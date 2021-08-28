FROM golang:1.17

WORKDIR /opt/server

COPY . .
RUN go build -o server .
CMD [ "./server" ]
