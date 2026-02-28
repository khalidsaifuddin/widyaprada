/**
 * Configuration for Aplikasi Widyaprada Frontend
 */

export interface AppConfig {
  api: {
    baseUrl: string;
    timeout: number;
  };
  app: {
    name: string;
    title: string;
    description: string;
    version: string;
    display_version: string;
  };
  ui: {
    logo: { src: string; alt: string };
    header: { title: string; subtitle?: string };
    theme: {
      primaryColor: string;
      secondaryColor: string;
      gradient: { from: string; to: string };
    };
  };
  env: {
    isDevelopment: boolean;
    isProduction: boolean;
    isTest: boolean;
  };
}

const defaultConfig: AppConfig = {
  api: {
    baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api",
    timeout: parseInt(process.env.NEXT_PUBLIC_API_TIMEOUT ?? "30000", 10),
  },
  app: {
    name: "Widyaprada",
    title: "Aplikasi Widyaprada",
    description: "Layanan penjaminan mutu pendidikan – PAUD, dasar, menengah, dan masyarakat",
    version: "0.1.0",
    display_version: "0.1.0",
  },
  ui: {
    logo: { src: "/logo-widyaprada.svg", alt: "Widyaprada" },
    header: {
      title: "Widyaprada",
      subtitle: "Penjaminan Mutu Pendidikan",
    },
    theme: {
      primaryColor: "#2563eb",
      secondaryColor: "#64748b",
      gradient: { from: "#1e40af", to: "#2563eb" },
    },
  },
  env: {
    isDevelopment: process.env.NODE_ENV === "development",
    isProduction: process.env.NODE_ENV === "production",
    isTest: process.env.NODE_ENV === "test",
  },
};

const getConfig = (): AppConfig => {
  const config = { ...defaultConfig };
  if (process.env.NEXT_PUBLIC_APP_NAME) config.app.name = process.env.NEXT_PUBLIC_APP_NAME;
  if (process.env.NEXT_PUBLIC_APP_TITLE) config.app.title = process.env.NEXT_PUBLIC_APP_TITLE;
  if (process.env.NEXT_PUBLIC_HEADER_TITLE) config.ui.header.title = process.env.NEXT_PUBLIC_HEADER_TITLE;
  if (process.env.NEXT_PUBLIC_HEADER_SUBTITLE) config.ui.header.subtitle = process.env.NEXT_PUBLIC_HEADER_SUBTITLE;
  return config;
};

export const appConfig = getConfig();
export const { api, app, ui, env } = appConfig;
