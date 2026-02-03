# AJ Teknik Backend Admin

Sistem backend untuk **Agung Jaya Teknik**, dibangun menggunakan Go dan Echo v5. Proyek ini menangani logika inti untuk katalog digital (toko web) dan sistem manajemen administratif.

## Struktur Proyek

```text
.
├── common-lib/             # Infrastruktur bersama dan utilitas
│   ├── infrastructure/     # Pengaturan server HTTP
│   ├── logger/             # Sistem pencatatan (logging) terpusat
│   ├── metadata/           # Metadata permintaan/tanggapan
│   ├── session/            # Format sesi dan hasil
│   ├── sql_lib/            # Pembungkus (wrapper) database SQL
│   ├── stacktrace/         # Penanganan kesalahan tingkat lanjut
│   └── shared/             # Fungsi utilitas (uuid, dsb.)
├── internal/               # Logika utama aplikasi
│   ├── handlers/           # Penangan data HTTP (REST)
│   ├── interfaces/         # Registrasi rute (routing)
│   ├── model/              # Model data
│   │   ├── entities/       # Pemetaan tabel database
│   │   └── dtos/           # Objek Transfer Data untuk kueri kustom
│   ├── repositories/       # Lapisan akses data (MySQL)
│   ├── usecases/           # Implementasi logika bisnis
│   │   ├── administration/ # Fitur admin (Brand, Kategori)
│   │   └── product/        # Fitur katalog produk
│   └── shared/             # Utilitas internal
├── public/                 # Aset statis
└── scripts/                # Skrip pengembangan
```

## Proses Bisnis

### 1. Toko Web (Katalog Digital)
Sistem ini dirancang untuk menampilkan produk kepada pelanggan dengan performa yang dioptimalkan:
- **Pencarian Produk**: Pelanggan dapat melihat daftar produk, melakukan pencarian, serta memfilter berdasarkan kategori atau brand.
- **Navigasi Hirarki**: Menampilkan kategori dalam struktur pohon untuk navigasi yang lebih mudah.
- **Kontak Langsung**: Alur bisnis difokuskan pada katalog digital di mana informasi stok tersedia, namun pembelian diarahkan melalui kontak langsung (sesuai kebutuhan bisnis Agung Jaya Teknik).

### 2. Panel Admin
Pusat kendali untuk mengelola data operasional toko:
- **Manajemen Kategori**: Mengelola struktur kategori induk dan anak secara fleksibel.
- **Manajemen Brand**: Pengelolaan brand produk yang terintegrasi dengan jumlah produk terkait.
- **Manajemen Produk**: 
    - Operasi CRUD produk dengan validasi ketat.
    - Sinkronisasi data antara `PriceId`, `Code`, dan `Stock` (menggunakan tipe `float64` untuk fleksibilitas unit produk).
    - Optimasi kueri kustom melalui **DTO** untuk menggabungkan data produk, kategori, dan brand dalam satu tanggapan yang efisien.

## Teknologi Utama
- **Bahasa Pemrograman**: Go
- **Kerangka Kerja Web**: Echo v5
- **Basis Data**: MySQL (jmoiron/sqlx)
- **Validasi**: go-playground/validator
- **Infrastruktur**: Penanganan kesalahan stacktrace & manajemen sesi yang kokoh.

## Pengembangan
Proyek ini menggunakan `Makefile` untuk mempermudah operasional:
- `make run`: Menjalankan server dalam mode pengembangan.
- `make build`: Mengompilasi proyek menjadi file biner.
