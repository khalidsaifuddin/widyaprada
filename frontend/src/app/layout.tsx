import { AuthWrapper } from "@/components";
import { I18nProvider } from "@/i18n/client";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Widyaprada",
  description: "Aplikasi Portal Widyaprada – Kemendikdasmen",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="id">
      <body className="min-h-screen bg-gray-50" suppressHydrationWarning>
        <I18nProvider lng="id">
          <AuthWrapper>{children}</AuthWrapper>
        </I18nProvider>
      </body>
    </html>
  );
}
