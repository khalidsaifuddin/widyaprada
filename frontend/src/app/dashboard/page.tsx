"use client";

import { Card } from "@/components";
import { app } from "@/config";

export default function DashboardPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="text-gray-600 mt-1">Ringkasan aplikasi {app.name}</p>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <Card title="Ringkasan" subtitle="Placeholder statistik">
          <p className="text-gray-500 text-sm">Konten dashboard akan diisi sesuai modul.</p>
        </Card>
      </div>
    </div>
  );
}
