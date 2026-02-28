package dokumenpersyaratanrepo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedDokumenPersyaratan seeds 13 dokumen persyaratan (from migration 007)
func SeedDokumenPersyaratan(db *gorm.DB) error {
	var count int64
	db.Model(&DokumenPersyaratanUjikom{}).Count(&count)
	if count > 0 {
		return nil
	}

	now := time.Now().UTC()
	docs := []DokumenPersyaratanUjikom{
		{ID: uuid.New().String(), Kode: "surat_usul_pimpinan", Nama: "Surat usul dari pimpinan satuan kerja...", Urutan: 1, TipeInput: "file", UntukJenisUjikom: "perpindahan_jabatan", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "sk_kenaikan_pangkat_terakhir", Nama: "Surat keputusan kenaikan pangkat terakhir", Urutan: 2, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "sk_jabatan_terakhir", Nama: "Surat keputusan jabatan terakhir", Urutan: 3, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_pernyataan_integritas_moralitas", Nama: "Surat pernyataan integritas dan moralitas", Urutan: 4, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_keterangan_sehat", Nama: "Surat keterangan sehat", Urutan: 5, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "fotokopi_ijazah", Nama: "Fotokopi ijazah pendidikan terakhir", Urutan: 6, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_keterangan_pengalaman_2tahun", Nama: "Surat keterangan pengalaman 2 tahun", Urutan: 7, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_pernyataan_lowongan", Nama: "Surat pernyataan lowongan", Urutan: 8, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_pernyataan_tidak_menuntut", Nama: "Surat pernyataan tidak menuntut", Urutan: 9, TipeInput: "file", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "penilaian_skp_2tahun", Nama: "Penilaian SKP 2 tahun", Urutan: 10, TipeInput: "file", Batasan: "24 aktivitas di bulan berbeda", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "portofolio", Nama: "Portofolio", Urutan: 11, TipeInput: "text_portofolio", Batasan: "Minimal 1 kegiatan perbulan; 2 tahun", UntukJenisUjikom: "", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "essay_inovasi_praktik_baik", Nama: "Essay inovasi praktik baik", Urutan: 12, TipeInput: "text_essay", Batasan: "Maksimal 1500 kata", UntukJenisUjikom: "perpindahan_jabatan", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Kode: "surat_pernyataan_orisinalitas_essay", Nama: "Surat pernyataan orisinalitas essay", Urutan: 13, TipeInput: "file", UntukJenisUjikom: "perpindahan_jabatan", CreatedAt: &now, UpdatedAt: &now},
	}

	return db.CreateInBatches(docs, 5).Error
}
