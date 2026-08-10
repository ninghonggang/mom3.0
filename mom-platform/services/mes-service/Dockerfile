FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/mes-service /app
COPY services/mes-service/config.yaml /etc/mom/config.yaml
EXPOSE 50052
ENTRYPOINT ["/app"]
