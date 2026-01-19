# 🍩 Donat

![icon](https://blogger.googleusercontent.com/img/a/AVvXsEj1Fy1Jq-dpT7v2lMuVutexOKsmA27qtJDxKk2Hp8rap-VHD7tWQkDeS3Sz4dHQH0Ux6Yx4hK1s1GHIOeeKINSp_t0dXZ-XwHewIXqferfYniZX2lnjbq1hu4Q16p469IwZoQaZUTyt_Y7I0VYqVkaf-fbTGZBfRnZfu4DycQAxjv9EEQMGdS5PlJ0NQNci)

**Donat** (Do Not Accept Trash) adalah aplikasi Command Line Interface (CLI) ringan yang ditulis menggunakan Go. Alat ini memungkinkan kamu membuat email sementara secara instan langsung dari terminal. Cocok untuk menghindari spam, mendaftar di situs yang mencurigakan, atau melakukan testing aplikasi.

> Menggunakan API dari GuerrillaMail.

---

## ✨ Fitur Utama

* **Cepat & Mudah:** Buat email siap pakai hanya dalam hitungan detik.
* **Auto-Copy:** Alamat email otomatis disalin ke clipboard.
* **Sesi Tersimpan:** Sesi inbox tetap aktif secara lokal sampai kamu menghapusnya.
* **Tampilan Bersih:** Output rapi tanpa elemen yang tidak perlu.

---

## 🚀 Instalasi

### Syarat

* **Go 1.20+** sudah terpasang.
* *(Khusus Linux)* Pastikan sudah menginstall `xclip` atau `xsel` agar fitur copy berjalan (`sudo apt install xclip`).

### Cara Install

Jalankan perintah berikut di terminal:

```bash
go install github.com/kikukafandi/donat@latest
```

---

## 📖 Cara Penggunaan

Donat menggunakan filosofi sederhana: **Bake, Eat, Crumbs**.

### 1. Bake (Buat Email Baru)

Membuat identitas email baru. Email akan disimpan di sesi lokal dan otomatis dicopy ke clipboard.

```bash
donat bake
```

### 2. Eat (Cek Inbox & Baca Pesan)

Melihat daftar pesan masuk di inbox kamu.

```bash
donat eat
```

Untuk membaca isi pesan tertentu, masukkan ID pesan dari daftar:

```bash
donat eat 123
```

### 3. Crumbs (Hapus Sesi)

Menghapus sesi lokal dan melupakan email yang sedang aktif.

```bash
donat crumbs
```

---

## 🛠️ Teknologi

* **Bahasa:** Go (Golang)
* **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
* **Clipboard:** [atotto/clipboard](https://github.com/atotto/clipboard)

---

## ☕ Dukung Proyek Ini

Jika Donat membantumu, pertimbangkan untuk mendukung pengembangannya:

👉 [https://buymeacoffee.com/kikukafandi](https://buymeacoffee.com/kikukafandi)

Dukunganmu sangat berarti agar proyek ini tetap hidup dan terawat. Terima kasih! 🙏

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah **MIT License**.

Silakan gunakan, modifikasi, dan distribusikan sesuai kebutuhan.

Lihat teks lisensi lengkap di file [`LICENSE`](./LICENSE).
