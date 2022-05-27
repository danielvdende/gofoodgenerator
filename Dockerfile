FROM golang:1.18.2-buster
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

ENTRYPOINT ["./gofoodgenerator"]