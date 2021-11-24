FROM golang:1.17.3-alpine3.14
WORKDIR /app
COPY . .
#RUN make 
RUN go build -v ./cmd/apiserver
EXPOSE 8080
CMD ["/app/apiserver"]