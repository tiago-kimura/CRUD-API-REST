FROM golang:1.14.2-alpine3.11  as builder
WORKDIR /app
ENV PORT 8080
RUN apk update && \
    apk add --no-cache \
    coreutils \
    git \
    make
COPY . .
RUN make build
CMD ["/app/bin/linux_amd64/account"]
EXPOSE ${PORT}