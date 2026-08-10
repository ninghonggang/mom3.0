FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/andon-service /app
COPY services/andon-service/config.yaml /config.yaml
EXPOSE 50058
ENTRYPOINT ["/app"]
