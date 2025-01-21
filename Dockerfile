FROM golang:1.23.5

WORKDIR /app 
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    pkg-config \
    libx11-dev \
    libxtst-dev \
    libxinerama-dev \
    libx11-xcb-dev \
    libxkbcommon-dev \
    libxkbcommon-x11-dev

RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server/

EXPOSE 7070

CMD ["./server"]
