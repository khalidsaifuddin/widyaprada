"use client";

import { Sidebar } from "@/components";
import { app, ui } from "@/config";
import Link from "next/link";

export default function Home() {
  return (
    <div className="flex h-screen bg-gray-50">
      <Sidebar />
      <main className="flex-1 overflow-y-auto">
        <div className="p-6 space-y-6">
          <div
            className="rounded-2xl shadow-sm p-6 flex items-center justify-between"
            style={{
              background: `linear-gradient(135deg, ${ui.theme.gradient.from}15, ${ui.theme.gradient.to}20)`,
            }}
          >
            <div>
              <h1 className="text-2xl font-bold text-gray-900 mb-2">
                Selamat Datang di {app.name}
              </h1>
              <p className="text-gray-600">{app.description}</p>
            </div>
            <div className="hidden md:block">
              <img src={ui.logo.src} alt={ui.logo.alt} className="h-16 w-auto opacity-90" />
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <Link
              href="/dashboard"
              className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md hover:border-blue-200 transition-all"
            >
              <h3 className="font-semibold text-gray-900 mb-1">Dashboard</h3>
              <p className="text-sm text-gray-500">Ringkasan dan statistik</p>
            </Link>
            <Link
              href="/wpdata"
              className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md hover:border-blue-200 transition-all"
            >
              <h3 className="font-semibold text-gray-900 mb-1">WPData</h3>
              <p className="text-sm text-gray-500">Manajemen data Widyaprada</p>
            </Link>
            <Link
              href="/wpujikom/bank-soal"
              className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md hover:border-blue-200 transition-all"
            >
              <h3 className="font-semibold text-gray-900 mb-1">Uji Kompetensi</h3>
              <p className="text-sm text-gray-500">Bank soal, paket soal, CBT</p>
            </Link>
          </div>
        </div>
      </main>
    </div>
  );
}
