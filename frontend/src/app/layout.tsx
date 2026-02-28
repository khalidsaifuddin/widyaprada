import { AuthWrapper } from "@/components";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Widyaprada",
  description: "Aplikasi layanan penjaminan mutu pendidikan – PAUD, dasar, menengah, dan masyarakat",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="id">
      <body className="min-h-screen bg-gray-50" suppressHydrationWarning>
        <AuthWrapper>{children}</AuthWrapper>
      </body>
    </html>
  );
}
