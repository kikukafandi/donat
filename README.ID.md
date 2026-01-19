# 🍩 Donat

**Donat** ("Do Not Accept Trash") adalah aplikasi Command Line Interface (CLI) ringan yang ditulis dengan Go. Alat ini berfungsi untuk membuat email sementara (disposable email) secara instan langsung dari terminal. Sangat berguna untuk menghindari spam saat mendaftar di website yang mencurigakan atau untuk keperluan testing aplikasi.

> Menggunakan API dari GuerrillaMail.

## ✨ Fitur Utama

* **Cepat & Mudah:** Buat email siap pakai hanya dalam hitungan detik.
* **Auto-Copy:** Alamat email otomatis disalin ke clipboard laptopmu.
* **Sesi Tersimpan:** Sesi inbox tetap aktif di laptop sampai kamu menghapusnya, jadi tidak perlu takut email hilang saat terminal ditutup.
* **Tampilan Bersih:** Output yang rapi dan mudah dibaca tanpa elemen yang membingungkan.

## 🚀 Instalasi

### Syarat
* **Go 1.20+** sudah terinstall.
* *(Khusus Linux)* Pastikan sudah menginstall `xclip` atau `xsel` agar fitur copy-paste jalan (`sudo apt install xclip`).

### Cara Install
Jalankan perintah ini di terminal:

```bash
go install [github.com/kikukafandi/donat@latest](https://github.com/kikukafandi/donat@latest)

```

## 📖 Cara Penggunaan

Donat menggunakan filosofi sederhana: "Bake, Eat, Crumbs".

### 1. Bake (Buat Email Baru)

Memanggang (membuat) identitas email baru. Email akan disimpan di sesi lokal dan dicopy ke clipboard.

```bash
donat bake

```

### 2. Eat (Cek Inbox & Baca Pesan)

Melihat daftar pesan masuk di piring (inbox) kamu.

```bash
donat eat

```

Untuk membaca isi pesan tertentu, masukkan ID pesan dari daftar tadi:

```bash
donat eat 123

```

### 3. Crumbs (Hapus Sesi)

Membersihkan remahan (menghapus sesi). Ini akan menghapus data sesi lokal dan melupakan email yang sedang aktif.

```bash
donat crumbs

```

## 🛠️ Teknologi

* **Bahasa:** Go (Golang)
* **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
* **Clipboard:** [atotto/clipboard](https://github.com/atotto/clipboard)

## 📄 Lisensi

Proyek ini dilisensikan di bawah MIT License. Silakan gunakan dan modifikasi sesuka hati!

