"use client";

import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

export interface SlideItem {
  id: string;
  image_url: string;
  title: string;
  subtitle: string;
  link_url?: string;
  cta_label?: string;
  sort_order: number;
}

interface HeroSliderProps {
  slides: SlideItem[];
}

export default function HeroSlider({ slides }: HeroSliderProps) {
  const { t } = useTranslation("common");
  const [current, setCurrent] = useState(0);
  const len = slides.length;

  const goNext = useCallback(() => {
    setCurrent((p) => (p + 1) % len);
  }, [len]);

  const goPrev = useCallback(() => {
    setCurrent((p) => (p - 1 + len) % len);
  }, [len]);

  useEffect(() => {
    if (len <= 1) return;
    const t = setInterval(goNext, 6000);
    return () => clearInterval(t);
  }, [len, goNext]);

  if (len === 0) return null;

  const slide = slides[current];
  const hasLink = slide.link_url && slide.link_url.trim() !== "";

  return (
    <div
      className="relative w-full overflow-hidden bg-gray-900"
      style={{ minHeight: 416 }}
      onMouseEnter={() => {
        // pause auto-play on hover would need ref to interval
      }}
    >
      <div className="relative w-full h-[416px] md:h-[546px]">
        {slide.image_url ? (
          <Image
            src={slide.image_url}
            alt={slide.title}
            fill
            className="object-cover"
            sizes="100vw"
            priority
            unoptimized={slide.image_url.startsWith("http")}
          />
        ) : (
          <div
            className="absolute inset-0"
            style={{
              background: "linear-gradient(135deg, #022a47 0%, #044a7d 100%)",
            }}
          />
        )}
        <div className="absolute inset-0 bg-black/40" />
        <div className="absolute inset-0 flex items-center justify-center">
          <div className="max-w-4xl mx-auto px-4 text-center text-white">
            <h2 className="text-2xl md:text-4xl font-bold drop-shadow-lg">
              {slide.title || t("hero.welcome")}
            </h2>
            {slide.subtitle && (
              <p className="mt-2 text-base md:text-lg text-white/90 drop-shadow">
                {slide.subtitle}
              </p>
            )}
            {hasLink && (
              <Link
                href={slide.link_url!}
                className="mt-6 inline-block rounded-lg bg-white px-6 py-2.5 text-sm font-medium text-[#033259] hover:bg-white/90 transition-colors"
              >
                {slide.cta_label || t("hero.readMore")}
              </Link>
            )}
          </div>
        </div>
      </div>

      {len > 1 && (
        <>
          <button
            type="button"
            onClick={goPrev}
            className="absolute left-2 top-1/2 -translate-y-1/2 p-2 rounded-full bg-black/40 text-white hover:bg-black/60 transition-colors"
            aria-label={t("hero.slidePrev")}
          >
            <ChevronLeftIcon className="h-6 w-6" />
          </button>
          <button
            type="button"
            onClick={goNext}
            className="absolute right-2 top-1/2 -translate-y-1/2 p-2 rounded-full bg-black/40 text-white hover:bg-black/60 transition-colors"
            aria-label={t("hero.slideNext")}
          >
            <ChevronRightIcon className="h-6 w-6" />
          </button>
          <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2">
            {slides.map((_, i) => (
              <button
                key={i}
                type="button"
                onClick={() => setCurrent(i)}
                className={`h-2 rounded-full transition-all ${
                  i === current ? "w-6 bg-white" : "w-2 bg-white/50 hover:bg-white/70"
                }`}
                aria-label={`Slide ${i + 1}`}
              />
            ))}
          </div>
        </>
      )}
    </div>
  );
}
