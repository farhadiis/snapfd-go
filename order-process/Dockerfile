FROM golang:1.19-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /go/bin/order-process


#FROM scratch
#COPY --from=builder /go/bin/order-process /go/bin/order-process
EXPOSE 3000
ENTRYPOINT ["/go/bin/order-process"]
