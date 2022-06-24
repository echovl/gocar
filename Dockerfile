FROM golang:1.18-buster

RUN mkdir /app

COPY . /app

WORKDIR /app

# Build golang app
RUN go build -o server .

CMD [ "/app/server" ]
