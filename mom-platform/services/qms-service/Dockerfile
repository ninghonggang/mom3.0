FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/qms-service /app
COPY services/qms-service/config.yaml /config.yaml
EXPOSE 50053
ENTRYPOINT ["/app"]
