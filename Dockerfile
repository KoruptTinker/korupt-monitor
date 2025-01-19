FROM golang:1.23.5

WORKDIR /app 
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/

EXPOSE 7070

CMD ["./server"]
