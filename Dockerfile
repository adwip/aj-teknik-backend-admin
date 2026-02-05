# STAGE 1: Builder
# Kita gunakan image golang berbasis alpine agar ringan namun lengkap dengan toolchain Go
FROM golang:1.23-alpine AS builder

# Install git dan build-base (perlu jika ada library yang menggunakan CGO)
RUN apk add --no-cache git build-base

# Tentukan direktori kerja di dalam container
WORKDIR /app

# Copy file dependency terlebih dahulu. 
# Ini adalah trik agar Docker bisa menyimpan cache layer 'go mod download' 
# selama tidak ada perubahan pada file-file ini.
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code proyek (termasuk .env jika ada, walaupun sebaiknya di-exclude via .dockerignore)
COPY . .

# Lakukan proses build.
# -o: Menentukan nama output file biner
# CGO_ENABLED=0: Membuat file biner statis (tidak bergantung pada shared library sistem)
# GOOS=linux: Memastikan target OS adalah linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web/main.go


# STAGE 2: Runner (Final Image)
# Kita gunakan image alpine yang sangat kecil (hanya ~5MB) untuk menjalankan aplikasi
FROM alpine:latest

# Tambahkan sertifikat SSL (diperlukan jika aplikasi melakukan request HTTPS ke luar)
# dan timezone data jika aplikasi membutuhkan waktu lokal yang akurat
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy file biner yang sudah jadi dari stage 'builder'
COPY --from=builder /app/main .

# Copy folder pendukung lainnya jika ada (karena aplikasi Anda melayani file statis di /public)
COPY --from=builder /app/public ./public

# Expose port sesuai dengan yang digunakan aplikasi Anda (misal 8080)
EXPOSE 8080

# Jalankan biner aplikasi
CMD ["./main"]
