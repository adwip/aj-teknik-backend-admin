# AJ Teknik Backend Admin

Backend system for **Agung Jaya Teknik**, built with Go and Echo v5. This project handles the core logic for both the digital catalog (web shop) and the administrative management system.

## Project Structure

```text
.
├── common-lib/             # Shared infrastructure and utilities
│   ├── infrastructure/     # HTTP server setup
│   ├── logger/             # Centralized logging system
│   ├── metadata/           # Request/response metadata
│   ├── session/            # Session and result formatting
│   ├── sql_lib/            # SQL database wrappers
│   ├── stacktrace/         # Enhanced error handling
│   └── shared/             # Utility functions (uuid, etc.)
├── internal/               # Main application logic
│   ├── handlers/           # HTTP Request handlers (REST)
│   ├── interfaces/         # Route registrations
│   ├── model/              # Data models
│   │   ├── entities/       # Database table mappings
│   │   └── dtos/           # Data Transfer Objects for custom queries
│   ├── repositories/       # Data access layer (MySQL)
│   ├── usecases/           # Business logic implementation
│   │   ├── administration/ # Admin features (Brands, Categories)
│   │   └── product/        # Product catalog features
│   └── shared/             # Internal utilities
├── public/                 # Static assets
└── scripts/                # Development scripts
```

## Business Process

### 1. Web Shop (Digital Catalog)
Sistem ini dirancang untuk menampilkan produk kepada pelanggan dengan performa yang dioptimalkan:
- **Product Discovery**: Pelanggan dapat melihat daftar produk, melakukan pencarian via query, serta memfilter berdasarkan kategori atau brand.
- **Hierarchical Navigation**: Menampilkan kategori dalam struktur pohon (*tree structure*) untuk navigasi yang lebih mudah.
- **Direct Contact**: Alur bisnis difokuskan pada katalog digital di mana informasi stok tersedia, namun pembelian diarahkan melalui kontak langsung (sesuai kebutuhan bisnis Agung Jaya Teknik).

### 2. Admin Panel
Pusat kontrol untuk mengelola data operasional toko:
- **Category Management**: Mengelola struktur kategori induk dan anak secara fleksibel.
- **Brand Management**: Pengelolaan brand produk yang terintegrasi dengan jumlah produk yang terkait.
- **Product Management**: 
    - CRUD produk dengan validasi ketat.
    - Sinkronisasi data antara `PriceId`, `Code`, dan `Stock` (tipe `float64` untuk fleksibilitas unit produk).
    - Optimasi query custom melalui **DTO** untuk menggabungkan data produk, kategori, dan brand dalam satu response efisien.

## Tech Stack
- **Language**: Go
- **Web Framework**: Echo v5
- **Database**: MySQL (jmoiron/sqlx)
- **Validation**: go-playground/validator
- **Infrastructure**: Custom-built robust error stacktraces & session handling.

## Development
Project ini menggunakan `Makefile` untuk mempermudah operasional:
- `make run`: Menjalankan server dalam mode development.
- `make build`: Mengkompilasi project menjadi binary.
