# [PRD] Fitur Logout
## Product Requirements Document | Auth & RBAC

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Logout  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Logout
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-01
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Logout memungkinkan pengguna yang sudah masuk (Widyaprada, Admin Satker, Super Admin) untuk mengakhiri sesi secara sengaja dari satu tempat yang mudah dijangkau (misalnya tombol atau tautan "Keluar" di menu/header). Setelah logout, sesi berakhir dan pengguna diarahkan ke halaman login. Pengalaman ini sama untuk semua peran: satu aksi keluar, lalu kembali ke layar login sehingga pengguna yakin sudah keluar dan bisa masuk lagi kapan pun dengan kredensial yang sama atau dari perangkat lain dengan aman.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur Logout dirumuskan per role dalam pola user story berikut, dalam format tabular.

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Semua | pengguna yang sudah login | mengakhiri sesi dari satu tombol atau tautan "Keluar" yang mudah ditemukan (misalnya di header atau menu profil) | keluar dari aplikasi tanpa bingung mencari cara logout |
| 2 | Semua | pengguna yang sudah login | setelah menekan Keluar langsung diarahkan ke halaman login dengan sesi yang benar-benar berakhir | yakin bahwa saya sudah keluar dan orang lain tidak bisa melanjutkan sesi saya di perangkat yang sama |
| 3 | Semua | pengguna yang sudah login | melihat umpan balik singkat setelah logout (misalnya "Anda telah keluar.") di halaman login | tahu bahwa aksi logout berhasil dan bukan karena error |
| 4 | Semua | pengguna yang sudah login | tombol/tautan Keluar menampilkan keadaan "Memproses…" atau loading singkat saat logout diproses (jika perlu) | yakin bahwa aplikasi sedang memproses keluar dan tidak hang |
| 5 | Widyaprada | Widyaprada yang sudah login | keluar dari aplikasi dari mana pun saya berada (dashboard, WPData, WPJurnal, WPUjikom, dll) dengan aksi yang sama | mengakhiri kerja atau ganti akun tanpa harus kembali ke halaman tertentu dulu |
| 6 | Admin Satker | Admin Satker yang sudah login | keluar dari aplikasi dari mana pun saya berada (CMS, Manajemen Data WP, Manajemen Pengguna, dll) dengan aksi yang sama | mengakhiri sesi admin atau ganti akun dengan cepat |
| 7 | Super Admin | Super Admin yang sudah login | keluar dari aplikasi dari mana pun saya berada dengan aksi yang sama | mengakhiri sesi atau ganti akun dengan konsisten |
| 8 | Semua | pengguna yang sudah login | setelah logout tidak bisa mengakses halaman yang memerlukan login tanpa masuk lagi | memastikan sesi benar-benar berakhir dan tidak ada akses sisa |
| 9 | Semua | pengguna yang sudah logout | jika mencoba membuka URL halaman yang memerlukan login diarahkan ke halaman login | tidak melihat halaman kosong atau error dan tahu bahwa saya harus masuk lagi |

**Keterangan**: *Semua* = berlaku untuk Widyaprada, Admin Satker, dan Super Admin.

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- Satu **tempat jelas** untuk keluar (tombol atau tautan "Keluar" / "Logout") yang mudah dijangkau dari mana pun di dalam aplikasi (misalnya header atau menu profil).
- Setelah menekan Keluar, **sesi benar-benar berakhir** dan pengguna **langsung diarahkan ke halaman login**.
- **Pesan singkat** setelah logout (misalnya "Anda telah keluar.") di halaman login agar pengguna yakin aksi berhasil.
- Perilaku **sama untuk semua peran** - tidak perlu memilih "logout sebagai apa"; satu aksi untuk keluar.
- Setelah logout, **tidak ada akses** ke halaman yang memerlukan login; upaya akses diarahkan ke login.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- Harus mencari-cari di mana tombol logout (misalnya tersembunyi di menu yang jarang dibuka).
- Setelah menekan Keluar, masih bisa mengakses halaman dalam aplikasi tanpa login lagi (sesi tidak benar-benar berakhir).
- Pesan atau tampilan yang membingungkan setelah logout (misalnya halaman kosong atau error tanpa penjelasan).
- Perbedaan alur logout antar peran yang tidak perlu (semua role cukup satu cara keluar).

---

## 3. Antarmuka Pengguna (UI)

*Fitur Logout tidak memiliki halaman khusus; yang diperlukan adalah elemen pemicu (tombol/tautan) dan perilaku setelah logout. Semua role melihat dan menggunakan elemen yang sama.*

### 3.1 Elemen yang Harus Ada
- **Tombol atau tautan "Keluar" (atau "Logout")**: Satu elemen yang jelas dan konsisten di seluruh aplikasi. Posisi yang umum: **header** (pojok kanan) atau **menu profil / menu akun** (dropdown berisi nama pengguna dan opsi "Keluar"). Label dapat berupa teks "Keluar", "Logout", atau ikon keluar dengan tooltip "Keluar".
- **Keterjangkauan**: Elemen Keluar harus **mudah ditemukan** dari halaman mana pun yang bisa diakses pengguna yang sudah login (dashboard, list, form, dll) - idealnya selalu terlihat di header atau satu klik dari menu profil.

### 3.2 Tata Letak dan Keterbacaan
- Tombol/tautan Keluar tidak tertutup oleh elemen lain dan ukurannya cukup untuk diklik (termasuk di layar sentuh).
- Jika menggunakan menu dropdown (misalnya menu profil), label "Keluar" harus jelas dan tidak ambigu.
- Di layar kecil (ponsel/tablet), tombol atau menu Keluar tetap bisa diakses (misalnya ikon menu yang memuat opsi Keluar).

### 3.3 Umpan Balik
- **Sedang memproses** (opsional): Jika logout membutuhkan waktu (misalnya panggilan ke server), tombol/tautan dapat menampilkan keadaan loading singkat (spinner atau "Memproses…") agar pengguna tidak mengira klik tidak terbaca.
- **Setelah logout berhasil**: Pengguna **diarahkan ke halaman login**. Di halaman login dapat ditampilkan **satu pesan sukses** yang ramah, misalnya **"Anda telah keluar."** - dengan gaya yang konsisten dengan pesan lain di aplikasi (misalnya warna hijau atau info).
- Pesan tersebut dapat hilang setelah beberapa detik atau saat pengguna mulai mengisi form login, agar tidak mengganggu.

---

## 4. Pengalaman Pengguna (UX) - Alur Logout

### 4.1 Alur Umum (Semua Role)
1. Pengguna sedang berada di dalam aplikasi (sudah login), di halaman mana pun (dashboard, list, form, dll).
2. Pengguna mencari dan menekan **tombol atau tautan "Keluar"** (di header atau menu profil).
3. (Opsional) Jika sistem menampilkan loading, pengguna melihat umpan balik singkat "Memproses…" atau spinner.
4. Sesi berakhir; pengguna **diarahkan ke halaman login**.
5. Di halaman login, pengguna melihat **pesan "Anda telah keluar."** (atau setara).
6. Pengguna bisa langsung **memasukkan kredensial lagi** jika ingin masuk kembali, atau menutup browser.

### 4.2 Perilaku yang Diharapkan
- **Satu aksi** (satu klik Keluar) cukup untuk mengakhiri sesi dan sampai di halaman login; tidak perlu langkah konfirmasi tambahan kecuali ada kebijakan khusus (misalnya konfirmasi "Yakin keluar?" bisa dipertimbangkan untuk konteks tertentu, tetapi untuk pengalaman standar satu klik langsung keluar lebih sederhana).
- Setelah logout, **semua upaya akses ke halaman yang memerlukan login** (misalnya mengklik tombol Back atau membuka bookmark halaman dalam) **mengarahkan ke halaman login**, bukan menampilkan halaman dalam aplikasi atau error yang membingungkan.
- Pengalaman logout **sama untuk Widyaprada, Admin Satker, dan Super Admin** - tidak ada perbedaan alur atau tempat tombol Keluar per role.

---

## 5. Kebutuhan per Role

*Fitur Logout tidak membedakan perilaku per role; semua pengguna yang sudah login menggunakan cara yang sama untuk keluar.*

### 5.1 Semua Role (Widyaprada, Admin Satker, Super Admin)
- **Siapa**: Setiap pengguna yang sudah masuk ke aplikasi (apa pun perannya).
- **Kebutuhan**: Satu tombol/tautan "Keluar" yang mudah dijangkau; satu klik → sesi berakhir → redirect ke halaman login dengan pesan "Anda telah keluar." (atau setara).
- **UI/UX**: Tombol/tautan Keluar hadir di setiap halaman yang bisa diakses setelah login (biasanya via header atau menu profil). Setelah logout, tidak ada perbedaan pesan atau alur per role - semua melihat halaman login yang sama.

---

## 6. Ringkasan Perbedaan per Role

| Aspek | Widyaprada | Admin Satker | Super Admin |
|-------|------------|--------------|-------------|
| **Cara logout** | Sama: tombol/tautan Keluar → redirect ke login | Sama | Sama |
| **Lokasi tombol Keluar** | Sama (header atau menu profil) | Sama | Sama |
| **Setelah logout** | Halaman login + pesan "Anda telah keluar." | Sama | Sama |
| **Pengalaman** | Satu klik keluar, lalu bisa login lagi kapan pun | Sama | Sama |

*Tidak ada perbedaan fungsional per role untuk fitur Logout.*

---

## 7. Cakupan Fitur Logout

### 7.1 Termasuk dalam Fitur Logout
- Tombol atau tautan **"Keluar"** (atau "Logout") yang konsisten dan mudah dijangkau (header atau menu profil) di seluruh aplikasi.
- Satu aksi (klik) untuk mengakhiri sesi dan mengarahkan pengguna ke **halaman login**.
- Pesan singkat di halaman login setelah logout (misalnya **"Anda telah keluar."**).
- Umpan balik loading pada tombol/tautan Keluar saat logout diproses (jika diperlukan).
- Setelah logout, upaya akses ke halaman yang memerlukan login diarahkan ke **halaman login** (tidak ada akses sisa).

### 7.2 Tidak Termasuk (Fitur Lain)
- Login → PRD Login.
- Lupa password → PRD Lupa Password.
- Manajemen pengguna, role, permission → PRD masing-masing modul.
- Sesi habis otomatis (timeout) → disebut di PRD Login; implementasi logout saat timeout dapat mengikuti perilaku yang sama (redirect ke login + pesan "Sesi Anda telah berakhir").

---

## 8. Persyaratan Produk (Nonteknis)

Detail implementasi (invalidasi sesi/token, pembersihan state klien) didokumentasikan di SDD.

- **Akhir sesi:** Setelah pengguna menekan Keluar, sesi benar-benar berakhir. Upaya mengakses halaman yang memerlukan login (misalnya dengan tombol Back atau bookmark) harus mengarahkan ke halaman login, bukan menampilkan halaman dalam aplikasi.
- **Ketersediaan:** Tombol/tautan Keluar berfungsi kapan pun pengguna yang sudah login menggunakannya.
- **Performa:** Waktu dari menekan "Keluar" sampai redirect ke halaman login harus wajar (tidak sampai puluhan detik) agar pengguna yakin proses berhasil.
- **Akses perangkat:** Tombol/tautan Keluar dapat digunakan dengan nyaman dari komputer, tablet, dan ponsel (ukuran dan posisi tetap bisa diklik).

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-01 | Dokumen awal, mengikuti pola PRD Login & Lupa Password, fokus user & UI/UX | - |

---

**Catatan**: Dokumen ini hanya mencakup fitur **Logout**. Fitur **Login** dan **Lupa Password** didokumentasikan dalam PRD masing-masing.
