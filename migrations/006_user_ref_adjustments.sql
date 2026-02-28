-- Migration 006: Penyesuaian user & ref untuk PRD terbaru
-- 1. user.satker_id: nullable untuk calon peserta yang registrasi mandiri (belum punya satker)
-- 2. ref.permission: tambah kolom "group", description (SDD RBAC)
-- 3. ref.role: tambah kolom description (SDD RBAC)

-- ========== 1. user.satker_id → nullable ==========
ALTER TABLE public."user" ALTER COLUMN satker_id DROP NOT NULL;

COMMENT ON COLUMN public."user".satker_id IS 'Nullable untuk calon peserta (registrasi mandiri). Diisi saat verifikasi/assign satker.';

-- ========== 2. ref.permission: tambah group, description ==========
ALTER TABLE ref.permission ADD COLUMN IF NOT EXISTS "group" VARCHAR(100);
ALTER TABLE ref.permission ADD COLUMN IF NOT EXISTS description TEXT;

COMMENT ON COLUMN ref.permission."group" IS 'Group permission untuk filter/kelompok (mis: AUTH, RBAC, WP_DATA)';

-- ========== 3. ref.role: tambah description ==========
ALTER TABLE ref.role ADD COLUMN IF NOT EXISTS description TEXT;
