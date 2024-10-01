FROM golang:1.23.1-bookworm as builder
COPY main.go .
RUN go build -o /image-optimization main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y procps dumb-init
COPY --from=builder /image-optimization /image-optimization
CMD ["/image-optimization", "--start-coordinator"]

ENTRYPOINT ["dumb-init", "--verbose", "--"]