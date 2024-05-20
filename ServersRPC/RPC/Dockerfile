FROM golang:alpine as builder

WORKDIR /app


COPY .. .

#RUN go mod download
RUN go mod tidy

RUN go build -o ./main ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main .

EXPOSE 1234

CMD ["./main"]