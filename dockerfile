FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

# Download modules necessary to compile the project
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /ninja

CMD ["/ninja"]