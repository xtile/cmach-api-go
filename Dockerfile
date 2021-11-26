#Build stage

FROM golang:1.17.3-alpine3.14 AS builder
WORKDIR /app
COPY . .
#RUN make 
RUN go build -v ./cmd/apiserver


#Run stage
FROM alpine:3.14 AS builder
WORKDIR /app
COPY --from=builder /app/apiserver .


EXPOSE 8080
CMD ["/app/apiserver"]
