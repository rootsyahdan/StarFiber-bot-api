# StarFiber-bot-api

Deskripsi Produk: Proyek ini bertujuan untuk membuat sistem manajemen iuran internet bulanan yang memungkinkan pengguna untuk membayar iuran bulanan mereka dan menerima reminder melalui Telegram Bot. Tujuan produk ini adalah mempermudah pengguna dalam membayar iuran internet dan mengingatkan mereka secara berkala.

## Instalasi

Untuk menjalankan API ini, Anda perlu mengikuti langkah-langkah instalasi berikut:

1. Clone repositori ini ke komputer Anda.
2. Instal semua dependensi yang diperlukan dengan menjalankan perintah `go mod tidy`.
3. Konfigurasi koneksi database MySQL di file `.env` atau konfigurasi yang sesuai.
4. Jalankan API dengan perintah `go run main.go`.
5. API akan berjalan di `http://localhost:1312` secara default.

## Penggunaan API

Anda dapat menggunakan API ini untuk keperluan seperti Pendaftaran Pendaftaran Iuran Internet Bulanan, dapat membuat Admin,Users,Memberships dan Transaksi secara otomatis

## Endpoints

| Route         | HTTP Method | Deskripsi                           |
| ------------- | ----------- | ----------------------------------- |
| /admins       | POST        | Membuat Admins baru                 |
| /admins/login | POST        | Login data Admin berdasarkan Admins |

| Route      | HTTP Method | Deskripsi                            |
| ---------- | ----------- | ------------------------------------ |
| /users     | GET         | Mendapatkan semua data user          |
| /users/:id | GET         | Mendapatkan data user berdasarkan ID |
| /users     | POST        | Membuat user baru                    |
| /users/:id | PUT         | Mengubah data user berdasarkan ID    |
| /users/:id | DELETE      | Menghapus data user berdasarkan ID   |

| Route            | HTTP Method | Deskripsi                                   |
| ---------------- | ----------- | ------------------------------------------- |
| /memberships     | GET         | Mendapatkan semua data membership           |
| /memberships/:id | GET         | Mendapatkan data memberships berdasarkan ID |
| /memberships     | POST        | Membuat membership baru                     |
| /memberships/:id | PUT         | Mengubah data membership berdasarkan ID     |
| /memberships/:id | DELETE      | Menghapus data membership berdasarkan ID    |

| Route             | HTTP Method | Deskripsi                                    |
| ----------------- | ----------- | -------------------------------------------- |
| /transactions     | GET         | Mendapatkan semua data transactions          |
| /transactions/:id | GET         | Mendapatkan data transactions berdasarkan ID |
| /transactions     | POST        | Membuat transactions baru                    |
| /transactions/:id | PUT         | Mengubah data transactions berdasarkan ID    |
| /transactions/:id | DELETE      | Menghapus data transactions berdasarkan ID   |
