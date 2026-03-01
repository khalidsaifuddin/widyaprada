-- Migration: Add pdf_url column for journal PDF file upload
-- Table journals is managed by GORM; this adds PDF support

-- For PostgreSQL (journals table from GORM)
ALTER TABLE public.journals ADD COLUMN IF NOT EXISTS pdf_url VARCHAR(1000) DEFAULT '';
