"use client";

import { Card } from "@/components";

export default function PaketSoalPage() {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Paket Soal</h1>
      <Card title="Paket Soal" subtitle="Kumpulan soal untuk ujian">
        <p className="text-gray-500 text-sm">List, Detail, Create, Edit, Delete, Verifikasi. Skeleton placeholder.</p>
      </Card>
    </div>
  );
}
