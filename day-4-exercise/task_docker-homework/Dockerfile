# Stage 1

FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o backend

# Stage 2

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/backend .

COPY ./config/app/* ./

EXPOSE 8080

CMD ["./backend"]
