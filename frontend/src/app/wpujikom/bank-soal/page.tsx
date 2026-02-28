"use client";

import { Card } from "@/components";

export default function BankSoalPage() {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Bank Soal</h1>
      <Card title="Bank Soal" subtitle="List, Detail, Create, Edit, Delete, Verifikasi">
        <p className="text-gray-500 text-sm">Kelola soal PG, Benar–Salah, Essay. Skeleton placeholder.</p>
      </Card>
    </div>
  );
}
