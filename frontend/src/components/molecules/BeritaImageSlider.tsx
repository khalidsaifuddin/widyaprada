"use client";

import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import { useCallback, useEffect, useState } from "react";
import { resolveImageUrl } from "@/lib/image";

interface BeritaImageSliderProps {
  images: string[];
  title: string;
  className?: string;
}

export default function BeritaImageSlider({ images, title, className = "" }: BeritaImageSliderProps) {
  const [current, setCurrent] = useState(0);
  const len = images.length;

  const goNext = useCallback(() => {
    setCurrent((p) => (p + 1) % len);
  }, [len]);

  const goPrev = useCallback(() => {
    setCurrent((p) => (p - 1 + len) % len);
  }, [len]);

  useEffect(() => {
    if (len <= 1) return;
    const t = setInterval(goNext, 5000);
    return () => clearInterval(t);
  }, [len, goNext]);

  if (len === 0) return null;

  const src = resolveImageUrl(images[current]);
  return (
    <div className={`relative h-64 md:h-96 rounded-xl overflow-hidden bg-gray-100 ${className}`}>
      <Image
        key={current}
        src={src}
        alt={`${title} - gambar ${current + 1}`}
        fill
        className="object-cover"
        sizes="(max-width: 768px) 100vw, 672px"
        unoptimized
      />
      {len > 1 && (
        <>
          <button
            type="button"
            onClick={goPrev}
            className="absolute left-2 top-1/2 -translate-y-1/2 p-2 rounded-full bg-black/40 text-white hover:bg-black/60"
            aria-label="Sebelumnya"
          >
            <ChevronLeftIcon className="h-6 w-6" />
          </button>
          <button
            type="button"
            onClick={goNext}
            className="absolute right-2 top-1/2 -translate-y-1/2 p-2 rounded-full bg-black/40 text-white hover:bg-black/60"
            aria-label="Selanjutnya"
          >
            <ChevronRightIcon className="h-6 w-6" />
          </button>
          <div className="absolute bottom-3 left-1/2 -translate-x-1/2 flex gap-2">
            {images.map((_, i) => (
              <button
                key={i}
                type="button"
                onClick={() => setCurrent(i)}
                className={`h-2 rounded-full transition-all ${
                  i === current ? "w-6 bg-white" : "w-2 bg-white/50 hover:bg-white/70"
                }`}
                aria-label={`Gambar ${i + 1}`}
              />
            ))}
          </div>
        </>
      )}
    </div>
  );
}
