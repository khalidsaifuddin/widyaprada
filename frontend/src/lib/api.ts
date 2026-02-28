import { getAuthData } from "./auth";
import { SecurityUtils, VALIDATION_PATTERNS } from "./security";

export interface ApiResponse<T = unknown> {
  data: T;
  message?: string;
  success: boolean;
  code?: number;
}

const rateLimiter = SecurityUtils.createRateLimiter(100, 60000);

class ApiService {
  private baseUrl: string;
  private defaultHeaders: Record<string, string> = {
    "Content-Type": "application/json",
    Accept: "application/json",
    "X-Requested-With": "XMLHttpRequest",
  };

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl.replace(/\/$/, "");
  }

  private sanitizeParams(params: Record<string, unknown>): Record<string, unknown> {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(params)) {
      if (v === undefined || v === null) continue;
      if (typeof v === "string") out[k] = SecurityUtils.sanitizeInput(v);
      else out[k] = v;
    }
    return out;
  }

  private async authHeaders(): Promise<Record<string, string>> {
    const auth = await getAuthData();
    const h = { ...this.defaultHeaders };
    if (auth?.access_token) h["Authorization"] = `Bearer ${auth.access_token}`;
    return h;
  }

  private checkRateLimit(): boolean {
    return rateLimiter("api");
  }

  private async handleResponse<T>(res: Response): Promise<ApiResponse<T>> {
    const contentType = res.headers.get("content-type");
    if (contentType && !contentType.includes("application/json")) {
      return { success: false, data: null as T, message: "Invalid response type" };
    }
    const data = await res.json().catch(() => ({}));
    return {
      success: res.ok,
      data: data.data ?? data,
      message: data.message,
      code: data.code ?? res.status,
    };
  }

  async get<T>(endpoint: string, params?: Record<string, unknown>): Promise<ApiResponse<T>> {
    if (!this.checkRateLimit()) {
      return { success: false, data: null as T, message: "Rate limit exceeded" };
    }
    const path = endpoint.startsWith("/") ? endpoint.slice(1) : endpoint;
    const url = new URL(path, this.baseUrl + "/");
    if (params) {
      const safe = this.sanitizeParams(params);
      Object.entries(safe).forEach(([k, v]) => url.searchParams.set(k, String(v)));
    }
    const headers = await this.authHeaders();
    const res = await fetch(url.toString(), { method: "GET", headers, credentials: "include" });
    return this.handleResponse<T>(res);
  }

  async post<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
    if (!this.checkRateLimit()) {
      return { success: false, data: null as T, message: "Rate limit exceeded" };
    }
    const path = endpoint.startsWith("/") ? endpoint.slice(1) : endpoint;
    const url = new URL(path, this.baseUrl + "/");
    const headers = await this.authHeaders();
    const res = await fetch(url.toString(), {
      method: "POST",
      headers,
      body: body ? JSON.stringify(body) : undefined,
      credentials: "include",
    });
    return this.handleResponse<T>(res);
  }

  async put<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
    if (!this.checkRateLimit()) {
      return { success: false, data: null as T, message: "Rate limit exceeded" };
    }
    const path = endpoint.startsWith("/") ? endpoint.slice(1) : endpoint;
    const url = new URL(path, this.baseUrl + "/");
    const headers = await this.authHeaders();
    const res = await fetch(url.toString(), {
      method: "PUT",
      headers,
      body: body ? JSON.stringify(body) : undefined,
      credentials: "include",
    });
    return this.handleResponse<T>(res);
  }

  async delete<T>(endpoint: string, body?: unknown): Promise<ApiResponse<T>> {
    if (!this.checkRateLimit()) {
      return { success: false, data: null as T, message: "Rate limit exceeded" };
    }
    const path = endpoint.startsWith("/") ? endpoint.slice(1) : endpoint;
    const url = new URL(path, this.baseUrl + "/");
    const headers = await this.authHeaders();
    const res = await fetch(url.toString(), {
      method: "DELETE",
      headers,
      body: body ? JSON.stringify(body) : undefined,
      credentials: "include",
    });
    return this.handleResponse<T>(res);
  }
}

const api = new ApiService(process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api");

export const apiService = {
  get: <T>(endpoint: string, params?: Record<string, unknown>) => api.get<T>(endpoint, params),
  post: <T>(endpoint: string, body?: unknown) => api.post<T>(endpoint, body),
  put: <T>(endpoint: string, body?: unknown) => api.put<T>(endpoint, body),
  delete: <T>(endpoint: string, body?: unknown) => api.delete<T>(endpoint, body),
};

export function extractApiData<T>(response: ApiResponse<unknown>, fallback: T): T {
  if (!response?.data) return fallback;
  const d = response.data as Record<string, unknown>;
  return (d?.data ?? response.data) as T;
}

export function formatCacheTime(cacheTime: string): string {
  try {
    return new Date(cacheTime).toLocaleDateString("id-ID", {
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return "";
  }
}

// Re-export for components that need patterns
export { VALIDATION_PATTERNS };
