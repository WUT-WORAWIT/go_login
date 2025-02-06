FROM golang:1.21.1 AS build

WORKDIR /app

# ใช้ go.mod จากโฟลเดอร์หลัก
COPY ./../go.mod ./../go.sum ./
# ดาวน์โหลดและติดตั้ง dependencies
RUN go mod download

COPY . .
RUN go build -o sell-app ./src/sell/sell.go

CMD ["./sell-app"]
