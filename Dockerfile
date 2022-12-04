FROM golang:1.19-alpine3.17

EXPOSE 7777

WORKDIR /app

COPY . .

RUN go build

CMD [ "./gomq" ]
