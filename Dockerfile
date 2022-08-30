
FROM golang:1.16-alpine
WORKDIR /build
COPY . .
RUN go get -u github.com/lib/pq
RUN go get -u github.com/gorilla/mux

RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

