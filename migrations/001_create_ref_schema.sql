-- Migration: Buat schema ref dan tabel-tabel referensi
-- Konvensi: id UUID PK, metadata (created_at, created_by, updated_at, updated_by, deleted_at, deleted_by), timestamptz

-- Schema ref untuk data referensi
CREATE SCHEMA IF NOT EXISTS ref;

-- Extension untuk UUID (jika belum ada)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ========== Tabel ref tanpa FK ke tabel ref lain ==========

CREATE TABLE ref.agama (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.jenis_kelamin (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) NOT NULL,
    nama VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.jenis_kepegawaian (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.golongan_ruang (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(20) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.jenjang_jabatan_fungsional (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.satker (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_satker_id UUID REFERENCES ref.satker(id),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.pendidikan_terakhir (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.status (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    tipe VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.role (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.permission (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(100) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.kategori_berita (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

-- ========== Tabel ref dengan FK ke tabel ref lain ==========

CREATE TABLE ref.pangkat (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    golongan_ruang_id UUID NOT NULL REFERENCES ref.golongan_ruang(id),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.unit_kerja (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    parent_unit_kerja_id UUID REFERENCES ref.unit_kerja(id),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

-- ========== Tabel ref wilayah ==========

CREATE TABLE ref.level_wilayah (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    urutan INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.mst_wilayah (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    level_wilayah_id UUID NOT NULL REFERENCES ref.level_wilayah(id),
    parent_wilayah_id UUID REFERENCES ref.mst_wilayah(id),
    kode VARCHAR(8) NOT NULL,
    nama VARCHAR(60) NOT NULL,
    parent_kode_wilayah VARCHAR(8),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

-- Indeks untuk kolom yang sering di-filter (soft delete)
CREATE INDEX idx_ref_agama_deleted_at ON ref.agama(deleted_at);
CREATE INDEX idx_ref_jenis_kelamin_deleted_at ON ref.jenis_kelamin(deleted_at);
CREATE INDEX idx_ref_jenis_kepegawaian_deleted_at ON ref.jenis_kepegawaian(deleted_at);
CREATE INDEX idx_ref_golongan_ruang_deleted_at ON ref.golongan_ruang(deleted_at);
CREATE INDEX idx_ref_pangkat_deleted_at ON ref.pangkat(deleted_at);
CREATE INDEX idx_ref_jenjang_jabatan_fungsional_deleted_at ON ref.jenjang_jabatan_fungsional(deleted_at);
CREATE INDEX idx_ref_satker_parent_satker_id ON ref.satker(parent_satker_id);
CREATE INDEX idx_ref_satker_deleted_at ON ref.satker(deleted_at);
CREATE INDEX idx_ref_unit_kerja_satker_id ON ref.unit_kerja(satker_id);
CREATE INDEX idx_ref_unit_kerja_parent_unit_kerja_id ON ref.unit_kerja(parent_unit_kerja_id);
CREATE INDEX idx_ref_unit_kerja_deleted_at ON ref.unit_kerja(deleted_at);
CREATE INDEX idx_ref_pendidikan_terakhir_deleted_at ON ref.pendidikan_terakhir(deleted_at);
CREATE INDEX idx_ref_status_deleted_at ON ref.status(deleted_at);
CREATE INDEX idx_ref_role_deleted_at ON ref.role(deleted_at);
CREATE INDEX idx_ref_permission_deleted_at ON ref.permission(deleted_at);
CREATE INDEX idx_ref_kategori_berita_deleted_at ON ref.kategori_berita(deleted_at);
CREATE INDEX idx_ref_level_wilayah_deleted_at ON ref.level_wilayah(deleted_at);
CREATE UNIQUE INDEX idx_ref_level_wilayah_kode ON ref.level_wilayah(kode) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_ref_mst_wilayah_kode_unique ON ref.mst_wilayah(kode) WHERE deleted_at IS NULL;
CREATE INDEX idx_ref_mst_wilayah_level_wilayah_id ON ref.mst_wilayah(level_wilayah_id);
CREATE INDEX idx_ref_mst_wilayah_parent_wilayah_id ON ref.mst_wilayah(parent_wilayah_id);
CREATE INDEX idx_ref_mst_wilayah_parent_kode_wilayah ON ref.mst_wilayah(parent_kode_wilayah);
CREATE INDEX idx_ref_mst_wilayah_deleted_at ON ref.mst_wilayah(deleted_at);

-- ========== Data awal ref.level_wilayah ==========
INSERT INTO ref.level_wilayah (id, kode, nama, urutan, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'negara', 'Negara', 1, NOW(), NOW()),
    (uuid_generate_v4(), 'provinsi', 'Provinsi', 2, NOW(), NOW()),
    (uuid_generate_v4(), 'kabupaten_kota', 'Kabupaten/Kota', 3, NOW(), NOW()),
    (uuid_generate_v4(), 'kecamatan', 'Kecamatan', 4, NOW(), NOW()),
    (uuid_generate_v4(), 'kelurahan_desa', 'Kelurahan/Desa', 5, NOW(), NOW());
