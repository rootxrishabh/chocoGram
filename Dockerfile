FROM golang:latest

WORKDIR /

ADD . .

RUN go build -o socialMediaApp

CMD ["./socialMediaApp"]