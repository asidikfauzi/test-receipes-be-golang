# Golang Recipes Test

## Panduan Instalasi

### 1. Membuat Skema Database PostgreSQL

Buat skema database PostgreSQL dengan nama "recipes_kaskus". Pastikan database telah berhasil dibuat.

### 2. Penggunaan UUIDv4 (Auto Generate)

Jalankan SQL console untuk mengaktifkan ekstensi UUIDv4 dengan perintah berikut:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### 3. Menjalankan Aplikasi
Untuk menjalankan aplikasi, Anda perlu mengonfigurasi dan menjalankan kode Go. Pastikan Anda memiliki Go (Golang) terinstal.

1. ***Pastikan semua dependensi aplikasi telah diunduh. Jika belum, Anda dapat mengunduhnya dengan perintah:***
```ssh
go get / go mod tidy
```

2.  ***Setelah semua dependensi diunduh, jalankan aplikasi dengan perintah:***
```ssh
go run main.go
```
Ini akan menjalankan aplikasi Anda, serta melakukan migrasi otomatis dan mengisi basis data Anda.

Harap dicatat bahwa Anda perlu mengatur konfigurasi basis data Anda di aplikasi Anda sesuai dengan koneksi PostgreSQL Anda sebelum menjalankannya. Pastikan juga bahwa server PostgreSQL Anda telah berjalan.

