FROM arm64v8/golang:1.19-alpine3.16

EXPOSE 7777

WORKDIR /app

COPY . .

RUN go build

CMD [ "./gomq" ]
