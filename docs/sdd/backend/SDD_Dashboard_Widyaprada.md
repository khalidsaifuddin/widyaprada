## SDD – Dashboard Widyaprada

**Aplikasi**: Widyaprada  
**Modul**: Dashboard User  
**Fitur**: Dashboard Widyaprada – Assignment Uji Kompetensi & Jurnal Di-Submit  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Dashboard Widyaprada` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `usecase`: `DashboardUsecase.GetAssignments(ctx, userID, limit)`, `DashboardUsecase.GetMyJournals(ctx, userID, limit)`.
  - `delivery/http`: `GET /api/v1/dashboard/assignments`, `GET /api/v1/dashboard/journals`.
  - **Sumber data**: Manajemen Uji Kompetensi (peserta per ujian), WPJurnal (jurnal per user).

---

## 2. Kontrak API

### 2.1 Daftar Assignment (Tugas Saya)

- **Endpoint**: `GET /api/v1/dashboard/assignments`
- **Query**: `limit` (default 10), `page` (opsional).
- **Response**:

```json
{
  "data": [
    {
      "id": "uuid",
      "exam_name": "string",
      "deadline": "datetime",
      "status": "belum_dikerjakan|sudah_dikerjakan",
      "score": 85,
      "can_start": true,
      "can_view_result": true,
      "can_view_leaderboard": false
    }
  ],
  "meta": { "total": 5 }
}
```

### 2.2 Daftar Jurnal Saya

- **Endpoint**: `GET /api/v1/dashboard/journals`
- **Query**: `limit` (default 10), `page`.
- **Response**:

```json
{
  "data": [
    {
      "id": "uuid",
      "title": "string",
      "submitted_at": "datetime",
      "status": "Draft|Menunggu Verifikasi|Diverifikasi|Ditolak|Published"
    }
  ],
  "meta": { "total": 3 }
}
```

---

## 3. Aturan Bisnis

- **Assignments**: Hanya ujian yang (1) user termasuk pesertanya, (2) status Diterbitkan/Berlangsung/Selesai, (3) urutan batas waktu terdekat / belum dikerjakan dulu.
- **Journals**: Hanya jurnal yang `created_by`/`submitted_by` = user login; filter status sesuai WPJurnal.

---

## 4. RBAC

- Role: **Widyaprada** (minimal). Endpoint dilindungi auth middleware.
