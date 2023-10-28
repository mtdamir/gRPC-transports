FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /pricefetcher

EXPOSE 8080

CMD [ "/pricefetcher"]