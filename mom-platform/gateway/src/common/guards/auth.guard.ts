import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
  Logger,
} from '@nestjs/common';
import { Observable } from 'rxjs';

/**
 * JWT authentication guard (stub).
 *
 * In production this will verify the access token carried in the
 * `Authorization: Bearer <token>` header, typically by calling the
 * platform's auth/SSO service or validating a signed JWT locally.
 *
 * For Phase 1 the guard is disabled so the gateway is reachable
 * without authentication. Wire it up globally (APP_GUARD) once the
 * auth service is integrated.
 */
@Injectable()
export class AuthGuard implements CanActivate {
  private readonly logger = new Logger(AuthGuard.name);

  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const request = context.switchToHttp().getRequest();
    const authHeader = request.headers['authorization'];

    // TODO: validate the JWT token against the auth service / key set.
    // Example:
    //   const token = authHeader?.replace('Bearer ', '');
    //   if (!token) throw new UnauthorizedException('Missing token');
    //   const payload = await this.jwtService.verifyAsync(token);
    //   request.user = payload;
    //   return true;

    if (!authHeader) {
      this.logger.debug('No Authorization header — allowing (stub mode)');
      // return true to keep the gateway open during development;
      // switch to throwing UnauthorizedException once auth is wired up.
      return true;
    }

    return true;
  }
}
