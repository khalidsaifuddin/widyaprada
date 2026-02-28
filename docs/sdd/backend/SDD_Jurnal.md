## SDD - Jurnal (List dan Detail Publik)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: List Jurnal, Detail Jurnal  

Stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur

- Data dari WPJurnal. Hanya jurnal status Published.
- Usecase: JurnalUsecase - ListPublished, GetDetail.
- Delivery: REST /api/v1/jurnal.

---

## 2. Kontrak API

- GET /api/v1/jurnal - List. Query: q, tahun, kategori, sort, page, page_size.
- GET /api/v1/jurnal/:id - Detail.

---

## 3. Skema

- journals: id, title, author, abstract, content, published_at, status, created_at, updated_at.

---

## 4. RBAC

- Endpoint publik.
