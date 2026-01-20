# Product Management API (Golang + Docker)

API RESTful sederhana untuk mengelola data produk, dibangun menggunakan bahasa pemrograman **Go Language** dan dideploy menggunakan **Docker** untuk memastikan skalabilitas dan konsistensi infrastruktur.

## 🚀 Fitur Utama
* **CRUD Operations**: Mendukung penambahan, pengambilan, pembaruan, dan penghapusan data produk.
* **PostgreSQL Integration**: Menggunakan PostgreSQL sebagai penyimpanan data utama.
* **Database Migration**: Automigrasi skema database menggunakan GORM.
* **Containerized Environment**: Aplikasi dan database berjalan sepenuhnya di dalam Docker.

## 🛠️ Teknologi yang Digunakan
* **Language**: Go (Golang)
* **Framework**: Gin Gonic
* **ORM**: GORM
* **Database**: PostgreSQL
* **DevOps**: Docker & Docker-Compose

## 📋 Syarat
Sebelum menjalankan proyek ini, pastikan Anda sudah menginstal:
* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Git](https://git-scm.com/)

## 🏃 Cara Menjalankan
1. **Clone Repository**
   ```bash
   git clone [https://github.com/username/namarepo.git](https://github.com/username/namarepo.git)
   cd 

2. Jalankan dengan Docker Compose Cukup satu perintah untuk membangun image dan menjalankan container:
    ```bash
   docker-compose up --build
3. Akses API Aplikasi akan berjalan di: http://localhost:8080 Database PostgreSQL dapat diakses di port: 5434
   
   



