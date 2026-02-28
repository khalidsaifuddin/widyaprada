"use client";

import BeritaDetail from "@/components/organisms/BeritaDetail";
import { useParams } from "next/navigation";

export default function BeritaDetailPage() {
  const params = useParams();
  const slug = (params?.slug as string) ?? "";
  return <BeritaDetail slug={slug} />;
}
