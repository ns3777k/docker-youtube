FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y procps dumb-init
COPY ./init.sh /init.sh

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/init.sh"]