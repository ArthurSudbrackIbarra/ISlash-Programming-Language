# Stage to build the ISlash executable using Go.
FROM golang:1.19 AS builder
WORKDIR /usr/local/islash
COPY . .
RUN go build -o islash

# Copying the ISlash executable to another container.
FROM alpine:3.14
COPY --from=builder /usr/local/islash/islash /usr/bin
CMD ["sleep", "infinity"]
