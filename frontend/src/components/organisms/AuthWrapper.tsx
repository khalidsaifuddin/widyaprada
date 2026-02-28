"use client";

import { getAuthCookie, isLoggedIn } from "@/lib/auth";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function AuthWrapper({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const pathname = usePathname();
  const [loading, setLoading] = useState(true);
  const [authenticated, setAuthenticated] = useState(false);

  useEffect(() => {
    if (pathname === "/auth/login") {
      setAuthenticated(true);
      setLoading(false);
      return;
    }
    const ext = [".jpg", ".jpeg", ".png", ".gif", ".svg", ".ico", ".css", ".js"];
    if (ext.some((e) => pathname.toLowerCase().endsWith(e))) {
      setAuthenticated(true);
      setLoading(false);
      return;
    }

    const check = async () => {
      try {
        if (!getAuthCookie()) {
          router.push(`/auth/login?redirect=${encodeURIComponent(pathname)}`);
          return;
        }
        const ok = await isLoggedIn();
        if (!ok) {
          router.push(`/auth/login?redirect=${encodeURIComponent(pathname)}`);
          return;
        }
        setAuthenticated(true);
      } catch {
        router.push(`/auth/login?redirect=${encodeURIComponent(pathname)}`);
      } finally {
        setLoading(false);
      }
    };
    check();
  }, [pathname, router]);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto" />
          <p className="mt-4 text-gray-600">Memverifikasi autentikasi...</p>
        </div>
      </div>
    );
  }

  if (!authenticated) return null;
  return <>{children}</>;
}
