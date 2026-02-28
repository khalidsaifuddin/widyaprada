## SDD - Berita (List dan Detail Publik)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: List Berita, Detail Berita  

Stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur

- Data dari CMS Berita. Hanya berita status Published.
- Usecase: BeritaUsecase - ListPublished, GetBySlug.
- Delivery: REST /api/v1/berita.

---

## 2. Kontrak API

- GET /api/v1/berita - List. Query: q, kategori, sort, page, page_size.
- GET /api/v1/berita/:slug - Detail by slug.

---

## 3. Skema

- articles: id, title, slug, content, excerpt, thumbnail_url, published_at, status, author_id, category_id.

---

## 4. RBAC

- Endpoint publik.
