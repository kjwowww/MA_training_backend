# ใช้ base image ของ golang
FROM golang:1.22.5-alpine

# ตั้งค่า environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# สร้างและตั้งค่าไดเรกทอรีสำหรับโปรเจค
WORKDIR /app

# คัดลอก go.mod และ go.sum เพื่อดาวน์โหลด dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

# คัดลอกไฟล์โปรเจคทั้งหมด
COPY . .

# คอมไพล์โปรเจค
RUN go build -o ma-backend-training .

# สร้างไดเรกทอรีสำหรับการเก็บไฟล์ที่อัปโหลด
RUN mkdir -p /app/uploads

# ตั้งค่า port ที่จะ expose
EXPOSE 3000

# คำสั่งเพื่อรันแอปพลิเคชัน
CMD ["./ma-backend-training"]
