export function resolveImageUrl(url: string): string {
  if (!url) return "";
  if (url.startsWith("http")) return url;
  const base = process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/api\/?$/, "") ?? "http://localhost:8080";
  return base + (url.startsWith("/") ? url : "/" + url);
}
