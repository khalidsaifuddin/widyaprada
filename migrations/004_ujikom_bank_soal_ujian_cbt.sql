-- Migration 004: WPUjikom – Bank Soal, Paket Soal, Manajemen Uji Kompetensi, CBT, Assignment
-- Sesuai PRD: Bank Soal, Paket Soal, Manajemen Uji Kompetensi, CBT, Assignment (tampilkan_leaderboard).
-- Hasil jawaban peserta: satu record per peserta per ujian; jawaban disimpan di ujian_peserta.jawaban_json (JSONB).

-- ========== REF: Tabel referensi untuk Ujikom ==========

CREATE TABLE ref.tipe_soal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE ref.kategori_kompetensi (
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

CREATE TABLE ref.tingkat_kesulitan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_ref_tipe_soal_deleted_at ON ref.tipe_soal(deleted_at);
CREATE INDEX idx_ref_kategori_kompetensi_deleted_at ON ref.kategori_kompetensi(deleted_at);
CREATE INDEX idx_ref_tingkat_kesulitan_deleted_at ON ref.tingkat_kesulitan(deleted_at);

-- Data awal ref.tipe_soal
INSERT INTO ref.tipe_soal (id, kode, nama, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'PG', 'Pilihan Ganda', NOW(), NOW()),
    (uuid_generate_v4(), 'B-S', 'Benar-Salah', NOW(), NOW()),
    (uuid_generate_v4(), 'Essay', 'Essay (Uraian)', NOW(), NOW());

-- Data awal ref.tingkat_kesulitan
INSERT INTO ref.tingkat_kesulitan (id, kode, nama, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'mudah', 'Mudah', NOW(), NOW()),
    (uuid_generate_v4(), 'sedang', 'Sedang', NOW(), NOW()),
    (uuid_generate_v4(), 'sukar', 'Sukar', NOW(), NOW());

-- Data awal ref.status untuk soal, paket_soal, ujian, verifikasi (gunakan tipe)
INSERT INTO ref.status (id, kode, nama, tipe, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'soal_draft', 'Draft', 'soal', NOW(), NOW()),
    (uuid_generate_v4(), 'soal_aktif', 'Aktif', 'soal', NOW(), NOW()),
    (uuid_generate_v4(), 'paket_draft', 'Draft', 'paket_soal', NOW(), NOW()),
    (uuid_generate_v4(), 'paket_aktif', 'Aktif', 'paket_soal', NOW(), NOW()),
    (uuid_generate_v4(), 'ujian_draft', 'Draft', 'ujian', NOW(), NOW()),
    (uuid_generate_v4(), 'ujian_diterbitkan', 'Diterbitkan', 'ujian', NOW(), NOW()),
    (uuid_generate_v4(), 'ujian_berlangsung', 'Berlangsung', 'ujian', NOW(), NOW()),
    (uuid_generate_v4(), 'ujian_selesai', 'Selesai', 'ujian', NOW(), NOW()),
    (uuid_generate_v4(), 'verifikasi_belum', 'Belum Diverifikasi', 'verifikasi', NOW(), NOW()),
    (uuid_generate_v4(), 'verifikasi_sudah', 'Sudah Diverifikasi', 'verifikasi', NOW(), NOW());

-- ========== PUBLIC: Bank Soal ==========

CREATE TABLE public.soal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tipe_soal_id UUID NOT NULL REFERENCES ref.tipe_soal(id),
    kategori_kompetensi_id UUID NOT NULL REFERENCES ref.kategori_kompetensi(id),
    tingkat_kesulitan_id UUID REFERENCES ref.tingkat_kesulitan(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    status_verifikasi_id UUID NOT NULL REFERENCES ref.status(id),
    kode VARCHAR(50) NOT NULL,
    teks_soal TEXT NOT NULL,
    kunci_essay TEXT,
    bobot_nilai DECIMAL(5,2) DEFAULT 1.00,
    keterangan TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255),
    deleted_reason TEXT
);

CREATE TABLE public.soal_opsi (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    soal_id UUID NOT NULL REFERENCES public.soal(id) ON DELETE CASCADE,
    urutan INT NOT NULL DEFAULT 0,
    teks_opsi TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_soal_kode_unique ON public.soal(kode) WHERE deleted_at IS NULL;
CREATE INDEX idx_soal_tipe_soal_id ON public.soal(tipe_soal_id);
CREATE INDEX idx_soal_kategori_kompetensi_id ON public.soal(kategori_kompetensi_id);
CREATE INDEX idx_soal_status_id ON public.soal(status_id);
CREATE INDEX idx_soal_status_verifikasi_id ON public.soal(status_verifikasi_id);
CREATE INDEX idx_soal_deleted_at ON public.soal(deleted_at);

CREATE INDEX idx_soal_opsi_soal_id ON public.soal_opsi(soal_id);
CREATE INDEX idx_soal_opsi_deleted_at ON public.soal_opsi(deleted_at);

-- ========== PUBLIC: Paket Soal ==========

CREATE TABLE public.paket_soal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    status_verifikasi_id UUID NOT NULL REFERENCES ref.status(id),
    kode VARCHAR(50) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    keterangan TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255),
    deleted_reason TEXT
);

CREATE TABLE public.paket_soal_item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    paket_soal_id UUID NOT NULL REFERENCES public.paket_soal(id) ON DELETE CASCADE,
    soal_id UUID NOT NULL REFERENCES public.soal(id) ON DELETE CASCADE,
    urutan INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_paket_soal_kode_unique ON public.paket_soal(kode) WHERE deleted_at IS NULL;
CREATE INDEX idx_paket_soal_status_id ON public.paket_soal(status_id);
CREATE INDEX idx_paket_soal_status_verifikasi_id ON public.paket_soal(status_verifikasi_id);
CREATE INDEX idx_paket_soal_deleted_at ON public.paket_soal(deleted_at);

CREATE INDEX idx_paket_soal_item_paket_soal_id ON public.paket_soal_item(paket_soal_id);
CREATE INDEX idx_paket_soal_item_soal_id ON public.paket_soal_item(soal_id);
CREATE INDEX idx_paket_soal_item_deleted_at ON public.paket_soal_item(deleted_at);

-- ========== PUBLIC: Ujian (Manajemen Uji Kompetensi) + Assignment (tampilkan_leaderboard) ==========

CREATE TABLE public.ujian (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    status_verifikasi_id UUID NOT NULL REFERENCES ref.status(id),
    kode_ujian VARCHAR(50) NOT NULL,
    nama_ujian VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    jadwal_mulai TIMESTAMPTZ NOT NULL,
    jadwal_selesai TIMESTAMPTZ NOT NULL,
    durasi_menit INT NOT NULL,
    tampilkan_leaderboard BOOLEAN NOT NULL DEFAULT FALSE,
    acak_urutan_soal BOOLEAN DEFAULT FALSE,
    acak_opsi BOOLEAN DEFAULT FALSE,
    keterangan TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255),
    deleted_reason TEXT
);

CREATE TABLE public.ujian_soal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ujian_id UUID NOT NULL REFERENCES public.ujian(id) ON DELETE CASCADE,
    soal_id UUID NOT NULL REFERENCES public.soal(id) ON DELETE CASCADE,
    urutan INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE TABLE public.ujian_paket (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ujian_id UUID NOT NULL REFERENCES public.ujian(id) ON DELETE CASCADE,
    paket_soal_id UUID NOT NULL REFERENCES public.paket_soal(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_ujian_kode_ujian_unique ON public.ujian(kode_ujian) WHERE deleted_at IS NULL;
CREATE INDEX idx_ujian_status_id ON public.ujian(status_id);
CREATE INDEX idx_ujian_status_verifikasi_id ON public.ujian(status_verifikasi_id);
CREATE INDEX idx_ujian_jadwal_mulai ON public.ujian(jadwal_mulai);
CREATE INDEX idx_ujian_jadwal_selesai ON public.ujian(jadwal_selesai);
CREATE INDEX idx_ujian_deleted_at ON public.ujian(deleted_at);

CREATE INDEX idx_ujian_soal_ujian_id ON public.ujian_soal(ujian_id);
CREATE INDEX idx_ujian_soal_soal_id ON public.ujian_soal(soal_id);
CREATE INDEX idx_ujian_soal_deleted_at ON public.ujian_soal(deleted_at);

CREATE INDEX idx_ujian_paket_ujian_id ON public.ujian_paket(ujian_id);
CREATE INDEX idx_ujian_paket_paket_soal_id ON public.ujian_paket(paket_soal_id);
CREATE INDEX idx_ujian_paket_deleted_at ON public.ujian_paket(deleted_at);

-- ========== PUBLIC: Ujian Peserta (Assignment + CBT attempt + hasil jawaban) ==========
-- Satu user satu kali pengerjaan per ujian. status_pengerjaan: belum_mengerjakan | sedang | sudah_submit
-- Jawaban seluruh soal disimpan dalam satu record: kolom jawaban_json (JSONB).
-- Format jawaban_json: { "<soal_id>": { "jawaban_opsi_id": "uuid" } | { "jawaban_teks": "..." }, ... }
-- PG/B-S pakai jawaban_opsi_id; Essay pakai jawaban_teks.

CREATE TABLE public.ujian_peserta (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ujian_id UUID NOT NULL REFERENCES public.ujian(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES public."user"(id) ON DELETE CASCADE,
    status_pengerjaan VARCHAR(30) NOT NULL DEFAULT 'belum_mengerjakan',
    started_at TIMESTAMPTZ,
    submitted_at TIMESTAMPTZ,
    nilai DECIMAL(5,2),
    jawaban_json JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_ujian_peserta_ujian_user_unique ON public.ujian_peserta(ujian_id, user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_ujian_peserta_ujian_id ON public.ujian_peserta(ujian_id);
CREATE INDEX idx_ujian_peserta_user_id ON public.ujian_peserta(user_id);
CREATE INDEX idx_ujian_peserta_submitted_at ON public.ujian_peserta(submitted_at);
CREATE INDEX idx_ujian_peserta_deleted_at ON public.ujian_peserta(deleted_at);
CREATE INDEX idx_ujian_peserta_jawaban_json ON public.ujian_peserta USING GIN (jawaban_json);
