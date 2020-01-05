FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o greeter-service/greeter-service greeter-service/main.go

FROM alpine
COPY --from=builder /app/greeter-service/greeter-service /app/greeter-service
EXPOSE 8091
ENTRYPOINT [ "/app/greeter-service" ]
