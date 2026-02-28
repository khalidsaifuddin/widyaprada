# [PRD] Fitur Manajemen Pengguna
## Product Requirements Document | Auth & RBAC

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Pengguna (List, Detail, Create, Edit, Delete)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Manajemen Pengguna
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-01
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Manajemen Pengguna memungkinkan Admin Satker (jika diizinkan) dan Super Admin untuk melihat daftar pengguna, melihat detail pengguna, menambah pengguna baru, mengubah data pengguna, dan menghapus pengguna dari satu area yang terstruktur. **Satu pengguna dapat memiliki lebih dari satu role** (misalnya Widyaprada sekaligus Admin Uji Kompetensi, atau Widyaprada sekaligus Verifikator Uji Kompetensi). Saat Create atau Edit pengguna, pengelola memilih **satu atau lebih role** untuk pengguna tersebut; menu dan wewenang yang tampil setelah login adalah gabungan dari semua role yang dimiliki. Pengguna dapat mencari, memfilter, dan mengelola akun sesuai wewenang masing-masing (Admin Satker terbatas pada satker/unit kerjanya; Super Admin untuk seluruh sistem) sehingga data pengguna tetap rapi dan akses terkendali.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur Manajemen Pengguna dirumuskan per role dalam pola user story berikut, dalam format tabular.

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Admin Satker / Super Admin | pengelola pengguna | melihat daftar pengguna dalam satu halaman/list yang terstruktur (nama, email/username, role/roles, satker, status) | dengan cepat menemukan dan memantau pengguna yang saya kelola |
| 2 | Admin Satker / Super Admin | pengelola pengguna | mencari pengguna berdasarkan nama, email, atau username | menemukan pengguna tertentu tanpa harus menggulir panjang |
| 3 | Admin Satker / Super Admin | pengelola pengguna | memfilter daftar berdasarkan role, satker, atau status (aktif/nonaktif) | fokus pada kelompok pengguna yang relevan dengan tugas saya |
| 4 | Admin Satker / Super Admin | pengelola pengguna | mengurutkan daftar (misalnya nama, tanggal daftar) | mengorganisir tampilan sesuai kebutuhan |
| 5 | Admin Satker / Super Admin | pengelola pengguna | membuka detail satu pengguna dari list (klik baris atau tombol Detail) | melihat informasi lengkap pengguna sebelum mengedit atau mengambil keputusan |
| 6 | Admin Satker / Super Admin | pengelola pengguna | menambah pengguna baru melalui form Create yang jelas (nama, email/username, kata sandi, satu atau lebih role, satker jika berlaku) | mendaftarkan anggota baru ke sistem; satu user bisa punya banyak role (mis. Widyaprada + Admin Uji Kompetensi) |
| 7 | Admin Satker / Super Admin | pengelola pengguna | mengubah data pengguna yang ada (Edit) dari halaman detail atau dari list, termasuk menambah atau mengurangi role pengguna | memperbarui informasi atau role tanpa harus menghapus dan buat ulang |
| 8 | Admin Satker / Super Admin | pengelola pengguna | menghapus atau menonaktifkan pengguna (Delete) dengan konfirmasi yang jelas | mencabut akses pengguna yang tidak lagi berhak atau salah input |
| 9 | Admin Satker | Admin Satker | hanya melihat dan mengelola pengguna dalam satker/unit kerja/lembaga/instansi saya | tidak mengakses atau mengubah pengguna di satker lain |
| 10 | Super Admin | Super Admin | melihat dan mengelola semua pengguna di sistem | mengawasi dan mengkonfigurasi seluruh akun |
| 11 | Admin Satker / Super Admin | pengelola pengguna | mendapat umpan balik jelas setelah Create/Edit/Delete (sukses atau pesan error) | tahu apakah aksi berhasil dan apa yang harus dilakukan jika gagal |
| 12 | Admin Satker / Super Admin | pengelola pengguna | paginasi atau lazy load pada list jika data banyak | daftar tetap cepat dan nyaman dipakai |
| 13 | Admin Satker / Super Admin | pengelola pengguna | tombol/aksi (Create, Edit, Delete) menampilkan loading saat diproses | yakin bahwa aplikasi sedang bekerja dan tidak hang |
| 14 | Admin Satker / Super Admin | pengelola pengguna | saat Create pengguna yang salah satu role-nya **Widyaprada**, sistem otomatis juga membuat data Widyaprada (profil/entitas Widyaprada) untuk pengguna tersebut | pengguna dengan role Widyaprada langsung memiliki data/profil Widyaprada yang lengkap |
| 15 | Admin Satker / Super Admin | pengelola pengguna | memilih lebih dari satu role untuk satu pengguna (multi-select role) saat Create atau Edit | satu user bisa sekaligus Widyaprada dan Admin Uji Kompetensi, atau Widyaprada dan Verifikator Uji Kompetensi, dll |
| 15 | Admin Satker / Super Admin | pengelola pengguna | saat Delete, wajib mengisi alasan/deskripsi kenapa data ini dihapus dalam prompt dialog konfirmasi | ada jejak dokumentasi alasan penghapusan dan mengurangi hapus tidak sengaja |

**Keterangan**: *Admin Satker* = akses Manajemen Pengguna jika role-nya diizinkan; *Super Admin* = akses penuh. Widyaprada tidak mengakses fitur ini. **Satu pengguna dapat memiliki lebih dari satu role.** Jika salah satu role yang dipilih adalah **Widyaprada**, Create pengguna wajib disertai pembuatan data Widyaprada.

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **List**: Daftar pengguna yang terstruktur dengan kolom relevan (nama, email/username, **role(s)**, satker, status), search, filter, sort, dan paginasi.
- **Detail**: Satu halaman/view untuk melihat informasi lengkap satu pengguna sebelum Edit atau Delete.
- **Create**: Form yang jelas untuk menambah pengguna baru dengan validasi (email unik, kata sandi memenuhi kebijakan, role dan satker sesuai wewenang). **Khusus role Widyaprada**: ketika pengguna baru dibuat dengan role Widyaprada, sistem harus sekaligus membuat/mendaftarkan data Widyaprada yang terkait (entitas/data profil Widyaprada untuk pengguna tersebut).
- **Edit**: Form untuk mengubah data pengguna dengan validasi yang sama; perubahan kata sandi opsional (field kosong = tidak diubah).
- **Delete**: Aksi hapus atau nonaktifkan dengan konfirmasi; **wajib ada prompt untuk mengisi deskripsi/alasan kenapa data ini dihapus** (field teks dalam dialog konfirmasi). Setelah konfirmasi dan alasan diisi, baru pesan sukses/error ditampilkan.
- **Umpan balik**: Loading pada tombol/aksi; pesan sukses atau error yang ramah dan jelas setelah setiap aksi.
- **Wewenang**: Admin Satker hanya mengelola pengguna dalam satker/unit kerjanya; Super Admin mengelola semua.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- List tanpa search/filter sehingga sulit menemukan pengguna saat data banyak.
- Form Create/Edit yang tidak memberi tahu field mana yang salah (validasi tidak jelas).
- Delete tanpa konfirmasi sehingga risiko hapus tidak sengaja.
- Delete tanpa wajib mengisi alasan/deskripsi penghapusan sehingga tidak ada jejak dokumentasi kenapa pengguna dihapus.
- Admin Satker bisa melihat atau mengubah pengguna di satker lain.
- Pesan error yang teknis atau membingungkan.

---

## 3. Antarmuka Pengguna (UI)

*Manajemen Pengguna terdiri atas: halaman List, halaman/view Detail, form Create, form Edit, dan aksi Delete (biasanya dari Detail atau dari list). Semua mengikuti design system aplikasi.*

---

### 3.1 List Pengguna

**Deskripsi**: Halaman yang menampilkan daftar pengguna dalam bentuk tabel atau card list.

**Elemen yang Harus Ada**:
- **Judul halaman**: Misalnya "Manajemen Pengguna" atau "Daftar Pengguna".
- **Tombol "Tambah Pengguna"** (atau "Create"): Memicu navigasi ke form Create.
- **Kotak pencarian**: Untuk mencari berdasarkan nama, email, atau username (placeholder jelas).
- **Filter** (opsional tapi disarankan): Dropdown atau pilihan untuk filter berdasarkan role, satker, status (aktif/nonaktif).
- **Tabel/List**: Kolom minimal — Nama (atau nama lengkap), Email/Username, **Role(s)** (satu atau lebih role, tampil sebagai daftar atau label), Satker/Unit Kerja (jika berlaku), Status (Aktif/Nonaktif). Kolom aksi: tombol atau link "Detail", "Edit", "Hapus" (sesuai wewenang).
- **Paginasi**: Jika data banyak, tampilkan paginasi (nomor halaman atau "Load more") agar performa tetap baik.
- **Sort**: Header kolom dapat diklik untuk mengurutkan (nama, tanggal, dll) jika diinginkan.

**Tata Letak dan Keterbacaan**:
- Tabel/list rapi; jarak antar baris nyaman; teks tidak terpotong sembarangan.
- Tombol aksi (Detail, Edit, Hapus) konsisten per baris dan mudah diklik (termasuk di layar sentuh).
- Di layar kecil (ponsel/tablet), list dapat di-responsive (misalnya card per pengguna atau tabel scroll horizontal).

**Umpan Balik**:
- Saat data sedang dimuat: tampilkan skeleton atau spinner.
- Jika tidak ada hasil search/filter: pesan "Tidak ada pengguna yang sesuai." atau serupa.
- Setelah aksi Delete dari list (jika ada): baris hilang atau status diperbarui; pesan sukses singkat.

---

### 3.2 Detail Pengguna

**Deskripsi**: Halaman atau panel yang menampilkan informasi lengkap satu pengguna (read-only atau dengan tombol Edit/Delete).

**Elemen yang Harus Ada**:
- **Judul**: "Detail Pengguna" atau nama pengguna yang ditampilkan.
- **Informasi yang ditampilkan**: Nama lengkap, Email/Username, **Role(s)** (daftar semua role pengguna — satu user bisa punya lebih dari satu role), Satker/Unit Kerja (jika ada), Status (Aktif/Nonaktif), tanggal dibuat/diubah (opsional).
- **Tombol "Edit"**: Navigasi ke form Edit untuk pengguna ini.
- **Tombol "Hapus"** (atau "Nonaktifkan"): Memicu alur Delete dengan konfirmasi.
- **Tombol "Kembali"** atau link ke "Daftar Pengguna": Kembali ke List.

**Tata Letak dan Keterbacaan**:
- Informasi tersusun rapi (misalnya label di kiri, nilai di kanan; atau daftar vertikal).
- Tombol Edit dan Hapus terlihat jelas dan tidak membingungkan.

**Umpan Balik**:
- Saat data detail dimuat: spinner atau skeleton.
- Jika pengguna tidak ditemukan (misalnya dihapus orang lain): pesan "Pengguna tidak ditemukan." dan opsi kembali ke List.

---

### 3.3 Create Pengguna (Form Tambah Pengguna)

**Deskripsi**: Form untuk menambah pengguna baru.

**Elemen yang Harus Ada**:
- **Judul**: "Tambah Pengguna" atau "Buat Pengguna Baru".
- **Field wajib** (sesuai kebijakan sistem): Nama lengkap, Email atau Username (unik), Kata sandi, Konfirmasi kata sandi, **Role(s)** — **pilih satu atau lebih role** (multi-select: checkbox, dropdown multi, atau tag), Satker/Unit Kerja (jika berlaku untuk role dan wewenang).
- **Field opsional** (jika ada): No. HP, dll.
- **Keterangan role Widyaprada**: Jika **salah satu** role yang dipilih adalah **Widyaprada**, proses Create pengguna harus sekaligus membuat data Widyaprada (profil/entitas Widyaprada) untuk pengguna tersebut. Tanpa data Widyaprada, pengguna dengan role Widyaprada tidak dianggap lengkap.
- **Tombol "Simpan"** (atau "Buat Pengguna"): Submit form.
- **Tombol "Batal"** atau "Kembali": Kembali ke List tanpa menyimpan.
- **Opsi tampilkan/sembunyikan kata sandi** (ikon mata) pada field Kata sandi dan Konfirmasi kata sandi.

**Tata Letak dan Keterbacaan**:
- Label jelas; placeholder tidak menggantikan label.
- Pesan validasi tampil di dekat field yang salah (misalnya "Email sudah digunakan", "Kata sandi minimal 8 karakter").
- Pilihan **Role(s)** (multi-select) dan Satker hanya menampilkan opsi yang diizinkan untuk role pengguna yang login (Admin Satker hanya satker sendiri; Super Admin semua). Minimal satu role harus dipilih.

**Umpan Balik**:
- Saat submit: tombol Simpan menampilkan loading ("Menyimpan…" atau spinner) dan disabled sampai selesai.
- **Sukses**: Pesan "Pengguna berhasil ditambahkan." dan redirect ke List atau ke Detail pengguna baru.
- **Gagal**: Pesan error di atas form atau per field (misalnya "Email sudah terdaftar", "Isian tidak valid") — ramah dan dapat ditindaklanjuti.

---

### 3.4 Edit Pengguna (Form Ubah Data Pengguna)

**Deskripsi**: Form untuk mengubah data pengguna yang sudah ada.

**Elemen yang Harus Ada**:
- **Judul**: "Edit Pengguna" atau "Ubah Data Pengguna" (dapat disertai nama pengguna).
- **Field**: Nama lengkap, Email atau Username (unik, boleh readonly jika kebijakan tidak mengizinkan ubah), **Role(s)** (satu atau lebih role, multi-select — dapat menambah atau mengurangi role), Satker/Unit Kerja (jika berlaku), Status (Aktif/Nonaktif).
- **Field kata sandi (opsional)**: "Kata sandi baru" dan "Konfirmasi kata sandi baru" — jika kosong dianggap tidak mengubah kata sandi; jika diisi maka divalidasi dan di-update.
- **Tombol "Simpan"** (atau "Perbarui"): Submit form.
- **Tombol "Batal"** atau "Kembali": Kembali ke Detail atau List tanpa menyimpan.
- **Opsi tampilkan/sembunyikan** pada field kata sandi jika ada.

**Tata Letak dan Keterbacaan**:
- Sama seperti Create: label jelas, validasi per field, dropdown sesuai wewenang.

**Umpan Balik**:
- Saat submit: tombol Simpan loading dan disabled.
- **Sukses**: Pesan "Data pengguna berhasil diperbarui." dan redirect ke Detail atau List.
- **Gagal**: Pesan error yang jelas (misalnya "Email sudah digunakan oleh pengguna lain").

---

### 3.5 Delete Pengguna

**Deskripsi**: Aksi menghapus atau menonaktifkan pengguna, dengan konfirmasi dan **wajib mengisi alasan/deskripsi penghapusan**.

**Elemen yang Harus Ada**:
- **Pemicu**: Tombol "Hapus" atau "Nonaktifkan" di halaman Detail atau di kolom aksi List.
- **Dialog konfirmasi**: Sebelum aksi benar-benar dijalankan, tampilkan dialog (modal) dengan pesan jelas, misalnya: "Yakin ingin menghapus pengguna [Nama]? Tindakan ini tidak dapat dibatalkan." (atau "Akun ini akan dinonaktifkan." jika kebijakan nonaktifkan bukan hard delete).
- **Field wajib: Alasan/deskripsi penghapusan**: Dalam dialog konfirmasi harus ada field teks (textarea atau input) yang **wajib diisi** oleh pengelola — deskripsi atau alasan kenapa data/pengguna ini dihapus (misalnya: "Pindah satker", "Tidak lagi berwenang", "Duplikat data"). Tanpa mengisi alasan, tombol "Hapus" / "Ya, Hapus" tidak aktif atau submit ditolak dengan pesan "Silakan isi alasan penghapusan."
- **Tombol di dialog**: "Batal" (tutup dialog, tidak hapus) dan "Hapus" / "Ya, Hapus" (warna peringatan jika perlu); tombol Hapus hanya bisa diklik setelah alasan diisi.
- **Loading**: Saat proses hapus berjalan, tombol "Hapus" dalam dialog menampilkan loading agar pengguna tidak klik dua kali.

**Umpan Balik**:
- **Sukses**: Dialog tertutup; pesan singkat "Pengguna berhasil dihapus." atau "Pengguna berhasil dinonaktifkan."; list atau detail diperbarui (pengguna hilang dari list atau status jadi nonaktif).
- **Gagal**: Pesan error dalam dialog atau toast, misalnya "Gagal menghapus pengguna. Silakan coba lagi." atau "Pengguna tidak dapat dihapus karena [alasan]."

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur List
1. Pengelola (Admin Satker / Super Admin) membuka menu Manajemen Pengguna.
2. Melihat daftar pengguna (dengan loading jika data di-fetch).
3. Dapat mencari (ketik di kotak search), memfilter (role/satker/status), mengurutkan (klik header kolom).
4. Dapat klik "Tambah Pengguna" untuk ke form Create, atau klik baris/Detail untuk ke Detail, atau klik Edit/Hapus dari list (jika didukung).

### 4.2 Alur Detail
1. Dari List, pengelola klik satu pengguna (baris atau tombol Detail).
2. Halaman Detail menampilkan informasi lengkap pengguna.
3. Dapat klik "Edit" untuk ke form Edit, atau "Hapus" untuk memicu konfirmasi Delete.
4. Dapat klik "Kembali" untuk ke List.

### 4.3 Alur Create
1. Dari List, pengelola klik "Tambah Pengguna".
2. Form Create tampil; pengelola mengisi field wajib (nama, email/username, kata sandi, konfirmasi kata sandi, role, satker jika berlaku).
3. Klik "Simpan" → loading → sukses: redirect ke List atau Detail; gagal: pesan error tampil, form tetap bisa diperbaiki.

### 4.4 Alur Edit
1. Dari List atau Detail, pengelola klik "Edit".
2. Form Edit tampil dengan data terisi; pengelola mengubah field yang perlu; kata sandi baru opsional.
3. Klik "Simpan" → loading → sukses: redirect ke Detail atau List; gagal: pesan error tampil.

### 4.5 Alur Delete
1. Dari List atau Detail, pengelola klik "Hapus" (atau "Nonaktifkan").
2. Dialog konfirmasi tampil dengan pesan jelas dan **field wajib: Alasan/deskripsi kenapa data ini dihapus**.
3. Pengelola mengisi alasan penghapusan (wajib); jika tidak diisi, tombol "Ya, Hapus" tidak aktif atau sistem meminta isi alasan.
4. Klik "Batal" → dialog tutup, tidak ada perubahan. Setelah alasan diisi, klik "Ya, Hapus" → loading → sukses: dialog tutup, pesan sukses, list/detail diperbarui; gagal: pesan error.

---

## 5. Kebutuhan per Role

### 5.1 Role: Admin Satker

**Siapa**: Admin satuan kerja yang mengelola konten dan data di satker/unit kerjanya; akses Manajemen Pengguna jika role-nya diizinkan.

**Cakupan Manajemen Pengguna**:
- **List**: Hanya melihat pengguna yang terkait dengan satker/unit kerja/lembaga/instansi Admin Satker tersebut. Search, filter, sort hanya dalam scope tersebut.
- **Detail**: Hanya bisa membuka detail pengguna dalam satker mereka.
- **Create**: Hanya bisa menambah pengguna dengan Satker/Unit Kerja yang sama dengan satker Admin Satker (dropdown Satker hanya menampilkan satker sendiri; daftar role yang bisa dipilih mungkin dibatasi).
- **Edit**: Hanya bisa mengedit pengguna dalam satker mereka.
- **Delete**: Hanya bisa menghapus/nonaktifkan pengguna dalam satker mereka.

**UI/UX yang diharapkan**:
- Tidak ada opsi memilih satker lain saat Create; tidak ada baris pengguna dari satker lain di List.
- Pesan yang jelas jika secara tidak sengaja mengakses URL detail pengguna dari satker lain (misalnya "Anda tidak memiliki wewenang untuk mengakses pengguna ini.").

---

### 5.2 Role: Super Admin

**Siapa**: Pengguna dengan wewenang penuh untuk mengelola seluruh sistem, termasuk semua pengguna.

**Cakupan Manajemen Pengguna**:
- **List**: Melihat semua pengguna di sistem. Filter by satker, role, status tersedia.
- **Detail**: Bisa membuka detail pengguna mana pun.
- **Create**: Bisa menambah pengguna dengan satu atau lebih Role dan Satker mana pun (multi-select role, dropdown satker penuh).
- **Edit**: Bisa mengedit pengguna mana pun, termasuk mengubah role(s) dan satker.
- **Delete**: Bisa menghapus/nonaktifkan pengguna mana pun (dengan kebijakan: misalnya tidak boleh hapus diri sendiri).

**UI/UX yang diharapkan**:
- Semua filter dan dropdown tidak dibatasi oleh satker. Navigasi dan aksi konsisten dengan "kendali penuh".

---

## 6. Ringkasan Perbedaan per Role

| Aspek | Admin Satker | Super Admin |
|-------|--------------|-------------|
| **List** | Hanya pengguna dalam satker sendiri | Semua pengguna |
| **Detail** | Hanya pengguna dalam satker sendiri | Semua pengguna |
| **Create** | Hanya bisa pilih satker sendiri; role(s) mungkin dibatasi | Bisa pilih satu atau lebih role dan satker mana pun |
| **Edit** | Hanya pengguna dalam satker sendiri | Semua pengguna |
| **Delete** | Hanya pengguna dalam satker sendiri | Semua pengguna (dengan kebijakan, misalnya tidak hapus diri sendiri) |
| **Pengalaman** | Fokus pada satker sendiri | Kendali penuh seluruh sistem |

---

## 7. Cakupan Fitur Manajemen Pengguna

### 7.1 Termasuk
- **List**: Daftar pengguna dengan search, filter (role, satker, status), sort, paginasi; tombol Tambah, Detail, Edit, Hapus sesuai wewenang.
- **Detail**: Halaman detail satu pengguna dengan tombol Edit dan Hapus.
- **Create**: Form tambah pengguna (nama, email/username, kata sandi, **satu atau lebih role**, satker) dengan validasi dan umpan balik. Jika salah satu role yang dipilih adalah **Widyaprada**, Create pengguna wajib disertai pembuatan data Widyaprada (entitas/profil Widyaprada) agar pengguna tersebut dapat berfungsi penuh.
- **Edit**: Form ubah data pengguna (termasuk kata sandi opsional) dengan validasi dan umpan balik.
- **Delete**: Konfirmasi sebelum hapus/nonaktifkan; **wajib prompt field alasan/deskripsi kenapa data ini dihapus** (diisi dalam dialog konfirmasi); pesan sukses/error; pembatasan wewenang (Admin Satker hanya satker sendiri; Super Admin dengan kebijakan misalnya tidak hapus diri sendiri).
- **Wewenang**: Scope data dan aksi sesuai role (Admin Satker vs Super Admin) seperti di atas.

### 7.2 Tidak Termasuk (Fitur Lain)
- Login, Logout, Lupa Password → PRD masing-masing (Auth).
- Manajemen Role dan Permission → PRD modul Role & Permission.
- Registrasi mandiri oleh pengguna (jika ada) → PRD terpisah.
- Audit log siapa mengubah siapa (jika ada) → dapat didokumentasikan di PRD terpisah atau tambahan.

---

## 8. Persyaratan Produk (Nonteknis)

Detail implementasi (validasi wewenang, kebijakan kata sandi, schema, API) didokumentasikan di SDD.

- **Wewenang:** Admin Satker hanya boleh melihat dan mengubah pengguna yang berada dalam satker/unit kerjanya. Jika mencoba mengakses pengguna satker lain, sistem menampilkan pesan jelas (misalnya "Anda tidak memiliki wewenang untuk mengakses pengguna ini.").
- **Kata sandi:** Saat Create dan Edit, kebijakan kata sandi (panjang minimal, kompleksitas) diterapkan; konfirmasi kata sandi harus sama dengan kata sandi. Pesan error tampil jika tidak memenuhi.
- **Unik:** Email dan username harus unik. Jika sudah terdaftar, tampilkan pesan ramah (misalnya "Email sudah terdaftar") tanpa membocorkan data pengguna lain.
- **Role Widyaprada:** Saat Create pengguna yang salah satu role-nya Widyaprada, sistem wajib sekaligus membuat data Widyaprada (profil/entitas Widyaprada) yang terhubung ke pengguna tersebut. Create pengguna dianggap berhasil hanya jika pengguna (dan data Widyaprada jika role Widyaprada dipilih) berhasil dibuat.
- **Multi-role:** Satu pengguna dapat memiliki lebih dari satu role. Relasi user–role bersifat many-to-many. Menu dan wewenang setelah login adalah gabungan dari semua permission role yang dimiliki.
- **Delete:** Alasan penghapusan wajib diisi di dialog konfirmasi; tanpa alasan, tombol "Ya, Hapus" tidak aktif atau submit ditolak. Sistem menyimpan alasan tersebut. Jika kebijakan memakai soft delete (nonaktif), list dan filter status harus mendukung "Nonaktif".
- **Umpan balik:** List menampilkan paginasi jika data banyak; setiap aksi (Create, Edit, Delete) menampilkan loading lalu pesan sukses atau error.
- **Akses perangkat:** Halaman List, Detail, dan form Create/Edit dapat dipakai dengan nyaman dari desktop dan tablet (layout responsif).

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-01 | Dokumen awal: List, Detail, Create, Edit, Delete; fokus kebutuhan user & UI/UX per role | - |
| 1.1 | 2025-02-11 | Satu user bisa lebih dari satu role; form Create/Edit: multi-select role; list/detail tampilkan daftar role | - |

---

**Catatan**: Dokumen ini mencakup fitur **Manajemen Pengguna** (List, Detail, Create, Edit, Delete). Fitur **Login**, **Logout**, **Lupa Password**, dan **Role & Permission** didokumentasikan dalam PRD masing-masing.
