-- Migration 003: User–Role many-to-many (sesuai PRD RBAC & Manajemen Pengguna)
-- Satu pengguna dapat memiliki lebih dari satu role. Relasi user–role bersifat many-to-many.
-- Setelah migrasi: public.user tidak lagi punya role_id; pemilihan role via public.user_role.

-- ========== 1. Buat tabel user_role ==========
CREATE TABLE public.user_role (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES public."user"(id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES ref.role(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_user_role_unique ON public.user_role(user_id, role_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_user_role_user_id ON public.user_role(user_id);
CREATE INDEX idx_user_role_role_id ON public.user_role(role_id);
CREATE INDEX idx_user_role_deleted_at ON public.user_role(deleted_at);

-- ========== 2. Migrasi data: salin role_id dari user ke user_role ==========
INSERT INTO public.user_role (user_id, role_id, created_at, updated_at)
SELECT id, role_id, created_at, updated_at
FROM public."user"
WHERE deleted_at IS NULL AND role_id IS NOT NULL;

-- ========== 3. Hapus kolom role_id dari user ==========
ALTER TABLE public."user" DROP COLUMN IF EXISTS role_id;

-- Catatan: Aplikasi wajib memastikan setiap user memiliki minimal satu role (validasi di layer aplikasi).
