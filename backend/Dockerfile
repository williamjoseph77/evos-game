FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o evos-be

EXPOSE 8080

CMD ./evos-be