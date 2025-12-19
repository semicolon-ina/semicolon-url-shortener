# Pake Go 1.24 (Alpine) sesuai request
FROM golang:1.24-alpine

RUN apk add --no-cache git

# 1. Set Root sebagai base
WORKDIR /app

# 2. Copy folder 'repo' (Shared Code) dari Root Context
COPY repo ./repo

# 3. Pindah ke folder project spesifik
WORKDIR /app/url-shortener

# 4. Copy go.mod & go.sum
# Perhatikan path-nya: kita ambil dari root context "url-shortener/..."
COPY ./go.mod ./go.sum ./

# 5. Download dependency
RUN go mod download

# 6. Copy sisa source code project
COPY url-shortener .

EXPOSE 8080

# 7. Jalankan main.go
# Karena WORKDIR udah di /app/url-shortener, tinggal panggil main.go
CMD ["go", "run", "main.go"]