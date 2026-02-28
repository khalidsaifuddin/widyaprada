// Security utilities for Widyaprada

export const VALIDATION_PATTERNS = {
  USERNAME: /^[a-zA-Z0-9_]{3,50}$/,
  PASSWORD: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{6,}$/,
  EMAIL: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
  SAFE_TEXT: /^[^<>{}()\[\]]*$/,
  URL: /^https?:\/\/[^\s/$.?#].[^\s]*$/i,
};

export class SecurityUtils {
  static sanitizeHtml(input: string): string {
    if (!input) return "";
    return input
      .replace(/<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi, "")
      .replace(/javascript:/gi, "")
      .replace(/on\w+\s*=/gi, "");
  }

  static sanitizeInput(input: string): string {
    if (!input || typeof input !== "string") return "";
    return this.sanitizeHtml(input.trim());
  }

  static createRateLimiter(maxRequests: number, windowMs: number) {
    const requests = new Map<string, { count: number; resetTime: number }>();
    return (identifier: string): boolean => {
      const now = Date.now();
      const userRequests = requests.get(identifier);
      if (!userRequests || now > userRequests.resetTime) {
        requests.set(identifier, { count: 1, resetTime: now + windowMs });
        return true;
      }
      if (userRequests.count >= maxRequests) return false;
      userRequests.count++;
      return true;
    };
  }

  static validateFile(file: File): { isValid: boolean; error?: string } {
    const maxSize = 5 * 1024 * 1024;
    const allowedTypes = ["image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp"];
    if (file.size > maxSize) return { isValid: false, error: "File size exceeds 5MB limit" };
    if (!allowedTypes.includes(file.type)) return { isValid: false, error: "Invalid file type" };
    return { isValid: true };
  }
}

export const validateInput = (schema: Record<string, RegExp>) => {
  return (data: Record<string, unknown>) => {
    const errors: Record<string, string> = {};
    for (const [field, pattern] of Object.entries(schema)) {
      const value = data[field];
      if (value !== undefined && value !== null && value !== "") {
        const inputStr = String(value);
        if (!pattern.test(inputStr)) errors[field] = `Invalid ${field} format`;
      }
    }
    return { isValid: Object.keys(errors).length === 0, errors };
  };
};
