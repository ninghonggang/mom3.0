import {
  Injectable,
  OnModuleInit,
  Logger,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { existsSync, readdirSync } from 'fs';
import { join } from 'path';
import * as protoLoader from '@grpc/proto-loader';
import * as grpc from '@grpc/grpc-js';
import { deepCamelToSnake, deepSnakeToCamel } from './case.util';

/** gRPC status code -> HTTP status 映射 */
const GRPC_STATUS_HTTP: Record<number, number> = {
  1: HttpStatus.INTERNAL_SERVER_ERROR, // CANCELLED
  2: HttpStatus.INTERNAL_SERVER_ERROR, // UNKNOWN
  3: HttpStatus.BAD_REQUEST, // INVALID_ARGUMENT
  4: HttpStatus.GATEWAY_TIMEOUT, // DEADLINE_EXCEEDED
  5: HttpStatus.NOT_FOUND, // NOT_FOUND
  6: HttpStatus.CONFLICT, // ALREADY_EXISTS
  7: HttpStatus.FORBIDDEN, // PERMISSION_DENIED
  8: HttpStatus.TOO_MANY_REQUESTS, // RESOURCE_EXHAUSTED
  9: HttpStatus.CONFLICT, // FAILED_PRECONDITION
  10: HttpStatus.CONFLICT, // ABORTED
  11: HttpStatus.GONE, // OUT_OF_RANGE
  12: HttpStatus.NOT_IMPLEMENTED, // UNIMPLEMENTED
  13: HttpStatus.INTERNAL_SERVER_ERROR, // INTERNAL
  14: HttpStatus.SERVICE_UNAVAILABLE, // UNAVAILABLE
  15: HttpStatus.SERVICE_UNAVAILABLE, // DATA_LOSS
  16: HttpStatus.UNAUTHORIZED, // UNAUTHENTICATED
};

/**
 * 统一 gRPC 客户端工厂。
 *
 * 启动时用 @grpc/proto-loader 一次性加载 gateway/proto 下的全部 .proto，
 * 之后各业务域 service 通过 call() 调用对应 package/service/method。
 */
@Injectable()
export class GrpcService implements OnModuleInit {
  private readonly logger = new Logger(GrpcService.name);
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private pkg: any;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private readonly clientCache = new Map<string, any>();

  onModuleInit() {
    const protoRoot = join(__dirname, '..', '..', 'proto');
    if (!existsSync(protoRoot)) {
      this.logger.error(`Proto 目录不存在: ${protoRoot}`);
      return;
    }
    const files = this.collectProtos(protoRoot);
    const def = protoLoader.loadSync(files, {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true,
      includeDirs: [protoRoot],
    });
    this.pkg = grpc.loadPackageDefinition(def);
    this.logger.log(`已加载 ${files.length} 个 proto 文件`);
  }

  private collectProtos(dir: string): string[] {
    const out: string[] = [];
    const walk = (d: string) => {
      for (const e of readdirSync(d, { withFileTypes: true })) {
        const p = join(d, e.name);
        if (e.isDirectory()) walk(p);
        else if (e.name.endsWith('.proto')) out.push(p);
      }
    };
    walk(dir);
    return out;
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private resolvePackage(packageName: string): any {
    return packageName.split('.').reduce((acc: any, key) => acc?.[key], this.pkg);
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private getClient(packageName: string, serviceName: string, url: string): any {
    const key = `${packageName}.${serviceName}@${url}`;
    let client = this.clientCache.get(key);
    if (!client) {
      const ServiceCtor = this.resolvePackage(packageName)?.[serviceName];
      if (!ServiceCtor) {
        throw new HttpException(
          `gRPC service ${packageName}.${serviceName} 未找到`,
          HttpStatus.BAD_GATEWAY,
        );
      }
      client = new ServiceCtor(url, grpc.credentials.createInsecure());
      this.clientCache.set(key, client);
    }
    return client;
  }

  /**
   * 调用远端 gRPC 方法。
   * @param packageName proto package，如 'mom.mdm'
   * @param serviceName gRPC service，如 'MaterialService'
   * @param method RPC 方法名
   * @param request 请求对象（camelCase，会自动转 snake_case）
   * @param url 服务地址（如 mdm-service:50051）
   */
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  async call<T = any>(
    packageName: string,
    serviceName: string,
    method: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    request: any,
    url: string,
  ): Promise<T> {
    if (!this.pkg) {
      throw new HttpException('gRPC 客户端未初始化', HttpStatus.SERVICE_UNAVAILABLE);
    }
    const client = this.getClient(packageName, serviceName, url);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const reqSnake = deepCamelToSnake(request ?? {});
    return new Promise<T>((resolve, reject) => {
      client[method](reqSnake, (err: grpc.ServiceError | null, data: unknown) => {
        if (err) {
          const httpStatus = GRPC_STATUS_HTTP[err.code] ?? HttpStatus.INTERNAL_SERVER_ERROR;
          return reject(new HttpException(err.details || err.message, httpStatus));
        }
        resolve(deepSnakeToCamel<T>(data));
      });
    });
  }
}
