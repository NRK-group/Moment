FROM golang:1.19.2
WORKDIR /moment-backend-img
COPY . /moment-backend-img
RUN go build -o moment-backend-img .
CMD [ "go", "run", "server.go" ]
