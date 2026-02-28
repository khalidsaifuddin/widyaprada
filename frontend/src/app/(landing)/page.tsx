"use client";

import AnnouncementPanel from "@/components/organisms/AnnouncementPanel";
import HeroSlider, { type SlideItem } from "@/components/organisms/HeroSlider";
import JournalPanel from "@/components/organisms/JournalPanel";
import LinksPanel from "@/components/organisms/LinksPanel";
import NewsPanel from "@/components/organisms/NewsPanel";
import { apiService } from "@/lib/api";
import { useEffect, useState } from "react";

interface LandingHomeData {
  slider?: SlideItem[];
}

export default function BerandaPage() {
  const [slider, setSlider] = useState<SlideItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    apiService
      .get<LandingHomeData>("v1/landing/home")
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as LandingHomeData;
          setSlider(d.slider ?? []);
        }
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  return (
    <>
      <HeroSlider slides={loading ? [] : slider} />
      <NewsPanel />
      <LinksPanel />
      <JournalPanel />
      <AnnouncementPanel />
    </>
  );
}
