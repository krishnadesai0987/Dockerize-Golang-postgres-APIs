
FROM golang:1.16-alpine
RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/krishnadesai0987/sample-app   
RUN cd /build && git clone https://github.com/krishnadesai0987/sample-app.git

COPY . .
RUN go get -u github.com/lib/pq
RUN go get -u github.com/gorilla/mux

RUN  go build -o main .
EXPOSE 8080
CMD ["./main"]
