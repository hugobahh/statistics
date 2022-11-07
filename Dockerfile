FROM golang:1.17-alpine AS builder


WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o credit-assignment main.go

FROM alpine as final
MAINTAINER "hugo_bh@yahoo.com"
LABEL service="credit-assignment"
LABEL owner="banwire"
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir /app
RUN chmod 777 /app
RUN mkdir /app/credit-assignment
RUN chmod 777 /app/credit-assignment
WORKDIR /app/credit-assignment
COPY --from=builder /app/configDB.json /app/credit-assignment
COPY --from=builder /app/configHttp.json /app/credit-assignment
COPY --from=builder /app/configMsSearchUsr.json /app/credit-assignment
COPY --from=builder /app/credit-assignment /app/credit-assignment
RUN chmod 777 /app/credit-assignment/credit-assignment
ENTRYPOINT ["/app/credit-assignment/credit-assignment"]
