FROM golang:1.23.1-bookworm as builder
COPY main.go .
RUN go build -o /zombie-example main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y procps
COPY --from=builder /zombie-example /zombie-example
ENTRYPOINT ["/zombie-example"]