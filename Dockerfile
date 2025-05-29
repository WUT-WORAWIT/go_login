# เลือก base image สำหรับการเริ่มต้น
FROM golang:1.24.3

# ติดตั้ง PostgreSQL
# RUN apt-get update && apt-get install -y postgresql postgresql-contrib

# กำหนด working directory
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum สำหรับโปรเจ็กต์
COPY go.mod go.sum ./

# ดาวน์โหลดและติดตั้ง dependencies
RUN go mod download

# คัดลอกโค้ดทั้งหมดไปยัง working directory
COPY . .

# Compile แอปพลิเคชัน
RUN go build -o main .

# กำหนด port ที่แอปพลิเคชันจะใช้
EXPOSE 8080

# กำหนดคำสั่งที่จะใช้เมื่อ container ถูกเริ่มต้น
CMD ["./main"]
