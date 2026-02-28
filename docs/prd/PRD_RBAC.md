# [PRD] Fitur RBAC – Manajemen Role, Permission, dan Kelola Role
## Product Requirements Document | Auth & RBAC

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Role, Manajemen Permission, Kelola Role (Role–Permission), dan integrasi menu dengan Manajemen Pengguna  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: RBAC (Role-Based Access Control) – Manajemen Role, Permission, Kelola Role
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Penulis**: Khalid Saifuddin
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
RBAC memungkinkan **Super Admin** untuk mengelola **Role** (daftar role, tambah/ubah/hapus role), mengelola **Permission** (daftar permission, tambah/ubah/hapus permission), dan **mengatur permission per role** (kelola role: role mana punya permission apa). **Satu pengguna dapat memiliki lebih dari satu role** (relasi user–role bersifat many-to-many); saat Create/Edit pengguna di Manajemen Pengguna, pengelola memilih satu atau lebih role untuk pengguna tersebut. Akses ke fitur RBAC (Kelola Role & Permission) hanya untuk Super Admin agar konfigurasi akses sistem tetap konsisten. Menu **Kelola Role** dan **Kelola Permission** ditempatkan di area yang sama dengan **Manajemen Pengguna** (submenu atau item sejajar di bawah Auth & RBAC).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur RBAC dirumuskan per role dalam pola user story berikut.

### 2.1 Manajemen Role

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Super Admin | pengelola RBAC | melihat daftar role dalam satu halaman (kode, nama, jumlah permission) | memantau role yang ada di sistem |
| 2 | Super Admin | pengelola RBAC | mencari role berdasarkan kode atau nama | menemukan role tertentu dengan cepat |
| 3 | Super Admin | pengelola RBAC | membuka detail satu role (termasuk daftar permission yang di-assign) | memeriksa dan mengatur permission role tersebut |
| 4 | Super Admin | pengelola RBAC | menambah role baru (Create) dengan kode dan nama | mendefinisikan role baru untuk dipakai di pengguna |
| 5 | Super Admin | pengelola RBAC | mengubah data role (Edit) dan mengatur permission yang di-assign ke role | menyesuaikan wewenang suatu role |
| 6 | Super Admin | pengelola RBAC | menghapus role (Delete) dengan konfirmasi dan alasan penghapusan | membersihkan role yang tidak dipakai |
| 7 | Super Admin | pengelola RBAC | di halaman Detail/Edit role memilih atau membatalkan permission untuk role tersebut | mengontrol akses per role tanpa halaman terpisah |

### 2.2 Manajemen Permission

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Super Admin | pengelola RBAC | melihat daftar permission (kode, nama, modul/group) | memantau permission yang tersedia |
| 2 | Super Admin | pengelola RBAC | mencari permission berdasarkan kode atau nama | menemukan permission tertentu |
| 3 | Super Admin | pengelola RBAC | membuka detail satu permission | melihat informasi dan role yang memakai permission ini (opsional) |
| 4 | Super Admin | pengelola RBAC | menambah permission baru (Create) | mendefinisikan permission untuk fitur baru |
| 5 | Super Admin | pengelola RBAC | mengubah data permission (Edit) | memperbaiki kode/nama/deskripsi |
| 6 | Super Admin | pengelola RBAC | menghapus permission (Delete) dengan konfirmasi dan alasan | membersihkan permission yang tidak dipakai |

### 2.3 Menu dan Navigasi (Kelola Role di menu Manajemen Pengguna)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Super Admin | pengelola sistem | dari area yang sama dengan Manajemen Pengguna bisa mengakses **Kelola Role** dan **Kelola Permission** | mengelola pengguna dan konfigurasi role/permission dari satu tempat |
| 2 | Admin Satker | pengelola pengguna | menu Manajemen Pengguna tetap seperti sekarang; saya tidak perlu akses ke Kelola Role/Permission | tidak bingung dan tidak mengubah konfigurasi role sistem |

**Keterangan**: Hanya **Super Admin** yang mengakses Manajemen Role, Manajemen Permission, dan Kelola Role (assign permission ke role). Admin Satker dan Widyaprada tidak mengakses fitur RBAC ini; mereka hanya memakai role yang sudah ditetapkan. Saat Create/Edit pengguna di Manajemen Pengguna, **satu atau lebih role** dapat dipilih untuk satu pengguna (multi-select); daftar role diisi dari data role yang dikelola Super Admin di sini.

---

### 2.4 Kebutuhan Pengguna – Yang Diinginkan
- **Manajemen Role**: List, Detail, Create, Edit, Delete; di Detail/Edit ada section untuk mengatur permission yang di-assign ke role (Kelola Role).
- **Manajemen Permission**: List, Detail, Create, Edit, Delete; kode unik, nama, dan opsional group/modul untuk pengelompokan.
- **Kelola Role (Role–Permission)**: Dikelola dari halaman Detail/Edit Role (pilih/batalkan permission), bukan hanya halaman list terpisah.
- **Menu**: Di menu navigasi (sidebar/drawer), di bawah atau sejajar dengan **Manajemen Pengguna**, tersedia **Kelola Role** dan **Kelola Permission** (hanya tampil untuk Super Admin).
- **Umpan balik**: Loading pada aksi; pesan sukses/error yang jelas; konfirmasi dan alasan penghapusan untuk Delete.

### 2.5 Kebutuhan Pengguna – Yang Tidak Diinginkan
- Admin Satker atau Widyaprada bisa mengubah role atau permission.
- List tanpa search sehingga sulit mencari role/permission.
- Delete tanpa konfirmasi dan tanpa alasan penghapusan.
- Menu Kelola Role/Permission tersembunyi atau berada di tempat yang tidak terkait dengan Manajemen Pengguna.

---

## 3. Antarmuka Pengguna (UI)

*Semua halaman mengikuti design system aplikasi. Hanya Super Admin yang melihat dan mengakses menu serta halaman RBAC.*

---

### 3.1 Struktur Menu (Navigasi)

**Deskripsi**: Di area Auth & RBAC (atau Manajemen Pengguna), menu berisi:

- **Manajemen Pengguna:** Sesuai PRD Manajemen Pengguna; akses Admin Satker + Super Admin sesuai wewenang.
- **Kelola Role:** List/Detail/Create/Edit/Delete Role dan pengaturan permission per role. **Hanya tampil untuk Super Admin.**
- **Kelola Permission:** List/Detail/Create/Edit/Delete Permission. **Hanya tampil untuk Super Admin.**

**Elemen yang Harus Ada**:
- Item menu "Manajemen Pengguna" dengan link ke halaman list pengguna.
- Item menu "Kelola Role" dengan link ke halaman list role (hanya terlihat untuk Super Admin).
- Item menu "Kelola Permission" dengan link ke halaman list permission (hanya terlihat untuk Super Admin).
- Urutan dan pengelompokan (misalnya submenu "Manajemen Pengguna" berisi: Daftar Pengguna, Kelola Role, Kelola Permission) dapat disesuaikan dengan design system; yang penting Kelola Role dan Kelola Permission mudah ditemukan di area yang sama dengan Manajemen Pengguna.

---

### 3.2 Manajemen Role – List Role

**Deskripsi**: Halaman daftar role.

**Elemen yang Harus Ada**:
- Judul: "Kelola Role" atau "Daftar Role".
- Tombol "Tambah Role" ke form Create.
- Kotak pencarian (kode, nama).
- Tabel: kolom minimal Kode, Nama, Jumlah Permission (opsional), tanggal dibuat/diubah (opsional). Kolom aksi: Detail, Edit, Hapus.
- Paginasi jika data banyak.
- Umpan balik: loading, pesan kosong, pesan sukses/error setelah Delete.

---

### 3.3 Manajemen Role – Detail Role

**Deskripsi**: Halaman detail satu role dan daftar permission yang di-assign.

**Elemen yang Harus Ada**:
- Judul: "Detail Role" atau nama role.
- Informasi role: Kode, Nama.
- **Section "Permission untuk Role ini"**: Daftar permission yang saat ini di-assign ke role (bisa tabel atau tag list). Tombol "Edit Role" untuk masuk ke form Edit (di mana permission bisa dipilih/dibatalkan).
- Tombol "Edit", "Hapus", "Kembali ke Daftar Role".

---

### 3.4 Manajemen Role – Create / Edit Role

**Deskripsi**: Form tambah atau ubah role; di Edit termasuk pengaturan permission.

**Elemen yang Harus Ada**:
- **Create**: Field Kode (unik), Nama. Tombol Simpan, Batal.
- **Edit**: Field Kode (boleh readonly), Nama; **section pilih permission**: checklist atau multi-select dari daftar permission (bisa dikelompokkan per modul). Tombol Simpan, Batal.
- Validasi: kode wajib dan unik; nama wajib.
- Umpan balik: loading saat simpan, pesan sukses/error.

---

### 3.5 Manajemen Role – Delete Role

**Deskripsi**: Hapus role dengan konfirmasi dan alasan.

**Elemen yang Harus Ada**:
- Dialog konfirmasi dengan field wajib **Alasan penghapusan**.
- Tombol Batal dan Ya, Hapus. Setelah alasan diisi dan konfirmasi, proses hapus; tampilkan pesan sukses/error.
- Validasi: jika role masih dipakai oleh pengguna, tampilkan pesan yang jelas (misalnya "Role ini masih digunakan oleh N pengguna. Ubah atau hapus pengguna tersebut terlebih dahulu.").

---

### 3.6 Manajemen Permission – List Permission

**Deskripsi**: Halaman daftar permission.

**Elemen yang Harus Ada**:
- Judul: "Kelola Permission" atau "Daftar Permission".
- Tombol "Tambah Permission".
- Kotak pencarian (kode, nama).
- Filter opsional: group/modul (jika permission punya field group).
- Tabel: Kode, Nama, Group/Modul (jika ada). Aksi: Detail, Edit, Hapus.
- Paginasi, loading, pesan kosong/sukses/error.

---

### 3.7 Manajemen Permission – Detail / Create / Edit / Delete

**Deskripsi**: Detail satu permission; form Create/Edit (kode, nama, deskripsi/group); Delete dengan konfirmasi dan alasan.

**Elemen yang Harus Ada**:
- **Detail**: Tampilkan kode, nama, group/deskripsi; tombol Edit, Hapus, Kembali.
- **Create/Edit**: Field Kode (unik), Nama, Group/Modul atau Deskripsi (opsional). Simpan, Batal.
- **Delete**: Konfirmasi + field wajib alasan penghapusan; validasi: jika permission masih di-assign ke role, beri pesan yang jelas.

---

## 4. Pengalaman Pengguna (UX) – Alur

### 4.1 Alur Kelola Role
1. Super Admin membuka menu **Kelola Role**.
2. Melihat list role; bisa search, klik Detail atau Edit.
3. Di Detail: melihat permission role; klik Edit untuk mengubah data role atau mengatur permission.
4. Di Edit: ubah nama; centang/uncentang permission; Simpan.
5. Create: klik Tambah Role, isi kode dan nama, Simpan; lalu bisa Edit untuk assign permission.
6. Delete: dari Detail atau List, klik Hapus → isi alasan → konfirmasi → pesan sukses/error.

### 4.2 Alur Kelola Permission
1. Super Admin membuka menu **Kelola Permission**.
2. List permission; search/filter; Detail, Edit, atau Hapus.
3. Create: Tambah Permission, isi kode/nama/group, Simpan.
4. Edit: ubah data, Simpan.
5. Delete: konfirmasi + alasan, proses hapus.

### 4.3 Integrasi dengan Manajemen Pengguna
- Dari **Manajemen Pengguna**, Super Admin bisa beralih ke **Kelola Role** atau **Kelola Permission** lewat menu yang sama (sidebar/submenu).
- Saat Create/Edit pengguna, dropdown **Role** diisi dari data role yang dikelola di Kelola Role; permission tidak perlu dipilih per pengguna (sudah melekat pada role).

---

## 5. Kebutuhan per Role

### 5.1 Super Admin
- **Manajemen Role**: List, Detail, Create, Edit, Delete; assign/unassign permission di Edit Role.
- **Manajemen Permission**: List, Detail, Create, Edit, Delete.
- **Menu**: Melihat dan mengakses "Manajemen Pengguna", "Kelola Role", "Kelola Permission".

### 5.2 Admin Satker
- Tidak mengakses Kelola Role dan Kelola Permission; menu tersebut tidak ditampilkan.
- Hanya menggunakan Manajemen Pengguna (dalam scope satker); dropdown role berisi role yang sudah dikonfigurasi Super Admin.

### 5.3 Widyaprada
- Tidak mengakses Manajemen Pengguna, Kelola Role, maupun Kelola Permission.

---

## 6. Ringkasan Perbedaan per Role

| Aspek | Super Admin | Admin Satker | Widyaprada |
|-------|--------------|--------------|------------|
| Manajemen Pengguna | Ya (semua pengguna) | Ya (dalam satker) | Tidak |
| Kelola Role | Ya penuh | Tidak (menu tidak tampil) | Tidak |
| Kelola Permission | Ya penuh | Tidak (menu tidak tampil) | Tidak |
| Assign permission ke role | Ya (via Edit Role) | Tidak | Tidak |

---

## 7. Acceptance Criteria (Ringkas)

**Kelola Role:** (1) Hanya Super Admin yang melihat menu dan mengakses halaman. (2) List menampilkan kode, nama, jumlah permission; search kode/nama; paginasi. (3) Create: form kode (unik), nama; validasi wajib dan unik. (4) Edit: ubah nama + section pilih/batalkan permission (checklist/multi-select). (5) Delete: dialog konfirmasi + field wajib alasan; jika role masih dipakai pengguna, tampilkan pesan jelas dan hapus tidak diproses. (6) Detail menampilkan permission yang di-assign; tombol Edit, Hapus, Kembali.

**Kelola Permission:** (1) Hanya Super Admin. (2) List: kode, nama, group/modul; search; filter group (jika ada). (3) Create/Edit: kode (unik), nama, group/deskripsi (opsional). (4) Delete: konfirmasi + alasan wajib; jika permission masih di-assign ke role, pesan jelas dan hapus tidak diproses atau assignment harus dicabut dulu.

**Menu:** Kelola Role dan Kelola Permission tampil di area yang sama dengan Manajemen Pengguna (sidebar/submenu); hanya terlihat untuk Super Admin. Admin Satker dan Widyaprada tidak melihat kedua menu tersebut.

---

## 8. Cakupan Fitur RBAC

### 8.1 Termasuk
- **Manajemen Role:** List, Detail, Create, Edit, Delete; pengaturan permission per role (Kelola Role) dari halaman Detail/Edit Role.
- **Manajemen Permission:** List, Detail, Create, Edit, Delete.
- **Menu:** Kelola Role dan Kelola Permission di menu yang terkait Manajemen Pengguna (area Auth & RBAC), hanya tampil untuk Super Admin.
- **Validasi:** Kode role dan permission unik; Delete dengan konfirmasi dan alasan; validasi jika role/permission masih dipakai sebelum hapus.
- **Umpan balik:** Loading, pesan sukses/error, pesan jelas saat data masih terpakai.

### 8.2 Tidak Termasuk
- Login, Logout, Lupa Password → PRD masing-masing.
- Manajemen Pengguna (CRUD pengguna) → PRD Manajemen Pengguna.
- Audit log perubahan role/permission (jika ada) → dapat didokumentasikan di PRD terpisah.
- Permission otomatis dari kode (code-based permission discovery) → dapat ditambah di fase berikutnya jika diperlukan.

---

## 9. Persyaratan Produk (Nonteknis)

Persyaratan berikut wajib dipenuhi dari sisi produk. Detail implementasi (schema, API, validasi backend) didokumentasikan di SDD.

- **Wewenang:** Hanya Super Admin yang boleh mengakses Kelola Role dan Kelola Permission. Menu tersebut hanya tampil untuk Super Admin; pengguna lain tidak melihat dan tidak bisa membuka halaman tersebut.
- **Unik:** Kode role dan kode permission harus unik di dalam sistem. Jika pengguna mengisi kode yang sudah ada, sistem menampilkan pesan error yang ramah (misalnya "Kode role sudah digunakan").
- **Delete role:** Sebelum role dihapus, sistem memeriksa apakah ada pengguna yang memakai role tersebut. Jika ada, tampilkan pesan jelas (misalnya "Role ini masih digunakan oleh N pengguna. Ubah atau hapus pengguna tersebut terlebih dahulu.") dan hapus tidak diproses. Alasan penghapusan wajib diisi dan disimpan.
- **Delete permission:** Jika permission masih di-assign ke satu atau lebih role, sistem memberi pesan jelas; hapus tidak diproses sampai assignment dicabut, atau produk memutuskan perilaku lain (tolak hapus dengan pesan).
- **Umpan balik:** Setiap aksi (simpan, hapus) menampilkan loading; setelah selesai tampilkan pesan sukses atau error. List menampilkan paginasi jika data banyak.
- **Akses perangkat:** Halaman List, Detail, dan form dapat dipakai dengan nyaman dari desktop dan tablet (layout responsif).

---

## 10. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Manajemen Role, Manajemen Permission, Kelola Role (assign permission di Edit Role), menu Kelola Role dan Kelola Permission di area Manajemen Pengguna | Khalid Saifuddin |

---

**Catatan**: Dokumen ini mencakup fitur **RBAC** (Manajemen Role, Manajemen Permission, Kelola Role). Fitur **Manajemen Pengguna** (CRUD pengguna) didokumentasikan di PRD Manajemen Pengguna. Menu **Kelola Role** dan **Kelola Permission** ditempatkan di menu yang sama dengan Manajemen Pengguna agar pengelolaan pengguna dan konfigurasi role/permission terpusat untuk Super Admin.
