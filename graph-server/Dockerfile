FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y procps

COPY graph-server /graph-server

ENTRYPOINT ["/graph-server"]
CMD ["run"]