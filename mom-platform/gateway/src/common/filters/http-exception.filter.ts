import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus,
  Logger,
} from '@nestjs/common';
import { Request, Response } from 'express';

/**
 * Global exception filter that normalises every error thrown in the
 * gateway into a consistent JSON envelope:
 *
 *   {
 *     "code": "ERROR_CODE",
 *     "message": "human readable message",
 *     "statusCode": 500,
 *     "timestamp": "ISO-8601",
 *     "path": "/api/..."
 *   }
 */
@Catch()
export class HttpExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(HttpExceptionFilter.name);

  catch(exception: unknown, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();

    let statusCode = HttpStatus.INTERNAL_SERVER_ERROR;
    let message = 'Internal server error';
    let code = 'INTERNAL_ERROR';
    let details: unknown;

    if (exception instanceof HttpException) {
      statusCode = exception.getStatus();
      const res = exception.getResponse();
      if (typeof res === 'string') {
        message = res;
      } else if (typeof res === 'object' && res !== null) {
        const r = res as Record<string, any>;
        message = r.message ?? message;
        code = r.code ?? code;
        details = r.details;
      }
    } else if (exception instanceof Error) {
      message = exception.message;
      this.logger.error(exception.stack);
    }

    // Map common HTTP statuses to stable error codes
    if (code === 'INTERNAL_ERROR') {
      code = this.codeForStatus(statusCode);
    }

    response.status(statusCode).json({
      code,
      message: Array.isArray(message) ? message : message,
      details,
      statusCode,
      timestamp: new Date().toISOString(),
      path: request.url,
    });
  }

  private codeForStatus(status: number): string {
    switch (status) {
      case HttpStatus.BAD_REQUEST:
        return 'BAD_REQUEST';
      case HttpStatus.UNAUTHORIZED:
        return 'UNAUTHORIZED';
      case HttpStatus.FORBIDDEN:
        return 'FORBIDDEN';
      case HttpStatus.NOT_FOUND:
        return 'NOT_FOUND';
      case HttpStatus.CONFLICT:
        return 'CONFLICT';
      case HttpStatus.UNPROCESSABLE_ENTITY:
        return 'UNPROCESSABLE_ENTITY';
      default:
        return 'INTERNAL_ERROR';
    }
  }
}
