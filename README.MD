# 🍩 Donat

![icon](https://blogger.googleusercontent.com/img/a/AVvXsEj1Fy1Jq-dpT7v2lMuVutexOKsmA27qtJDxKk2Hp8rap-VHD7tWQkDeS3Sz4dHQH0Ux6Yx4hK1s1GHIOeeKINSp_t0dXZ-XwHewIXqferfYniZX2lnjbq1hu4Q16p469IwZoQaZUTyt_Y7I0VYqVkaf-fbTGZBfRnZfu4DycQAxjv9EEQMGdS5PlJ0NQNci)

**Donat** (Do Not Accept Trash) is a lightweight Command Line Interface (CLI) tool written in Go. It allows you to generate disposable email addresses instantly, straight from your terminal. Perfect for bypassing sign-up forms, avoiding spam, and testing applications.

> Powered by GuerrillaMail API.

## ✨ Features

* **Zero Friction:** Generate a usable email address in seconds.
* **Clipboard Ready:** Automatically copies the new email address to your clipboard.
* **Session Persistence:** Keeps your inbox session active locally until you decide to clear it.
* **Clean Output:** Minimalist, standard output without unnecessary bloat.

## 🚀 Installation

### Prerequisites

* **Go 1.20+** installed on your machine.
* *(Linux Only)* `xclip` or `xsel` for clipboard functionality (`sudo apt install xclip`).

### Install via Go

```bash
go install github.com/kikukafandi/donat@latest
```

## 📖 Usage

Donat uses a simple "Bake, Eat, Crumbs" philosophy.

### 1. Bake (Generate Identity)

Create a new disposable email address. The address is automatically saved to your session and copied to your clipboard.

```bash
donat bake
```

### 2. Eat (Check Inbox)

List all incoming messages in your current inbox.

```bash
donat eat
```

To read a specific message, provide the ID from the list:

```bash
donat eat 123
```

### 3. Crumbs (Clear Session)

Remove your local session file and "forget" the current email address.

```bash
donat crumbs
```

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
* **Clipboard:** [atotto/clipboard](https://github.com/atotto/clipboard)
* **Architecture:** Command-based modular design.

## ☕ Support This Project

If Donat helps you, consider supporting its development:

👉 [https://buymeacoffee.com/kikukafandi](https://buymeacoffee.com/kikukafandi)

Your support keeps this project alive and maintained. Thank you! 🙏

## 📄 License

This project is licensed under the **MIT License**.

See the full license text in the [`LICENSE`](./LICENSE) file.

---

## ☕ Support This Project

If Donat helps you, consider supporting its development:

👉 [https://buymeacoffee.com/kikukafandi](https://buymeacoffee.com/kikukafandi)

Your support keeps this project alive and maintained. Thank you! 🙏
