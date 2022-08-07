FROM golang:1.19
WORKDIR /usr/local/islash
COPY . .
RUN go build -o /usr/bin/islash
CMD ["sleep", "infinity"]
