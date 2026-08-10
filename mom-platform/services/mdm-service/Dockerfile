FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/mdm-service /app
COPY services/mdm-service/config.yaml /config.yaml
EXPOSE 50051
ENTRYPOINT ["/app"]
