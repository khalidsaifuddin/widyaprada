import { NextRequest, NextResponse } from "next/server";

const publicPaths = [
  "/",
  "/beranda",
  "/berita",
  "/jurnal",
  "/auth/login",
  "/auth/register",
  "/auth/forgot-password",
  "/auth/reset-password",
  "/api",
  "/_next",
  "/favicon.ico",
];

function isPublic(pathname: string): boolean {
  if (publicPaths.some((p) => pathname === p || pathname.startsWith(p + "/"))) return true;
  if (/\.(jpg|jpeg|png|gif|svg|ico|css|js|woff|woff2|ttf|eot)$/i.test(pathname)) return true;
  return false;
}

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  if (isPublic(pathname)) return NextResponse.next();

  const token = request.cookies.get("auth_token")?.value;
  if (!token) {
    const login = new URL("/auth/login", request.url);
    login.searchParams.set("redirect", pathname);
    return NextResponse.redirect(login);
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
};
