import {
  CallHandler,
  ExecutionContext,
  Injectable,
  NestInterceptor,
  Logger,
} from '@nestjs/common';
import { Observable, tap } from 'rxjs';

/**
 * Logs every HTTP request that flows through the gateway: method, URL,
 * response time and status. Installed globally from main.ts.
 */
@Injectable()
export class LoggingInterceptor implements NestInterceptor {
  private readonly logger = new Logger('HTTP');

  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const request = context.switchToHttp().getRequest();
    const { method, url } = request;
    const now = Date.now();

    return next.handle().pipe(
      tap({
        next: () => {
          const response = context.switchToHttp().getResponse();
          const { statusCode } = response;
          const duration = Date.now() - now;
          this.logger.log(`${method} ${url} ${statusCode} - ${duration}ms`);
        },
        error: (error) => {
          const duration = Date.now() - now;
          const status = error?.status ?? 500;
          this.logger.error(
            `${method} ${url} ${status} - ${duration}ms - ${error.message}`,
          );
        },
      }),
    );
  }
}
