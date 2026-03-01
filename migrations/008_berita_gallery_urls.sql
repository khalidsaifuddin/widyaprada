-- Migration: Add gallery_urls for multiple images in berita/articles
-- Supports both public.articles (GORM) and public.berita (legacy) - run the one that matches your schema

-- For GORM-managed articles table:
ALTER TABLE public.articles ADD COLUMN IF NOT EXISTS gallery_urls JSONB DEFAULT '[]';

-- Add index for JSONB if needed for queries (optional)
-- CREATE INDEX IF NOT EXISTS idx_articles_gallery_urls ON public.articles USING GIN (gallery_urls);
