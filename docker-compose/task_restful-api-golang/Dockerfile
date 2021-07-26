# build stage
FROM golang:1.15-alpine AS build-env

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' 

# final stage
FROM alpine:latest

WORKDIR /main

# Copy result binary go app to /app folder
COPY --from=build-env /app/example /main
COPY --from=build-env /app/.env /main

ENTRYPOINT ["./example"]


