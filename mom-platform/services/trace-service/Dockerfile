FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/trace-service /app
COPY services/trace-service/config.yaml /config.yaml
EXPOSE 50057
ENTRYPOINT ["/app"]
