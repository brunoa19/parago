FROM golang:stretch AS builder

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build

FROM ubuntu:latest AS worker

RUN apt-get update \
    && apt-get install -y \
        tzdata \
        ca-certificates

WORKDIR /app
COPY --from=builder /app/shipa-gen ./
CMD ["./shipa-gen"]
EXPOSE 8080