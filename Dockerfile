# payment image

FROM golang:1.16-alpine as payment

WORKDIR /payment

COPY go.mod ./
COPY  go.sum ./
RUN go mod download

COPY . .

RUN go build -o payment cmd/payment/api/main.go 

EXPOSE 3000

CMD [ "./payment" ]


# consumer image
FROM golang:1.16-alpine as consumer

WORKDIR /app
COPY --from=payment ./app  ./
RUN go build -o consumer cmd/payment/queue/main.go 

CMD [ "./consumer" ]