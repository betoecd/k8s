FROM golang:1.15

WORKDIR /go/src/bitslab
COPY . . 
RUN GOOS=linux go build
CMD ["./bitslab"]
