-- Migration 007: Referensi dokumen persyaratan Uji Kompetensi (Non-WP ke WP / Widyaprada Ahli Madya)
-- Sesuai lampiran persyaratan: apply uji kompetensi jabatan fungsional Widyaprada Ahli Madya
-- Sumber: Persyaratan dokumen peraturan/permen/ pedoman uji kompetensi WF Widyaprada

-- ========== REF: Tabel dokumen persyaratan ujikom ==========
CREATE TABLE ref.dokumen_persyaratan_ujikom (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) NOT NULL,
    nama TEXT NOT NULL,
    urutan INT NOT NULL DEFAULT 0,
    tipe_input VARCHAR(20) NOT NULL DEFAULT 'file',
    batasan TEXT,
    deskripsi TEXT,
    untuk_jenis_ujikom VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(255),
    deleted_at TIMESTAMPTZ,
    deleted_by VARCHAR(255)
);

COMMENT ON TABLE ref.dokumen_persyaratan_ujikom IS 'Referensi dokumen persyaratan apply uji kompetensi WF Widyaprada';
COMMENT ON COLUMN ref.dokumen_persyaratan_ujikom.tipe_input IS 'file | text_portofolio | text_essay';
COMMENT ON COLUMN ref.dokumen_persyaratan_ujikom.batasan IS 'Batasan: mis. max 1500 kata, 24 aktivitas';
COMMENT ON COLUMN ref.dokumen_persyaratan_ujikom.untuk_jenis_ujikom IS 'perpindahan_jabatan | kenaikan_tingkat | NULL=keduanya';

CREATE UNIQUE INDEX idx_ref_dokumen_persyaratan_ujikom_kode ON ref.dokumen_persyaratan_ujikom(kode) WHERE deleted_at IS NULL;
CREATE INDEX idx_ref_dokumen_persyaratan_ujikom_urutan ON ref.dokumen_persyaratan_ujikom(urutan);
CREATE INDEX idx_ref_dokumen_persyaratan_ujikom_deleted_at ON ref.dokumen_persyaratan_ujikom(deleted_at);

-- ========== Data awal: 13 dokumen persyaratan (Non-WP ke WP / Ahli Madya) ==========
INSERT INTO ref.dokumen_persyaratan_ujikom (id, kode, nama, urutan, tipe_input, batasan, deskripsi, untuk_jenis_ujikom) VALUES
    (uuid_generate_v4(), 'surat_usul_pimpinan', 'Surat usul dari pimpinan satuan kerja kepada Sekretaris Direktorat Jenderal PAUD, Pendidikan Dasar, dan Pendidikan Menengah (pelamar di lingkungan Ditjen PAUD Dikdas Dikmen), atau surat usul dari Sekretaris Unit Utama (pelamar dari unit utama lain)', 1, 'file', NULL, NULL, 'perpindahan_jabatan'),
    (uuid_generate_v4(), 'sk_kenaikan_pangkat_terakhir', 'Surat keputusan kenaikan pangkat terakhir', 2, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'sk_jabatan_terakhir', 'Surat keputusan jabatan terakhir', 3, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'surat_pernyataan_integritas_moralitas', 'Surat pernyataan pimpinan yang menyatakan calon peserta memiliki integritas dan moralitas yang baik', 4, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'surat_keterangan_sehat', 'Surat keterangan sehat dari pusat pelayanan kesehatan yang berwenang', 5, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'fotokopi_ijazah', 'Fotokopi ijazah pendidikan terakhir', 6, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'surat_keterangan_pengalaman_2tahun', 'Surat keterangan pimpinan yang menyatakan calon peserta memiliki pengalaman dalam pelaksanaan tugas di bidang penjaminan mutu pendidikan paling singkat 2 (dua) tahun', 7, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'surat_pernyataan_lowongan', 'Surat pernyataan pimpinan unit kerja mengenai ketersediaan lowongan kebutuhan pada jenjang jabatan yang akan diduduki. (Calon dari unit kerja lain: lampirkan surat pernyataan dari pimpinan unit kerja tujuan)', 8, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'surat_pernyataan_tidak_menuntut', 'Surat pernyataan tidak menuntut untuk diangkat sebagai pejabat fungsional Widyaprada', 9, 'file', NULL, NULL, NULL),
    (uuid_generate_v4(), 'penilaian_skp_2tahun', 'Penilaian sasaran kinerja pegawai 2 (dua) tahun terakhir (24 buah aktivitas di bulan berbeda)', 10, 'file', '24 aktivitas di bulan berbeda', NULL, NULL),
    (uuid_generate_v4(), 'portofolio', 'Portofolio (format terlampir): rincian pelaksanaan tugas dalam 2 (dua) tahun di bidang penjaminan mutu pendidikan, memuat minimal 1 (satu) kegiatan perbulan berisi topik, tahapan pelaksanaan tugas, dan output', 11, 'text_portofolio', 'Minimal 1 kegiatan perbulan; 2 tahun', NULL, NULL),
    (uuid_generate_v4(), 'essay_inovasi_praktik_baik', 'Tulisan/essay terkait inovasi dan aksi nyata/praktik baik di bidang penjaminan mutu pendidikan yang sudah dilakukan (untuk calon peserta uji kompetensi Widyaprada Ahli Madya)', 12, 'text_essay', 'Maksimal 1500 kata', NULL, 'perpindahan_jabatan'),
    (uuid_generate_v4(), 'surat_pernyataan_orisinalitas_essay', 'Surat pernyataan calon peserta uji kompetensi perihal orisinalitas tulisan/essay terkait inovasi dan aksi nyata/praktik, yang diketahui oleh pimpinan satuan kerja', 13, 'file', NULL, NULL, 'perpindahan_jabatan');
