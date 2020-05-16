FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go build -o counter/service/counter-service counter/service/main.go

FROM alpine
COPY --from=builder /app/counter/service/counter-service /app/counter-service
EXPOSE 8091
ENTRYPOINT [ "/app/counter-service" ]
