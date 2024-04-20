FROM golang:latest

WORKDIR /opt/server

COPY server .
RUN go build -o server .
CMD [ "./server" ]
