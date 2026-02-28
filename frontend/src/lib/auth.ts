// Authentication utilities for Widyaprada

export interface UserRole {
  user_id: string;
  instansi_id: string;
  nama_instansi: string;
  role_aplikasi_id: string;
  role_aplikasi: string;
  [key: string]: unknown;
}

export interface AuthData {
  access_token: string;
  user_id: string;
  user_name: string;
  user_nik: string;
  user_fullname: string;
  expiry: string;
  role_user: UserRole[];
  [key: string]: unknown;
}

export interface UserProfile {
  user_id: string;
  user_name: string;
  user_nik: string;
  user_fullname: string;
  role_user: UserRole[];
  stored_at: string;
  [key: string]: unknown;
}

const AUTH_COOKIE_KEY = "auth_token";
const STORAGE_KEY = "current_auth";
const DB_NAME = "WidyapradaDB";

// IndexedDB
async function initDB(): Promise<IDBDatabase> {
  return new Promise((resolve, reject) => {
    const request = indexedDB.open(DB_NAME, 1);
    request.onerror = () => reject(request.error);
    request.onsuccess = () => resolve(request.result);
    request.onupgradeneeded = (event) => {
      const db = (event.target as IDBOpenDBRequest).result;
      if (!db.objectStoreNames.contains("auth")) {
        db.createObjectStore("auth", { keyPath: "id" });
      }
      if (!db.objectStoreNames.contains("user")) {
        db.createObjectStore("user", { keyPath: "user_id" });
      }
    };
  });
}

async function storeInDB(data: AuthData): Promise<boolean> {
  try {
    const db = await initDB();
    const tx = db.transaction(["auth", "user"], "readwrite");
    const authStore = tx.objectStore("auth");
    authStore.put({ id: "current_auth", ...data, stored_at: new Date().toISOString() });
    const userStore = tx.objectStore("user");
    userStore.put({
      user_id: data.user_id,
      user_name: data.user_name,
      user_fullname: data.user_fullname,
      role_user: data.role_user ?? [],
      stored_at: new Date().toISOString(),
    });
    return new Promise((res) => {
      tx.oncomplete = () => res(true);
      tx.onerror = () => res(false);
    });
  } catch {
    return false;
  }
}

async function getFromDB(): Promise<AuthData | null> {
  try {
    const db = await initDB();
    return new Promise((resolve) => {
      const tx = db.transaction("auth", "readonly");
      const req = tx.objectStore("auth").get("current_auth");
      req.onsuccess = () => resolve(req.result ?? null);
      req.onerror = () => resolve(null);
    });
  } catch {
    return null;
  }
}

async function clearDB(): Promise<boolean> {
  try {
    const db = await initDB();
    const tx = db.transaction(["auth", "user"], "readwrite");
    tx.objectStore("auth").clear();
    tx.objectStore("user").clear();
    return new Promise((res) => {
      tx.oncomplete = () => res(true);
      tx.onerror = () => res(false);
    });
  } catch {
    return false;
  }
}

export const setAuthCookie = (token: string): void => {
  if (typeof document === "undefined") return;
  const expiry = new Date();
  expiry.setDate(expiry.getDate() + 7);
  document.cookie = `${AUTH_COOKIE_KEY}=${encodeURIComponent(token)}; path=/; expires=${expiry.toUTCString()}; SameSite=Strict`;
};

export const clearAuthCookie = (): void => {
  if (typeof document === "undefined") return;
  document.cookie = `${AUTH_COOKIE_KEY}=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT`;
};

export const getAuthCookie = (): string | null => {
  if (typeof document === "undefined") return null;
  const match = document.cookie.match(new RegExp("(^| )" + AUTH_COOKIE_KEY + "=([^;]+)"));
  return match ? decodeURIComponent(match[2]) : null;
};

export const storeAuthData = async (authData: AuthData): Promise<void> => {
  const ok = await storeInDB(authData);
  if (!ok) throw new Error("Failed to store auth data");
  setAuthCookie(authData.access_token);
};

export const getAuthData = async (): Promise<AuthData | null> => {
  const data = await getFromDB();
  if (!data?.access_token) return null;
  if (data.expiry && new Date(data.expiry) <= new Date()) {
    await logout();
    return null;
  }
  return data;
};

export const getUserProfile = async (): Promise<UserProfile | null> => {
  const authData = await getAuthData();
  if (!authData) return null;
  return {
    user_id: authData.user_id,
    user_name: authData.user_name,
    user_nik: authData.user_nik ?? "",
    user_fullname: authData.user_fullname,
    role_user: authData.role_user ?? [],
    stored_at: new Date().toISOString(),
  };
};

export const isLoggedIn = async (): Promise<boolean> => {
  const data = await getAuthData();
  return data !== null;
};

export const logout = async (): Promise<void> => {
  await clearDB();
  clearAuthCookie();
  if (typeof window !== "undefined") window.location.href = "/auth/login";
};

export const detectStorageIssues = (): { issues: string[] } => {
  const issues: string[] = [];
  if (typeof indexedDB === "undefined") issues.push("IndexedDB not available");
  try {
    localStorage.setItem("_t", "1");
    localStorage.removeItem("_t");
  } catch {
    issues.push("localStorage not available");
  }
  return { issues };
};

export const showStorageWarning = (issues: string[]): void => {
  if (issues.length && typeof window !== "undefined") {
    console.warn("Storage issues:", issues);
  }
};
