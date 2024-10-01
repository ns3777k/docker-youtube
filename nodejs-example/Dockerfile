FROM node:22-bookworm-slim

RUN apt-get update && apt-get install -y procps
COPY ./index.js /index.js

ENTRYPOINT ["/usr/local/bin/node"]
CMD ["index.js"]