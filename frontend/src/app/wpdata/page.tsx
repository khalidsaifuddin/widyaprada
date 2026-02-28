"use client";

import { Card } from "@/components";

export default function WPDataPage() {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold text-gray-900">Manajemen Data WP</h1>
      <Card title="Daftar Data WP" subtitle="List, Detail, Create, Edit, Delete">
        <p className="text-gray-500 text-sm">Halaman ini akan menampilkan manajemen data Widyaprada (WPData).</p>
      </Card>
    </div>
  );
}
