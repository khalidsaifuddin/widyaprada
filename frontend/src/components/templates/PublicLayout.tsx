"use client";

import { app, ui } from "@/config";
import { getUserProfile } from "@/lib/auth";
import Link from "next/link";
import * as React from "react";
import { useTranslation } from "react-i18next";

export default function PublicLayout({ children }: { children: React.ReactNode }) {
  const { t } = useTranslation("common");
  return (
    <div className="min-h-screen flex flex-col bg-gray-50">
      <header
        className="sticky top-0 z-40 border-b border-gray-200 bg-white shadow-sm"
        style={{
          background: `linear-gradient(to right, ${ui.theme.gradient.from}, ${ui.theme.gradient.to})`,
        }}
      >
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex h-20 items-center justify-between">
            <Link href="/" className="flex items-center gap-3">
              <img src={ui.logo.src} alt={ui.logo.alt} className="h-10 w-auto" />
              <div>
                <span className="text-3xl font-bold">
                  <span style={{ color: "#057AC1" }}>Widya</span>
                  <span style={{ color: "#F9A702" }}>prada</span>
                </span>
                {ui.header.subtitle && (
                  <span className="ml-2 text-sm text-white/80 hidden sm:inline">
                    {ui.header.subtitle}
                  </span>
                )}
              </div>
            </Link>
            <nav className="flex items-center gap-4">
              <Link
                href="/"
                className="text-sm font-medium text-white/90 hover:text-white"
              >
                {t("nav.home")}
              </Link>
              <Link
                href="/berita"
                className="text-sm font-medium text-white/90 hover:text-white"
              >
                {t("nav.news")}
              </Link>
              <Link
                href="/jurnal"
                className="text-sm font-medium text-white/90 hover:text-white"
              >
                {t("nav.journal")}
              </Link>
              <LanguageSwitcher />
              <PublicHeaderAuth />
            </nav>
          </div>
        </div>
      </header>
      <main className="flex-1">{children}</main>
      <footer className="border-t border-gray-200 bg-gray-100 py-8">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row items-center justify-between gap-4">
            <div className="flex items-center gap-2">
              <img src={ui.logo.src} alt={ui.logo.alt} className="h-8 w-auto opacity-80" />
              <span className="text-gray-600 text-sm">{app.name}</span>
            </div>
            <p className="text-gray-500 text-sm text-center md:text-right">
              {t("app.descriptionLong")}
            </p>
            <div className="flex gap-6">
              <Link href="/berita" className="text-sm text-gray-600 hover:text-gray-900">
                {t("nav.news")}
              </Link>
              <Link href="/jurnal" className="text-sm text-gray-600 hover:text-gray-900">
                {t("nav.journal")}
              </Link>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

function LanguageSwitcher() {
  const { i18n: i18nInstance } = useTranslation();
  return (
    <div className="flex gap-1">
      {["id", "en"].map((lng) => (
        <button
          key={lng}
          type="button"
          onClick={() => i18nInstance.changeLanguage(lng)}
          className={`rounded px-2 py-1 text-xs font-medium ${i18nInstance.language === lng ? "bg-white/30 text-white" : "text-white/80 hover:text-white"
            }`}
        >
          {lng.toUpperCase()}
        </button>
      ))}
    </div>
  );
}

function PublicHeaderAuth() {
  const { t } = useTranslation("common");
  const [user, setUser] = React.useState<{ user_fullname?: string } | null>(null);

  React.useEffect(() => {
    getUserProfile().then(setUser).catch(() => setUser(null));
  }, []);

  if (user) {
    return (
      <Link
        href="/dashboard"
        className="rounded-lg bg-white/20 px-3 py-1.5 text-sm font-medium text-white hover:bg-white/30"
      >
        {t("nav.dashboard")}
      </Link>
    );
  }

  return (
    <Link
      href="/auth/login"
      className="rounded-lg bg-white px-3 py-1.5 text-sm font-medium text-[#033259] hover:bg-white/90"
    >
      {t("nav.login")}
    </Link>
  );
}
