import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  serverExternalPackages: [],
  trailingSlash: false,
  productionBrowserSourceMaps: false,
  poweredByHeader: false,
  images: {
    domains: ["localhost"],
    remotePatterns: [{ protocol: "https", hostname: "**" }],
  },
};

export default nextConfig;
