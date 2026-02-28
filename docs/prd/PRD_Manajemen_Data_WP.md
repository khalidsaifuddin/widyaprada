# [PRD] Fitur Manajemen Data WP
## Product Requirements Document | Manajemen Data Widyaprada

**Aplikasi**: Widyaprada  
**Modul**: Manajemen Data WP  
**Fitur**: Manajemen Data Widyaprada (List, Detail, Create, Edit, Delete)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Manajemen Data WP (Manajemen Data Widyaprada)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Manajemen Data WP memungkinkan Admin Satker dan Super Admin untuk melihat daftar data Widyaprada (entitas/profil Widyaprada), daftar **calon peserta** yang telah apply beserta dokumen persyaratan, melakukan **verifikasi dan validasi dokumen persyaratan**, serta mengelola data WP (CRUD). Admin Satker hanya mengelola data WP dalam satker/unit kerjanya; Super Admin mengelola seluruh data WP. **Tim Verval** memeriksa dokumen persyaratan calon; jika ditolak ada catatan (tanpa koreksi — langsung info tidak lulus syarat administrasi).

### 1.3 Field/Kolom Data Widyaprada (Referensi Data Kepegawaian PNS Indonesia)

Struktur field berikut mengacu pada praktik data kepegawaian PNS/ASN Indonesia (SAPK BKN, DUK, portal data kepegawaian) dan regulasi jabatan fungsional Widyaprada (Permendikbud 37/2020, Permenpanrb 93/2020).

| No | Field / Kolom | Wajib | Tipe | Keterangan / Referensi |
|----|----------------|-------|------|-------------------------|
| 1 | **NIP** | Ya | Teks (18 digit) | Nomor Induk Pegawai — unik, format standar BKN. Digunakan untuk identifikasi dan pencarian. |
| 2 | **Nama Lengkap** | Ya | Teks | Nama lengkap PNS sesuai dokumen kepegawaian. |
| 3 | **Jenis Kelamin** | Opsional | Pilihan (L/P) | Laki-laki / Perempuan — lazim di data statistik kepegawaian. |
| 4 | **Golongan Ruang** | Opsional | Pilihan | Golongan I s.d. IV dengan sub ruang (III/a, III/b, III/c, III/d, IV/a … IV/e). Standar kepangkatan PNS. |
| 5 | **Pangkat** | Opsional | Teks/Pilihan | Contoh: Penata Muda III/a, Penata III/c, Pembina IV/a. Mengikuti golongan ruang. |
| 6 | **Jenjang Jabatan Fungsional** | Opsional | Pilihan | Widyaprada Ahli Pertama / Ahli Muda / Ahli Madya / Ahli Utama (Permendikbud 37/2020). |
| 7 | **Satuan Kerja (Satker)** | Ya | Pilihan/Referensi | Unit organisasi tempat pegawai bertugas (Eselon I/II atau setara). Scope wewenang Admin Satker. |
| 8 | **Unit Kerja** | Opsional | Teks/Pilihan | Sub unit di bawah satker (direktorat, bidang, dll). |
| 9 | **Pendidikan Terakhir** | Opsional | Pilihan | S1, S2, S3, D4, D3, dll — umum di DUK dan profil ASN. |
| 10 | **TMT Golongan** | Opsional | Tanggal | TMT = Tanggal Mulai Tercatat. TMT kepangkatan/golongan. |
| 11 | **TMT Jabatan Fungsional** | Opsional | Tanggal | Tanggal mulai menjabat sebagai Widyaprada (jabatan fungsional). |
| 12 | **No. SK Pengangkatan** | Opsional | Teks | Nomor SK pengangkatan dalam jabatan fungsional Widyaprada. |
| 13 | **No. HP** | Opsional | Teks | Nomor telepon/HP untuk keperluan operasional dan kontak. |
| 14 | **Email** | Opsional | Teks | Email institusi/pribadi — untuk notifikasi dan login jika terhubung ke user. |
| 15 | **Alamat** | Opsional | Teks panjang | Alamat domisili atau kantor jika diperlukan. |
| 16 | **Status** | Ya | Pilihan | Aktif / Nonaktif (atau Cuti/Pensiun jika kebijakan mengakomodasi). Untuk filter dan tampilan list. |
| 17 | **Keterangan** | Opsional | Teks panjang | Catatan tambahan (mis. wilayah binaan, tugas khusus). |
| 18 | **Tanggal Dibuat** | Sistem | Timestamp | Diisi otomatis oleh sistem. |
| 19 | **Tanggal Diubah** | Sistem | Timestamp | Diisi otomatis saat data diperbarui. |

**Referensi singkat:**
- **SAPK BKN / MyASN**: Data Utama, Golongan, Jabatan, Posisi, Pendidikan, Data Pribadi (profil ASN).
- **DUK (Daftar Urut Kepangkatan)**: Penyusunan berdasarkan Pangkat, Jabatan, Masa Kerja, Pendidikan, Latihan Jabatan, Usia.
- **Jabatan Fungsional Widyaprada**: Permendikbud 37/2020 (Juknis), Permenpanrb 93/2020 (Standar Kompetensi). Jenjang: Ahli Pertama (III/a–III/b), Ahli Muda (III/c–III/d), Ahli Madya (IV/a–IV/c), Ahli Utama (IV/d–IV/e).

**Catatan implementasi:**
- Field **wajib minimal** untuk Create: NIP, Nama Lengkap, Satker, Status. Field lain dapat diwajibkan sesuai kebijakan satker.
- **NIP** harus unik di dalam sistem; validasi format 18 digit sesuai ketentuan BKN (jika diterapkan).
- Jika Data WP terhubung ke **akun pengguna** (user role Widyaprada), tambahkan field/relasi **User ID** (referensi ke tabel user) agar satu akun terhubung ke satu data WP.

### 1.4 Data Historis dan Relasi One-to-Many (Satu Widyaprada — Banyak Data Terkait)

Satu entitas **Data Widyaprada** (satu orang Widyaprada, satu NIP) dapat memiliki **banyak data terkait** yang bersifat **historis** atau **one-to-many**. Artinya: satu Widyaprada bisa punya lebih dari satu record untuk riwayat atau untuk data yang mengulang (misalnya banyak anggota keluarga, banyak periode gaji). Bagian ini memberi keterangan agar model data dan fitur bisa dirancang dengan relasi yang benar.

#### 1.4.1 Data yang Sifatnya Historis

**Pengertian:** Data historis adalah data yang mencatat **perubahan atau kejadian di waktu lampau** untuk satu orang Widyaprada yang sama. Satu Widyaprada (satu NIP) memiliki **banyak record historis**; setiap record punya periode/waktu (misalnya tanggal mulai–selesai atau tanggal efektif) dan tidak menimpa record lama, melainkan menambah riwayat.

| Contoh data historis | Keterangan | Relasi |
|---------------------|------------|--------|
| **Riwayat kepangkatan / golongan** | Setiap kenaikan pangkat atau perubahan golongan ruang (III/a → III/b, dll) dicatat sebagai satu baris dengan TMT, pangkat, golongan, no. SK. | Satu WP → banyak riwayat pangkat. |
| **Riwayat jabatan** | Perubahan jabatan (struktural/fungsional), mutasi jabatan, TMT jabatan. | Satu WP → banyak riwayat jabatan. |
| **Riwayat satker / unit kerja** | Mutasi atau pindah satker/unit kerja; tiap pindah dicatat dengan periode dan satker baru. | Satu WP → banyak riwayat penempatan. |
| **Riwayat gaji berkala** | Setiap periode kenaikan gaji (berkala) dicatat: periode, gaji pokok, no. SK, tanggal. | Satu WP → banyak riwayat gaji. |
| **Riwayat diklat / pelatihan** | Setiap diklat atau pelatihan yang diikuti (nama, tahun, jam, institusi). | Satu WP → banyak riwayat diklat. |
| **Riwayat penghargaan** | Tanda jasa, satyalancana, atau penghargaan lain per peristiwa. | Satu WP → banyak riwayat penghargaan. |

**Implikasi:**  
- Di **tabel master** Data WP (1.3), field seperti Golongan Ruang, Pangkat, Satker, TMT bisa menyimpan **nilai saat ini (current)**.  
- **Riwayat lengkap** disimpan di tabel terpisah (one-to-many ke Data WP), dengan **foreign key ke Data WP** (atau NIP).  
- Fitur **Detail** Data WP dapat menampilkan sub-section “Riwayat …” (riwayat pangkat, riwayat jabatan, riwayat gaji, dll) dalam bentuk list/tabel atau tab; Create/Edit riwayat bisa dari halaman Detail atau modul terpisah (lihat cakupan fitur 7).

#### 1.4.2 Data yang Sifatnya One-to-Many (Bukan Hanya Historis)

**Pengertian:** Data yang **satu orang punya banyak** record, tidak selalu berurutan waktu; bisa berupa daftar anggota atau daftar item yang terkait satu Widyaprada.

| Contoh data one-to-many | Keterangan | Relasi |
|-------------------------|------------|--------|
| **Data anggota keluarga** | Istri/suami, anak, orang tua, dll. Setiap anggota: nama, hubungan, tanggal lahir, NIK, pekerjaan, dll. | Satu WP → banyak anggota keluarga. |
| **Riwayat gaji berkala** | Sudah dicontohkan di atas; sekaligus historis dan one-to-many (banyak periode gaji). | Satu WP → banyak record gaji. |
| **Data dependen / tanggungan** | Orang yang menjadi tanggungan (untuk tunjangan, asuransi). | Satu WP → banyak tanggungan. |

**Implikasi:**  
- Disimpan di **tabel terpisah** dengan **foreign key ke Data WP** (atau NIP).  
- Di halaman **Detail** Data WP dapat ditampilkan sub-list “Anggota Keluarga”, “Riwayat Gaji”, dll, dengan aksi Tambah/Edit/Hapus per item (atau diatur di PRD/scope terpisah).

#### 1.4.3 Ringkasan Relasi

- **Master:** Satu record **Data Widyaprada** per orang (identifikasi unik NIP). Field di 1.3 menggambarkan data **master dan nilai saat ini**.
- **Historis:** Satu Widyaprada → **banyak** record riwayat (pangkat, jabatan, satker, gaji, diklat, penghargaan). Setiap record punya konteks waktu/periode dan referensi ke Data WP.
- **One-to-many:** Satu Widyaprada → **banyak** record untuk data seperti anggota keluarga atau tanggungan; bisa dikelola dari Detail Data WP atau modul terkait.

Apakah **pengelolaan** data historis dan one-to-many (Create/Edit/Delete per riwayat atau per anggota) masuk dalam cakupan fitur **Manajemen Data WP** atau menjadi modul/PRD terpisah, ditentukan di **Cakupan Fitur (bagian 7)** dan backlog.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur Manajemen Data WP dirumuskan per role dalam pola user story berikut, dalam format tabular.

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Admin Satker / Super Admin | pengelola data WP | melihat **daftar calon peserta** (yang telah apply dengan dokumen persyaratan) | memantau dan memverifikasi pendaftaran calon |
| 2 | Admin Satker / Super Admin / Tim Verval | verifikator | memeriksa dan **verifikasi serta validasi dokumen persyaratan** yang diupload calon | menyeleksi peserta yang lolos syarat administrasi |
| 3 | Admin Satker / Super Admin | pengelola | menolak pendaftaran dengan **catatan** (tanpa koreksi — langsung info tidak lulus syarat administrasi) | calon tahu alasan tidak lolos |
| 4 | Admin Satker / Super Admin | pengelola data WP | melihat daftar data Widyaprada dalam satu halaman/list yang terstruktur (nama, NIP, satker, unit kerja, status, dll) | dengan cepat menemukan dan memantau data WP yang saya kelola |
| 5 | Admin Satker / Super Admin | pengelola data WP | mencari data WP berdasarkan nama, NIP, atau satker/unit kerja | menemukan data WP tertentu tanpa harus menggulir panjang |
| 6 | Admin Satker / Super Admin | pengelola data WP | memfilter daftar berdasarkan satker, unit kerja, atau status (aktif/nonaktif) | fokus pada kelompok data WP yang relevan dengan tugas saya |
| 7 | Admin Satker / Super Admin | pengelola data WP | mengurutkan daftar (misalnya nama, NIP, tanggal dibuat) | mengorganisir tampilan sesuai kebutuhan |
| 8 | Admin Satker / Super Admin | pengelola data WP | membuka detail satu data WP atau detail calon peserta (dokumen persyaratan) dari list | melihat informasi lengkap sebelum mengedit atau memverifikasi |
| 9 | Admin Satker / Super Admin | pengelola data WP | menambah data WP baru melalui form Create yang jelas (nama, NIP, satker, unit kerja, kontak, dll) | mendaftarkan data Widyaprada baru ke sistem sesuai wewenang saya |
| 10 | Admin Satker / Super Admin | pengelola data WP | mengubah data WP yang ada (Edit) dari halaman detail atau dari list | memperbarui informasi tanpa harus menghapus dan buat ulang |
| 11 | Admin Satker / Super Admin | pengelola data WP | menghapus data WP (Delete) dengan konfirmasi yang jelas | mencabut atau membersihkan data yang tidak lagi valid atau salah input |
| 12 | Admin Satker | Admin Satker | hanya melihat dan mengelola data WP dalam satker/unit kerja/lembaga/instansi saya | tidak mengakses atau mengubah data WP di satker lain |
| 13 | Super Admin | Super Admin | melihat dan mengelola semua data WP di sistem | mengawasi dan mengkonfigurasi seluruh data Widyaprada |
| 14 | Admin Satker / Super Admin | pengelola data WP | mendapat umpan balik jelas setelah Create/Edit/Delete (sukses atau pesan error) | tahu apakah aksi berhasil dan apa yang harus dilakukan jika gagal |
| 15 | Admin Satker / Super Admin | pengelola data WP | paginasi atau lazy load pada list jika data banyak | daftar tetap cepat dan nyaman dipakai |
| 16 | Admin Satker / Super Admin | pengelola data WP | tombol/aksi (Create, Edit, Delete) menampilkan loading saat diproses | yakin bahwa aplikasi sedang bekerja dan tidak hang |
| 17 | Admin Satker / Super Admin | pengelola data WP | saat Delete, wajib mengisi alasan/deskripsi kenapa data ini dihapus dalam prompt dialog konfirmasi | ada jejak dokumentasi alasan penghapusan dan mengurangi hapus tidak sengaja |

**Keterangan**: *Admin Satker* = akses Manajemen Data WP terbatas pada satker/unit kerjanya; *Super Admin* = akses penuh ke semua data WP. Data WP = entitas/profil Widyaprada (terhubung ke pengguna role Widyaprada jika ada; atau data mandiri tergantung kebijakan sistem).

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **List**: Daftar data WP yang terstruktur dengan kolom relevan (nama, NIP, satker, unit kerja, status), search, filter, sort, dan paginasi.
- **Detail**: Satu halaman/view untuk melihat informasi lengkap satu data WP sebelum Edit atau Delete.
- **Create**: Form yang jelas untuk menambah data WP baru dengan validasi (NIP unik jika berlaku, satker/unit kerja sesuai wewenang).
- **Edit**: Form untuk mengubah data WP dengan validasi yang sama.
- **Delete**: Aksi hapus dengan konfirmasi; **wajib ada prompt untuk mengisi deskripsi/alasan kenapa data ini dihapus** (field teks dalam dialog konfirmasi). Setelah konfirmasi dan alasan diisi, baru pesan sukses/error ditampilkan.
- **Umpan balik**: Loading pada tombol/aksi; pesan sukses atau error yang ramah dan jelas setelah setiap aksi.
- **Wewenang**: Admin Satker hanya mengelola data WP dalam satker/unit kerjanya; Super Admin mengelola semua.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- List tanpa search/filter sehingga sulit menemukan data WP saat data banyak.
- Form Create/Edit yang tidak memberi tahu field mana yang salah (validasi tidak jelas).
- Delete tanpa konfirmasi sehingga risiko hapus tidak sengaja.
- Delete tanpa wajib mengisi alasan/deskripsi penghapusan sehingga tidak ada jejak dokumentasi.
- Admin Satker bisa melihat atau mengubah data WP di satker lain.
- Pesan error yang teknis atau membingungkan.

---

## 3. Antarmuka Pengguna (UI)

*Manajemen Data WP terdiri atas: halaman List, halaman/view Detail, form Create, form Edit, dan aksi Delete (biasanya dari Detail atau dari list). Semua mengikuti design system aplikasi.*

---

### 3.1 List Data WP

**Deskripsi**: Halaman yang menampilkan daftar data Widyaprada dalam bentuk tabel atau card list.

**Elemen yang Harus Ada**:
- **Judul halaman**: Misalnya "Manajemen Data WP" atau "Daftar Data Widyaprada".
- **Tombol "Tambah Data WP"** (atau "Create"): Memicu navigasi ke form Create.
- **Kotak pencarian**: Untuk mencari berdasarkan nama, NIP, atau satker/unit kerja (placeholder jelas).
- **Filter** (opsional tapi disarankan): Dropdown atau pilihan untuk filter berdasarkan satker, unit kerja, status (aktif/nonaktif).
- **Tabel/List**: Kolom mengacu pada field di **1.3** — minimal tampil: Nama Lengkap, NIP, Satker, Unit Kerja (opsional), Status. Dapat ditambah Pangkat/Jenjang Jabatan jika diinginkan. Kolom aksi: tombol atau link "Detail", "Edit", "Hapus" (sesuai wewenang).
- **Paginasi**: Jika data banyak, tampilkan paginasi (nomor halaman atau "Load more") agar performa tetap baik.
- **Sort**: Header kolom dapat diklik untuk mengurutkan (nama, NIP, tanggal, dll) jika diinginkan.

**Tata Letak dan Keterbacaan**:
- Tabel/list rapi; jarak antar baris nyaman; teks tidak terpotong sembarangan.
- Tombol aksi (Detail, Edit, Hapus) konsisten per baris dan mudah diklik (termasuk di layar sentuh).
- Di layar kecil (ponsel/tablet), list dapat di-responsive (misalnya card per data WP atau tabel scroll horizontal).

**Umpan Balik**:
- Saat data sedang dimuat: tampilkan skeleton atau spinner.
- Jika tidak ada hasil search/filter: pesan "Tidak ada data WP yang sesuai." atau serupa.
- Setelah aksi Delete dari list (jika ada): baris hilang atau status diperbarui; pesan sukses singkat.

---

### 3.2 Detail Data WP

**Deskripsi**: Halaman atau panel yang menampilkan informasi lengkap satu data Widyaprada (read-only atau dengan tombol Edit/Delete).

**Elemen yang Harus Ada**:
- **Judul**: "Detail Data WP" atau nama Widyaprada yang ditampilkan.
- **Informasi yang ditampilkan**: Sesuai field di **1.3** — Nama Lengkap, NIP, Jenis Kelamin, Golongan Ruang, Pangkat, Jenjang Jabatan Fungsional, Satker, Unit Kerja, Pendidikan Terakhir, TMT Golongan/TMT Jabatan, No. SK (jika ada), No. HP, Email, Alamat, Status, Keterangan, Tanggal Dibuat/Diubah.
- **Data historis & one-to-many (1.4)**: Halaman Detail dapat menampilkan sub-section atau tab untuk **data historis** (riwayat pangkat, riwayat jabatan, riwayat gaji berkala, riwayat diklat, dll) dan **data one-to-many** (anggota keluarga, tanggungan) dalam bentuk list/tabel; aksi Tambah/Edit/Hapus per item mengikuti keputusan cakupan fitur (7.1).
- **Tombol "Edit"**: Navigasi ke form Edit untuk data WP ini.
- **Tombol "Hapus"**: Memicu alur Delete dengan konfirmasi.
- **Tombol "Kembali"** atau link ke "Daftar Data WP": Kembali ke List.

**Tata Letak dan Keterbacaan**:
- Informasi tersusun rapi (misalnya label di kiri, nilai di kanan; atau daftar vertikal).
- Tombol Edit dan Hapus terlihat jelas dan tidak membingungkan.

**Umpan Balik**:
- Saat data detail dimuat: spinner atau skeleton.
- Jika data WP tidak ditemukan (misalnya dihapus orang lain): pesan "Data WP tidak ditemukan." dan opsi kembali ke List.

---

### 3.3 Create Data WP (Form Tambah Data WP)

**Deskripsi**: Form untuk menambah data Widyaprada baru.

**Elemen yang Harus Ada**:
- **Judul**: "Tambah Data WP" atau "Buat Data Widyaprada Baru".
- **Field** mengacu pada **1.3**: wajib minimal — Nama Lengkap, NIP, Satker, Status; opsional — Jenis Kelamin, Golongan Ruang, Pangkat, Jenjang Jabatan Fungsional, Unit Kerja, Pendidikan Terakhir, TMT Golongan, TMT Jabatan Fungsional, No. SK Pengangkatan, No. HP, Email, Alamat, Keterangan. Satker/Unit Kerja sesuai wewenang (Admin Satker hanya satker sendiri).
- **Tombol "Simpan"** (atau "Buat Data WP"): Submit form.
- **Tombol "Batal"** atau "Kembali": Kembali ke List tanpa menyimpan.

**Tata Letak dan Keterbacaan**:
- Label jelas; placeholder tidak menggantikan label.
- Pesan validasi tampil di dekat field yang salah (misalnya "NIP sudah digunakan", "Nama wajib diisi").
- Dropdown Satker/Unit Kerja hanya menampilkan opsi yang diizinkan (Admin Satker hanya satker sendiri; Super Admin semua).

**Umpan Balik**:
- Saat submit: tombol Simpan menampilkan loading ("Menyimpan…" atau spinner) dan disabled sampai selesai.
- **Sukses**: Pesan "Data WP berhasil ditambahkan." dan redirect ke List atau ke Detail data WP baru.
- **Gagal**: Pesan error di atas form atau per field (misalnya "NIP sudah terdaftar", "Isian tidak valid") — ramah dan dapat ditindaklanjuti.

---

### 3.4 Edit Data WP (Form Ubah Data WP)

**Deskripsi**: Form untuk mengubah data Widyaprada yang sudah ada.

**Elemen yang Harus Ada**:
- **Judul**: "Edit Data WP" atau "Ubah Data Widyaprada" (dapat disertai nama).
- **Field**: Sesuai **1.3** — Nama Lengkap, NIP (unik, boleh readonly jika kebijakan tidak mengizinkan ubah), Jenis Kelamin, Golongan Ruang, Pangkat, Jenjang Jabatan Fungsional, Satker, Unit Kerja (sesuai wewenang), Pendidikan Terakhir, TMT Golongan, TMT Jabatan Fungsional, No. SK, No. HP, Email, Alamat, Status, Keterangan.
- **Tombol "Simpan"** (atau "Perbarui"): Submit form.
- **Tombol "Batal"** atau "Kembali": Kembali ke Detail atau List tanpa menyimpan.

**Tata Letak dan Keterbacaan**:
- Sama seperti Create: label jelas, validasi per field, dropdown sesuai wewenang.

**Umpan Balik**:
- Saat submit: tombol Simpan loading dan disabled.
- **Sukses**: Pesan "Data WP berhasil diperbarui." dan redirect ke Detail atau List.
- **Gagal**: Pesan error yang jelas (misalnya "NIP sudah digunakan oleh data WP lain").

---

### 3.5 Delete Data WP

**Deskripsi**: Aksi menghapus data WP, dengan konfirmasi dan **wajib mengisi alasan/deskripsi penghapusan**.

**Elemen yang Harus Ada**:
- **Pemicu**: Tombol "Hapus" di halaman Detail atau di kolom aksi List.
- **Dialog konfirmasi**: Sebelum aksi benar-benar dijalankan, tampilkan dialog (modal) dengan pesan jelas, misalnya: "Yakin ingin menghapus data WP [Nama]? Tindakan ini tidak dapat dibatalkan."
- **Field wajib: Alasan/deskripsi penghapusan**: Dalam dialog konfirmasi harus ada field teks (textarea atau input) yang **wajib diisi** — deskripsi atau alasan kenapa data WP ini dihapus (misalnya: "Pindah satker", "Duplikat data", "Tidak lagi menjabat"). Tanpa mengisi alasan, tombol "Hapus" / "Ya, Hapus" tidak aktif atau submit ditolak dengan pesan "Silakan isi alasan penghapusan."
- **Tombol di dialog**: "Batal" (tutup dialog, tidak hapus) dan "Hapus" / "Ya, Hapus" (warna peringatan jika perlu); tombol Hapus hanya bisa diklik setelah alasan diisi.
- **Loading**: Saat proses hapus berjalan, tombol "Hapus" dalam dialog menampilkan loading agar pengguna tidak klik dua kali.

**Umpan Balik**:
- **Sukses**: Dialog tertutup; pesan singkat "Data WP berhasil dihapus."; list atau detail diperbarui (data WP hilang dari list).
- **Gagal**: Pesan error dalam dialog atau toast, misalnya "Gagal menghapus data WP. Silakan coba lagi." atau "Data WP tidak dapat dihapus karena [alasan]."

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur List
1. Pengelola (Admin Satker / Super Admin) membuka menu Manajemen Data WP.
2. Melihat daftar data WP (dengan loading jika data di-fetch).
3. Dapat mencari (ketik di kotak search), memfilter (satker/unit kerja/status), mengurutkan (klik header kolom).
4. Dapat klik "Tambah Data WP" untuk ke form Create, atau klik baris/Detail untuk ke Detail, atau klik Edit/Hapus dari list (jika didukung).

### 4.2 Alur Detail
1. Dari List, pengelola klik satu data WP (baris atau tombol Detail).
2. Halaman Detail menampilkan informasi lengkap data WP.
3. Dapat klik "Edit" untuk ke form Edit, atau "Hapus" untuk memicu konfirmasi Delete.
4. Dapat klik "Kembali" untuk ke List.

### 4.3 Alur Create
1. Dari List, pengelola klik "Tambah Data WP".
2. Form Create tampil; pengelola mengisi field wajib (nama, NIP, satker/unit kerja, dll).
3. Klik "Simpan" → loading → sukses: redirect ke List atau Detail; gagal: pesan error tampil, form tetap bisa diperbaiki.

### 4.4 Alur Edit
1. Dari List atau Detail, pengelola klik "Edit".
2. Form Edit tampil dengan data terisi; pengelola mengubah field yang perlu.
3. Klik "Simpan" → loading → sukses: redirect ke Detail atau List; gagal: pesan error tampil.

### 4.5 Alur Delete
1. Dari List atau Detail, pengelola klik "Hapus".
2. Dialog konfirmasi tampil dengan pesan jelas dan **field wajib: Alasan/deskripsi kenapa data ini dihapus**.
3. Pengelola mengisi alasan penghapusan (wajib); jika tidak diisi, tombol "Ya, Hapus" tidak aktif atau sistem meminta isi alasan.
4. Klik "Batal" → dialog tutup, tidak ada perubahan. Setelah alasan diisi, klik "Ya, Hapus" → loading → sukses: dialog tutup, pesan sukses, list/detail diperbarui; gagal: pesan error.

---

## 5. Kebutuhan per Role

### 5.1 Role: Admin Satker

**Siapa**: Admin satuan kerja yang mengelola konten dan data di satker/unit kerjanya, termasuk data Widyaprada di satker tersebut.

**Cakupan Manajemen Data WP**:
- **List**: Hanya melihat data WP yang terkait dengan satker/unit kerja/lembaga/instansi Admin Satker tersebut. Search, filter, sort hanya dalam scope tersebut.
- **Detail**: Hanya bisa membuka detail data WP dalam satker mereka.
- **Create**: Hanya bisa menambah data WP dengan Satker/Unit Kerja yang sama dengan satker Admin Satker (dropdown hanya menampilkan satker sendiri).
- **Edit**: Hanya bisa mengedit data WP dalam satker mereka.
- **Delete**: Hanya bisa menghapus data WP dalam satker mereka.

**UI/UX yang diharapkan**:
- Tidak ada opsi memilih satker lain saat Create; tidak ada baris data WP dari satker lain di List.
- Pesan yang jelas jika secara tidak sengaja mengakses URL detail data WP dari satker lain (misalnya "Anda tidak memiliki wewenang untuk mengakses data WP ini.").

---

### 5.2 Role: Super Admin

**Siapa**: Pengguna dengan wewenang penuh untuk mengelola seluruh sistem, termasuk semua data Widyaprada.

**Cakupan Manajemen Data WP**:
- **List**: Melihat semua data WP di sistem. Filter by satker, unit kerja, status tersedia.
- **Detail**: Bisa membuka detail data WP mana pun.
- **Create**: Bisa menambah data WP dengan Satker/Unit Kerja mana pun (dropdown penuh).
- **Edit**: Bisa mengedit data WP mana pun, termasuk mengubah satker/unit kerja.
- **Delete**: Bisa menghapus data WP mana pun (dengan kebijakan jika ada, misalnya validasi relasi ke pengguna/jurnal).

**UI/UX yang diharapkan**:
- Semua filter dan dropdown tidak dibatasi oleh satker. Navigasi dan aksi konsisten dengan "kendali penuh".

---

## 6. Ringkasan Perbedaan per Role

| Aspek | Admin Satker | Super Admin |
|-------|--------------|-------------|
| **List** | Hanya data WP dalam satker sendiri | Semua data WP |
| **Detail** | Hanya data WP dalam satker sendiri | Semua data WP |
| **Create** | Hanya bisa pilih satker sendiri | Bisa pilih satker mana pun |
| **Edit** | Hanya data WP dalam satker sendiri | Semua data WP |
| **Delete** | Hanya data WP dalam satker sendiri | Semua data WP (dengan kebijakan jika ada) |
| **Pengalaman** | Fokus pada satker sendiri | Kendali penuh seluruh sistem |

---

## 7. Cakupan Fitur Manajemen Data WP

### 7.1 Termasuk
- **Daftar calon peserta**: List calon yang telah apply beserta dokumen persyaratan; aksi verifikasi/validasi; jika ditolak: catatan (tanpa koreksi — langsung info tidak lulus syarat administrasi).
- **Verifikasi & validasi dokumen**: Tim Verval/Admin memeriksa kelengkapan berkas sesuai lampiran; setuju/tolak dengan catatan.
- **List**: Daftar data WP dengan search, filter (satker, unit kerja, status), sort, paginasi; tombol Tambah, Detail, Edit, Hapus sesuai wewenang.
- **Detail**: Halaman detail satu data WP dengan tombol Edit dan Hapus. **Data historis dan one-to-many** (lihat **1.4**) dapat ditampilkan di Detail sebagai sub-section/tab (mis. Riwayat Pangkat, Riwayat Gaji, Anggota Keluarga); apakah Create/Edit/Delete per riwayat atau per anggota keluarga masuk dalam fitur ini atau fase berikutnya ditentukan di backlog.
- **Create**: Form tambah data WP (nama, NIP, satker/unit kerja, kontak, dll) dengan validasi dan umpan balik.
- **Edit**: Form ubah data WP dengan validasi dan umpan balik.
- **Delete**: Konfirmasi sebelum hapus; **wajib prompt field alasan/deskripsi kenapa data ini dihapus** (diisi dalam dialog konfirmasi); pesan sukses/error; pembatasan wewenang (Admin Satker hanya satker sendiri; Super Admin penuh).
- **Wewenang**: Scope data dan aksi sesuai role (Admin Satker vs Super Admin) seperti di atas.

### 7.2 Tidak Termasuk (Fitur Lain)
- Manajemen Pengguna (akun, role) → PRD Auth & Manajemen Pengguna.
- WPData, WPJurnal, WPUjikom (modul lain) → PRD masing-masing.
- Pembuatan otomatis data WP saat Create pengguna role Widyaprada → tercakup di PRD Manajemen Pengguna; integrasi teknis tetap mengacu ke sana jika data WP terhubung ke user.
- Audit log perubahan data WP (jika ada) → dapat didokumentasikan di PRD terpisah atau tambahan.
- **Data historis & one-to-many (1.4):** Model data dan keterangan relasi (satu WP → banyak riwayat / banyak anggota keluarga) sudah didokumentasikan di **1.4** untuk keperluan implementasi. Pengelolaan CRUD penuh untuk tiap jenis riwayat (riwayat gaji, riwayat pangkat, dll) atau data anggota keluarga dapat masuk dalam fitur Manajemen Data WP (dari halaman Detail) atau menjadi modul/PRD terpisah sesuai keputusan produk.

---

## 8. Persyaratan Produk (Nonteknis)

Detail implementasi (validasi wewenang, schema, API) didokumentasikan di SDD.

- **Wewenang:** Admin Satker hanya boleh melihat dan mengubah data WP yang berada dalam satker/unit kerjanya. Jika mencoba mengakses data WP satker lain (misalnya lewat URL), sistem menampilkan pesan jelas (misalnya "Anda tidak memiliki wewenang untuk mengakses data WP ini.").
- **Unik:** NIP harus unik di dalam sistem. Jika pengguna mengisi NIP yang sudah dipakai, tampilkan pesan error yang ramah (misalnya "NIP sudah terdaftar.").
- **Delete:** Alasan penghapusan wajib diisi di dialog konfirmasi; tanpa alasan, tombol "Ya, Hapus" tidak aktif atau submit ditolak. Sistem menyimpan alasan tersebut. Jika kebijakan produk memakai soft delete (nonaktif), list dan filter status harus mendukung tampilan "Nonaktif".
- **Relasi:** Jika data WP terhubung ke pengguna (role Widyaprada), kebijakan saat hapus data WP (apakah pengguna ikut di-nonaktifkan, di-block, atau hanya putus relasi) harus jelas; perilaku ini ditetapkan produk dan didokumentasikan (bisa di PRD atau SDD).
- **Umpan balik:** List menampilkan paginasi jika data banyak; setiap aksi (Create, Edit, Delete) menampilkan loading lalu pesan sukses atau error.
- **Akses perangkat:** Halaman List, Detail, dan form Create/Edit dapat dipakai dengan nyaman dari desktop dan tablet (layout responsif).

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: List, Detail, Create, Edit, Delete; fokus kebutuhan user & UI/UX per role | - |
| 1.1 | 2025-02-28 | Daftar calon peserta; verifikasi & validasi dokumen persyaratan; Tim Verval; tolak dengan catatan (tanpa koreksi) | - |

---

**Catatan**: Dokumen ini mencakup fitur **Manajemen Data WP** (List, Detail, Create, Edit, Delete). Keterkaitan dengan **Manajemen Pengguna** (pembuatan data WP saat Create pengguna role Widyaprada) dan modul **WPData/WPJurnal/WPUjikom** didokumentasikan dalam PRD masing-masing.
