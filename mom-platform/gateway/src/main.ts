import 'reflect-metadata';
import { NestFactory } from '@nestjs/core';
import { ValidationPipe, Logger } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { HttpExceptionFilter } from './common/filters/http-exception.filter';
import { LoggingInterceptor } from './common/interceptors/logging.interceptor';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const logger = new Logger('Bootstrap');

  // Global prefix is intentionally omitted so route paths stay as defined
  // in each controller (e.g. /api/mes/orders).

  // CORS
  app.enableCors({
    origin: true,
    credentials: true,
  });

  // Global validation pipe — validates DTOs decorated with class-validator
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      forbidNonWhitelisted: true,
      transform: true,
      transformOptions: { enableImplicitConversion: true },
    }),
  );

  // Global exception filter & request logging interceptor
  app.useGlobalFilters(new HttpExceptionFilter());
  app.useGlobalInterceptors(new LoggingInterceptor());

  // Swagger / OpenAPI documentation
  const config = new DocumentBuilder()
    .setTitle('MOM Platform API Gateway')
    .setDescription(
      'BFF (Backend-For-Frontend) gateway aggregating all MOM microservices: ' +
        'MES, QMS, EAM, WMS, MDM, APS, Trace, Andon and Dashboard.',
    )
    .setVersion('0.1.0')
    .addBearerAuth()
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api/docs', app, document);

  await app.listen(3000);
  logger.log('MOM Gateway running on http://localhost:3000');
  logger.log('Swagger docs available at http://localhost:3000/api/docs');
}
bootstrap();
