# Belajar Golang Dasar

Repositori ini berisi pembelajaran dasar bahasa pemrograman Go (Golang).

## Persiapan Awal

### Instalasi Go

Pastikan Go sudah terinstal di komputer Anda. Unduh dari [situs resmi Go](https://golang.org/dl/).

### Verifikasi Instalasi

Periksa apakah Go sudah terinstal dengan benar:

```bash
go version
```

# Buat direktori baru untuk project Anda

mkdir nama-project
cd nama-project

# Inisialisasi modul Go

go mod init nama-module

# Contoh pada project ini

go mod init belajar-golang-dasar

# Buat file dengan Notepad/text editor

notepad helloworld.go

# Atau dengan command line

echo 'package main

import "fmt"

func main() {
fmt.Println("Hello, World!")
}' > helloworld.go

# Menjalankan Langsung

go run helloworld.go

# Build program menjadi executable

go build helloworld.go

# Jalankan executable yang dihasilkan (Windows)

helloworld.exe

# Jalankan executable yang dihasilkan (Linux/Mac)

./helloworld

# Build dengan Nama Custom

go build -o nama-program helloworld.go

# Inisialisasi modul Go

go mod init nama-module

# Jalankan program Go

go run namafile.go

# Kompilasi program Go menjadi executable

go build namafile.go

# Unduh dan instal package

go get nama-package

# Jalankan unit test

go test ./...

# Format kode secara otomatis

go fmt ./...
