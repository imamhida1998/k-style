# Gunakan image Go resmi sebagai base image
FROM golang:1.20-alpine

# Set environment variable
ENV GO111MODULE=on

# Atur working directory di dalam container
WORKDIR /app

# Salin go.mod dan go.sum ke working directory
COPY go.mod go.sum ./

# Unduh dependencies
RUN go mod download

# Salin semua file dan folder ke working directory
COPY . .


# Build aplikasi
RUN go build -o main .

# Ekspose port yang digunakan aplikasi
EXPOSE 3000

# Perintah untuk menjalankan aplikasi
CMD ["./main"]
