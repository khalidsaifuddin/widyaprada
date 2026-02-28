# SDD – Tech Stack & Spesifikasi Teknis

Dokumen ini mendefinisikan **pilihan teknologi** yang digunakan untuk membangun aplikasi Widyaprada. Semua implementasi (backend, frontend, database, cache, styling) mengacu pada keputusan di sini.

---

## 1. Backend

| Aspek        | Pilihan                | Keterangan |
|-------------|------------------------|------------|
| **Bahasa**  | **Go (Golang)**        | Performa tinggi, konkurensi bawaan, binary tunggal, cocok untuk API dan layanan backend. |
| **Arsitektur** | **Clean Architecture** | Pemisahan layer (domain, use case, delivery, infrastructure); maintainable, testable, dan independen dari framework/DB. |

### 1.1 Go (Golang)

| Pros | Cons |
|------|------|
| Performa tinggi, kompilasi ke binary native; startup cepat. | Kurang "sugar" dibanding bahasa lain (mis. generics baru di Go 1.18+). |
| Konkurensi bawaan (goroutine, channel) cocok untuk I/O dan banyak koneksi. | Error handling eksplisit (`if err != nil`) bisa terasa berulang. |
| Satu binary, deploy sederhana; tidak perlu runtime terpisah. | Ekosistem ORM/relasional tidak sebesar Java/C# (biasanya pakai query builder atau SQL). |
| Tooling standar kuat: `go build`, `go test`, `go mod`, formatter standar. | Dependency management (vendor, module) perlu disiplin agar repo tetap rapi. |
| Cepat dipelajari untuk tim yang sudah kenal C/Java; dokumentasi resmi bagus. | Untuk domain logic sangat kompleks, bahasa dengan type system kaya bisa lebih ekspresif. |

**Tradeoff:** Mengutamakan kesederhanaan, kecepatan eksekusi, dan kemudahan deploy; menerima lebih banyak boilerplate (error handling, kurang generics di kode lama) dan ekosistem relasional yang lebih "manual" dibanding ecosystem Java/Node.

### 1.2 Clean Architecture

| Pros | Cons |
|------|------|
| Domain dan use case tidak bergantung pada DB/framework; ganti tech bisa tanpa ubah bisnis logic. | Lebih banyak file dan layer; struktur awal terasa "berat" untuk fitur kecil. |
| Testing mudah: use case dan domain bisa di-unit test tanpa DB/HTTP. | Kurva belajar untuk tim baru; perlu disiplin agar tidak "bocor" (mis. entity pakai type dari DB). |
| Aturan bisnis di satu tempat; mengurangi duplikasi dan inkonsistensi. | Bisa over-engine jika scope aplikasi kecil; perlu menyesuaikan kedalaman layer dengan skala proyek. |
| Tim bisa kerja paralel per layer (domain vs delivery vs infrastructure). | Mapping antara entity domain ↔ DTO ↔ model DB menambah kode; perlu konvensi jelas. |

**Tradeoff:** Mengutamakan maintainability jangka panjang dan testability; menerima kompleksitas struktur dan waktu development awal lebih besar. Cocok untuk aplikasi yang akan berevolusi (fitur banyak, tim bisa bertambah).

Struktur layer Clean Architecture yang dipakai (ringkas):

- **Domain:** Entitas dan aturan bisnis murni.
- **Use case / Application:** Orkestrasi bisnis, interface repository.
- **Delivery:** HTTP handler (API), middleware (auth, logging).
- **Infrastructure:** Implementasi repository (PostgreSQL, Redis), koneksi eksternal.

---

## 2. Frontend

| Aspek        | Pilihan                | Keterangan |
|-------------|------------------------|------------|
| **Framework** | **React (ReactJS)**   | Library UI dominan, ekosistem besar, cocok untuk SPA dan dashboard. |
| **CSS**     | **Tailwind CSS**       | Utility-first, konsisten dengan design system, cepat untuk layout dan responsif. |
| **UI Library** | **shadcn/ui**       | Komponen siap pakai yang memakai Tailwind; berbasis Radix UI (aksesibel); copy-paste ke codebase (bukan dependency berat); cocok untuk form, tabel, modal, navigasi di aplikasi admin/ujikom. |

### 2.1 React (ReactJS)

| Pros | Cons |
|------|------|
| Ekosistem sangat besar: library, contoh, dan kandidat karyawan banyak. | Perlu pilihan tambahan untuk routing, state, data-fetching (tidak "batteries included"). |
| Komponen reusable dan komposisi yang jelas; cocok untuk UI kompleks (dashboard, form, tabel). | Perubahan cepat (React 18, concurrent features); beberapa lib pihak ketiga bisa tertinggal. |
| Virtual DOM dan React 18 (concurrent, Suspense) mendukung UX responsif. | Bundle size bisa membesar jika tidak hati-hati (code-splitting, lazy load perlu disengaja). |
| Satu arah data flow memudahkan debugging dan prediksi state. | JSX dan "JavaScript everywhere" tidak cocok untuk tim yang ingin pemisahan tegas HTML/CSS/JS. |

**Tradeoff:** Mengutamakan fleksibilitas dan ekosistem; menerima bahwa tim harus memilih dan memelihara stack (router, state, API client) serta menjaga performa bundle sendiri.

### 2.2 Tailwind CSS

| Pros | Cons |
|------|------|
| Utility-first: layout dan styling cepat tanpa keluar dari HTML/JSX. | Class name panjang di JSX; bisa mengurangi keterbacaan jika tidak dikelola (extract component). |
| Design system konsisten lewat config (warna, spacing, breakpoint); mudah di-branding. | Kurva belajar untuk yang terbiasa BEM/SCSS; perlu hafal/semantic class. |
| Purge/build hanya mengeluarkan CSS yang dipakai; bundle CSS kecil. | Custom design sangat "custom" kadang butuh `@apply` atau CSS biasa; bisa campur paradigma. |
| Responsif dan dark mode built-in lewat modifier (`md:`, `dark:`). | Ketergantungan pada Tailwind version; upgrade major bisa breaking. |

**Tradeoff:** Mengutamakan kecepatan development dan konsistensi; menerima bahwa markup jadi lebih padat dan tim harus berkomitmen pada konvensi (kapan pakai component vs utility langsung).

### 2.3 shadcn/ui

| Pros | Cons |
|------|------|
| Komponen masuk ke codebase (copy-paste); full control, tidak "black box", upgrade sukarela. | Perlu maintenance sendiri (patch keamanan/aksesibilitas di komponen yang di-copy). |
| Berbasis Radix UI; aksesibel (keyboard, ARIA, screen reader) out of the box. | Jumlah komponen terbatas dibanding library "full suite"; untuk niche case mungkin perlu buat sendiri. |
| Selaras dengan Tailwind; styling seragam, mudah dikustomisasi. | Awal setup (theme, dependency Radix) perlu sekali; setelah itu konsisten. |
| Cocok untuk admin/dashboard: Table, Form, Dialog, Dropdown, Navigation. | Bukan "design system siap pakai" seperti Ant Design; tampilan perlu dirapikan per halaman. |

**Tradeoff:** Mengutamakan kontrol dan aksesibilitas dengan tetap memakai Tailwind; menerima bahwa komponen "milik kita" jadi tanggung jawab maintenance dan tidak semua komponen tersedia siap pakai.

Alternatif yang dipertimbangkan: Radix UI (headless saja), DaisyUI (lebih "theme-ready"), Headless UI (resmi Tailwind Labs). shadcn/ui dipilih karena keseimbangan antara kemudahan pakai dan kontrol penuh dengan Tailwind.

---

## 3. Database

| Aspek     | Pilihan           | Keterangan |
|----------|-------------------|------------|
| **DB utama** | **PostgreSQL (PGSQL)** | Relasional, ACID, mendukung JSON, full-text search, cocok untuk data terstruktur (user, role, bank soal, ujian, jurnal). |

### 3.1 PostgreSQL

| Pros | Cons |
|------|------|
| ACID, transaksi andal; cocok untuk data kritis (user, nilai, bank soal). | Skala horizontal (sharding) lebih rumit daripada DB NoSQL; perlu perencanaan. |
| Tipe data kaya: JSON/JSONB, array, full-text search, range; fleksibel tanpa keluar dari SQL. | Tuning dan indexing perlu pemahaman (explain, index strategy); salah konfigurasi bisa lambat. |
| Open source, lisensi permissive; ekosistem dan tool (pgAdmin, migrasi) matang. | Replikasi dan HA (high availability) butuh setup (streaming replication, failover). |
| Cocok untuk relasi kompleks (role–permission, ujian–soal–jawaban); constraint dan foreign key menjaga integritas. | Backup/restore dan upgrade major version perlu prosedur yang jelas. |

**Tradeoff:** Mengutamakan integritas data dan fleksibilitas query (relasional + JSON); menerima bahwa scaling horizontal dan operasi lanjutan (HA, backup) membutuhkan keahlian dan perencanaan.

Schema dan migrasi mengacu pada:

- [docs/erd.puml](../erd.puml)
- File di [migrations/](../../migrations/)

---

## 4. Cache

| Aspek   | Pilihan   | Keterangan |
|--------|-----------|------------|
| **Cache** | **Redis** | In-memory, cocok untuk sesi, token, rate limit, cache response API, dan data sementara (misalnya state ujian CBT). |

### 4.1 Redis

| Pros | Cons |
|------|------|
| Sangat cepat (in-memory); latency rendah untuk session, cache, rate limit. | Data in-memory: restart/ crash bisa hilang kecuali pakai persistence (RDB/AOF); perlu konfigurasi. |
| Struktur data fleksibel: string, hash, list, set, sorted set; cocok untuk berbagai pola cache/session. | Memori terbatas; perlu kebijakan eviction (LRU dll.) dan monitoring penggunaan RAM. |
| Fitur siap pakai: TTL, pub/sub, Lua script; mendukung rate limiting dan state sementara (ujian CBT). | Bukan pengganti database; data penting harus tetap di PostgreSQL. |
| De facto standar untuk cache/session; driver dan dokumentasi banyak (termasuk Go). | Replikasi dan cluster (Redis Cluster) menambah kompleksitas operasional. |

**Tradeoff:** Mengutamakan kecepatan dan kemudahan untuk session/cache/rate limit; menerima bahwa Redis adalah lapisan tambahan (bukan penyimpanan utama) dan perlu kebijakan persistence/eviction serta monitoring memori.

Penggunaan khas: session/token login, cache query yang sering diakses, rate limiting, dan antrian/state singkat jika diperlukan nanti.

---

## 5. Komunikasi Backend–Frontend

| Aspek | Pilihan default | Keterangan |
|-------|------------------|------------|
| **Protokol** | **REST over HTTP(S)** dengan **JSON** | Endpoint GET/POST/PUT/DELETE; request/response JSON. Mudah di-debug (curl, DevTools), ekosistem React (fetch, axios, React Query) matang. |

Komunikasi dari React (browser) ke backend Go mengacu pada REST + JSON sebagai **default** untuk kemudahan development, debugging, dan kompatibilitas dengan tool dan middleware (CORS, auth header, dll.).

### 5.1 Opsi: gRPC / gRPC-Web

**Apakah mungkin pakai gRPC antara backend (Go) dan frontend (React)?**  
**Ya, mungkin**, dengan skema **gRPC-Web**: browser tidak mendukung gRPC murni (HTTP/2 + binary), sehingga dipakai klien gRPC-Web di frontend yang berbicara ke proxy atau backend yang mendukung gRPC-Web; proxy/backend itu yang menerjemahkan ke gRPC di sisi Go.

| Aspek | Keterangan |
|-------|------------|
| **Frontend** | Pakai klien gRPC-Web (mis. `@grpc/grpc-web`, `grpc-web`) di React; kode client bisa di-generate dari `.proto`. |
| **Backend / Proxy** | Go menerima gRPC; untuk browser perlu salah satu: (1) Envoy (atau proxy lain) yang menerima gRPC-Web dan meneruskan gRPC ke Go, atau (2) wrapper gRPC-Web di Go (mis. `github.com/improbable-eng/grpc-web/go/grpcweb`). |
| **Contract** | Definisi API di file `.proto` (Protocol Buffers); typed, bisa dipakai bersama Go dan generated JS/TS. |

| Pros (gRPC/gRPC-Web) | Cons |
|----------------------|------|
| Kontrak API terdefinisi (`.proto`), type-safe di backend dan frontend. | Setup lebih rumit: proxy atau wrapper gRPC-Web, generate kode dari `.proto`. |
| Payload binary (protobuf), biasanya lebih kecil dan cepat daripada JSON. | Debug lebih susah (binary); perlu tool (grpcurl, Postman dll.) untuk inspeksi. |
| Cocok untuk streaming (server/client/bidirectional) jika nanti dibutuhkan. | CORS dan integrasi dengan auth (cookie/header) perlu diperhatikan di proxy. |
| Satu definisi API untuk banyak klien (web, mobile, layanan lain). | Tim harus terbiasa dengan protobuf dan alur generate/update `.proto`. |

**Tradeoff:** gRPC/gRPC-Web mengutamakan kontrak ketat, performa, dan kemungkinan streaming; menerima kompleksitas setup dan operasional (proxy, tooling, debugging). Untuk kebanyakan CRUD dan dashboard Widyaprada, **REST + JSON tetap disarankan** sebagai default. gRPC/gRPC-Web masuk akal jika: (1) ada rencana klien lain (mobile, layanan internal) yang memakai API yang sama, (2) butuh streaming (real-time), atau (3) tim sudah nyaman dengan protobuf dan siap maintain infrastruktur gRPC-Web.

---

## Ringkasan Tech Stack

| Layer       | Teknologi              |
|-------------|-------------------------|
| Backend     | Go + Clean Architecture |
| Frontend    | ReactJS                 |
| CSS         | Tailwind CSS            |
| UI Library  | shadcn/ui               |
| Database    | PostgreSQL              |
| Cache       | Redis                   |
| Backend–Frontend | REST + JSON (default); gRPC/gRPC-Web opsi |

---

## Dokumen terkait

- [SDD README](README.md) – Ruang lingkup dan struktur SDD.
- [ERD](../erd.puml) – Diagram entitas dan relasi database.
- [Migrations](../../migrations/) – SQL schema.

Perubahan tech stack harus didiskusikan dan diperbarui di dokumen ini serta dampaknya di SDD modul terkait.
