FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o logger-service/logger-service logger-service/main.go

FROM alpine
COPY --from=builder /app/logger-service/logger-service /app/logger-service
ENTRYPOINT [ "/app/logger-service" ]
