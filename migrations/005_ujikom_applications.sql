-- Migration 005: Apply-first flow – Ujikom Applications & Documents
-- Sesuai PRD Assignment & SDD: Peserta apply dulu dengan dokumen persyaratan; penugasan ditentukan validasi Tim Verval.
-- Tabel: ujikom_applications, ujikom_application_documents.

-- ========== REF: Status untuk ujikom_application ==========
INSERT INTO ref.status (id, kode, nama, tipe, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'ujikom_menunggu_verifikasi', 'Menunggu Verifikasi', 'ujikom_application', NOW(), NOW()),
    (uuid_generate_v4(), 'ujikom_lolos', 'Lolos', 'ujikom_application', NOW(), NOW()),
    (uuid_generate_v4(), 'ujikom_tidak_lolos', 'Tidak Lolos', 'ujikom_application', NOW(), NOW());

-- ========== PUBLIC: Ujikom Applications (Apply Pendaftaran) ==========
CREATE TABLE public.ujikom_application (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES public."user"(id) ON DELETE CASCADE,
    jenis_ujikom VARCHAR(50) NOT NULL,
    status_id UUID NOT NULL REFERENCES ref.status(id),
    catatan_tolak TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

COMMENT ON COLUMN public.ujikom_application.jenis_ujikom IS 'perpindahan_jabatan | kenaikan_tingkat';

CREATE INDEX idx_ujikom_application_user_id ON public.ujikom_application(user_id);
CREATE INDEX idx_ujikom_application_status_id ON public.ujikom_application(status_id);
CREATE INDEX idx_ujikom_application_jenis_ujikom ON public.ujikom_application(jenis_ujikom);
CREATE INDEX idx_ujikom_application_deleted_at ON public.ujikom_application(deleted_at);

-- ========== PUBLIC: Ujikom Application Documents ==========
CREATE TABLE public.ujikom_application_document (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ujikom_application_id UUID NOT NULL REFERENCES public.ujikom_application(id) ON DELETE CASCADE,
    document_type VARCHAR(100) NOT NULL,
    file_path VARCHAR(500),
    portofolio_text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

COMMENT ON COLUMN public.ujikom_application_document.document_type IS 'Tipe dokumen sesuai lampiran persyaratan';
COMMENT ON COLUMN public.ujikom_application_document.portofolio_text IS 'Portofolio diketik per baris';

CREATE INDEX idx_ujikom_application_document_application_id ON public.ujikom_application_document(ujikom_application_id);
CREATE INDEX idx_ujikom_application_document_document_type ON public.ujikom_application_document(document_type);
CREATE INDEX idx_ujikom_application_document_deleted_at ON public.ujikom_application_document(deleted_at);
