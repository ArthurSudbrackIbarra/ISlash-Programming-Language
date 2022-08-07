FROM golang:1.19
WORKDIR /usr/islash
COPY . .
RUN go build -o islash
CMD ["sleep", "infinity"]