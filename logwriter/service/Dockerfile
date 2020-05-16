FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o logwriter/service/logwriter-service logwriter/service/main.go

FROM alpine
COPY --from=builder /app/logwriter/service/logwriter-service /app/logwriter-service
ENTRYPOINT [ "/app/logwriter-service" ]
