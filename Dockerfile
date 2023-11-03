FROM golang:alpine3.18 AS builder
RUN mkdir /app
ADD . /app 
WORKDIR /app 
RUN go build main.go 

FROM alpine:3.18 
WORKDIR /root/ 
COPY --from=builder /app/main .
EXPOSE 8000
CMD [ "./main" ]