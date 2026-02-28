"use client";

import { getUserProfile } from "@/lib/auth";
import { useEffect, useState } from "react";
import AssignmentBlock from "@/components/organisms/AssignmentBlock";
import JournalBlock from "@/components/organisms/JournalBlock";

export default function DashboardPage() {
  const [userName, setUserName] = useState<string | null>(null);

  useEffect(() => {
    getUserProfile().then((p) => setUserName(p?.user_fullname ?? null));
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">
          {userName ? `Selamat datang, ${userName}` : "Dashboard"}
        </h1>
        <p className="text-gray-600 mt-1">
          Ringkasan tugas uji kompetensi dan jurnal Anda
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="order-1">
          <AssignmentBlock />
        </div>
        <div className="order-2">
          <JournalBlock />
        </div>
      </div>
    </div>
  );
}
