import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  serverExternalPackages: [],
  trailingSlash: false,
  productionBrowserSourceMaps: false,
  poweredByHeader: false,
  images: {
    remotePatterns: [
      { protocol: "https", hostname: "**" },
      { protocol: "http", hostname: "localhost", pathname: "/**" },
    ],
  },
};

export default nextConfig;
