-- Migration: Buat tabel-tabel schema public (data utama & transaksional)
-- Konvensi: id UUID PK, metadata + deleted_reason (untuk audit), timestamptz, FK format nama_tabel_id

-- ========== public.user ==========
CREATE TABLE public."user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL REFERENCES ref.role(id),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    nama_lengkap VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    username VARCHAR(100),
    password_hash VARCHAR(255) NOT NULL,
    no_hp VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255),
    deleted_reason TEXT
);

CREATE UNIQUE INDEX idx_user_email_unique ON public."user"(email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_user_username_unique ON public."user"(username) WHERE deleted_at IS NULL;
CREATE INDEX idx_user_deleted_at ON public."user"(deleted_at);
CREATE INDEX idx_user_role_id ON public."user"(role_id);
CREATE INDEX idx_user_satker_id ON public."user"(satker_id);
CREATE INDEX idx_user_status_id ON public."user"(status_id);

-- ========== public.role_permission ==========
CREATE TABLE public.role_permission (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL REFERENCES ref.role(id),
    permission_id UUID NOT NULL REFERENCES ref.permission(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_role_permission_unique ON public.role_permission(role_id, permission_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_role_permission_role_id ON public.role_permission(role_id);
CREATE INDEX idx_role_permission_permission_id ON public.role_permission(permission_id);
CREATE INDEX idx_role_permission_deleted_at ON public.role_permission(deleted_at);

-- ========== public.widyaprada ==========
CREATE TABLE public.widyaprada (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES public."user"(id),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    unit_kerja_id UUID REFERENCES ref.unit_kerja(id),
    jenis_kelamin_id UUID REFERENCES ref.jenis_kelamin(id),
    golongan_ruang_id UUID REFERENCES ref.golongan_ruang(id),
    pangkat_id UUID REFERENCES ref.pangkat(id),
    jenjang_jabatan_fungsional_id UUID REFERENCES ref.jenjang_jabatan_fungsional(id),
    pendidikan_terakhir_id UUID REFERENCES ref.pendidikan_terakhir(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    agama_id UUID REFERENCES ref.agama(id),
    wilayah_id UUID REFERENCES ref.mst_wilayah(id),
    nip VARCHAR(18) NOT NULL,
    nama_lengkap VARCHAR(255) NOT NULL,
    tmt_golongan DATE,
    tmt_jabatan_fungsional DATE,
    no_sk_pengangkatan VARCHAR(100),
    no_hp VARCHAR(50),
    email VARCHAR(255),
    alamat TEXT,
    keterangan TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255),
    deleted_reason TEXT
);

CREATE UNIQUE INDEX idx_widyaprada_nip_unique ON public.widyaprada(nip) WHERE deleted_at IS NULL;
CREATE INDEX idx_widyaprada_user_id ON public.widyaprada(user_id);
CREATE INDEX idx_widyaprada_satker_id ON public.widyaprada(satker_id);
CREATE INDEX idx_widyaprada_wilayah_id ON public.widyaprada(wilayah_id);
CREATE INDEX idx_widyaprada_deleted_at ON public.widyaprada(deleted_at);

-- ========== public.slide ==========
CREATE TABLE public.slide (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    judul VARCHAR(500),
    subjudul TEXT,
    gambar_url VARCHAR(500) NOT NULL,
    url VARCHAR(500),
    label_cta VARCHAR(100),
    urutan INT NOT NULL DEFAULT 0,
    tanggal_mulai_tampil TIMESTAMPTZ,
    tanggal_selesai_tampil TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_slide_satker_id ON public.slide(satker_id);
CREATE INDEX idx_slide_status_id ON public.slide(status_id);
CREATE INDEX idx_slide_urutan ON public.slide(urutan);
CREATE INDEX idx_slide_deleted_at ON public.slide(deleted_at);

-- ========== public.berita ==========
CREATE TABLE public.berita (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    kategori_berita_id UUID REFERENCES ref.kategori_berita(id),
    judul VARCHAR(500) NOT NULL,
    slug VARCHAR(500),
    konten TEXT NOT NULL,
    ringkasan TEXT,
    thumbnail_url VARCHAR(500),
    tanggal_publikasi TIMESTAMPTZ,
    penulis VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE UNIQUE INDEX idx_berita_slug_unique ON public.berita(slug) WHERE deleted_at IS NULL;
CREATE INDEX idx_berita_satker_id ON public.berita(satker_id);
CREATE INDEX idx_berita_status_id ON public.berita(status_id);
CREATE INDEX idx_berita_tanggal_publikasi ON public.berita(tanggal_publikasi);
CREATE INDEX idx_berita_deleted_at ON public.berita(deleted_at);

-- ========== public.tautan ==========
CREATE TABLE public.tautan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    satker_id UUID NOT NULL REFERENCES ref.satker(id),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    judul VARCHAR(500) NOT NULL,
    url VARCHAR(500) NOT NULL,
    deskripsi TEXT,
    urutan INT DEFAULT 0,
    buka_tab_baru BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_tautan_satker_id ON public.tautan(satker_id);
CREATE INDEX idx_tautan_status_id ON public.tautan(status_id);
CREATE INDEX idx_tautan_deleted_at ON public.tautan(deleted_at);

-- ========== public.jurnal ==========
CREATE TABLE public.jurnal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    status_id UUID NOT NULL REFERENCES ref.status(id),
    widyaprada_id UUID REFERENCES public.widyaprada(id),
    judul VARCHAR(500) NOT NULL,
    penulis VARCHAR(255),
    abstrak TEXT,
    konten TEXT,
    tanggal_publish TIMESTAMPTZ,
    doi VARCHAR(255),
    issn VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_jurnal_status_id ON public.jurnal(status_id);
CREATE INDEX idx_jurnal_widyaprada_id ON public.jurnal(widyaprada_id);
CREATE INDEX idx_jurnal_tanggal_publish ON public.jurnal(tanggal_publish);
CREATE INDEX idx_jurnal_deleted_at ON public.jurnal(deleted_at);
