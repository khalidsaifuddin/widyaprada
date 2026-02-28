# Workflow Modul WPUjikom (PlantUML)

Diagram alur (activity diagram) per modul sesuai role — **satu file per diagram**.

## Daftar file

| File | Modul | Role |
|------|-------|------|
| `Workflow_Legenda.puml` | Legenda | - |
| `Workflow_Bank_Soal.puml` | Bank Soal | Admin/Super Admin, Verifikator |
| `Workflow_Paket_Soal.puml` | Paket Soal | Admin/Super Admin, Verifikator |
| `Workflow_Manajemen_Uji_Kompetensi.puml` | Manajemen Uji Kompetensi | Admin/Super Admin, Verifikator |
| `Workflow_Assignment.puml` | Assignment (Penugasan) | Widyaprada, Admin/Super Admin |
| `Workflow_CBT.puml` | CBT | Widyaprada, Admin/Verifikator (keterangan) |

## Cara generate gambar (PNG/SVG)

### Opsi 1: VS Code
1. Pasang ekstensi **PlantUML** (j.e. by jebbs).
2. Buka salah satu file `.puml`.
3. `Alt+D` untuk preview, atau klik kanan → **Export Current Diagram**.

### Opsi 2: PlantUML CLI (Java)
```bash
# Generate satu file
java -jar plantuml.jar Workflow_Bank_Soal.puml

# Generate semua file di folder ini
java -jar plantuml.jar docs/workflows/*.puml
```

### Opsi 3: Server online
Salin isi file `.puml` ke [PlantUML Online Server](https://www.plantuml.com/plantuml/uml) lalu export PNG/SVG.

## Keterangan role

- **Admin Uji Kompetensi / Super Admin**: CRUD + aksi khusus (Terbitkan, dll). Super Admin punya akses semua modul.
- **Verifikator Uji Kompetensi**: Hanya list, detail, dan aksi Verifikasi / Batalkan Verifikasi.
- **Widyaprada**: Hanya mengakses Tugas Saya (Assignment) dan CBT; tidak akses Bank Soal, Paket Soal, Manajemen Uji Kompetensi.
