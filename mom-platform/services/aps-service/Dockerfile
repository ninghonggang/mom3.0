FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY dist/aps-service /app
COPY services/aps-service/config.yaml /config.yaml
EXPOSE 50056
ENTRYPOINT ["/app"]
