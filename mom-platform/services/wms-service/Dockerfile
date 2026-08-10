FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/wms-service /app
COPY services/wms-service/config.yaml /config.yaml
EXPOSE 50055
ENTRYPOINT ["/app"]
