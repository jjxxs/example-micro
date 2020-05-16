FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o client/service/client client/service/main.go

FROM alpine
COPY --from=builder /app/client/service/client /app/client
EXPOSE 8091
ENTRYPOINT [ "/app/client" ]
