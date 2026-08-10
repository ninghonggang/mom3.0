FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/eam-service /app
COPY services/eam-service/config.yaml /config.yaml
EXPOSE 50054
ENTRYPOINT ["/app"]
