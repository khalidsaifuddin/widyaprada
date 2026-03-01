"use client";

import { useState } from "react";
import { apiService } from "@/lib/api";

interface ImageUrlInputProps {
  value: string;
  onChange: (url: string) => void;
  placeholder?: string;
  accept?: string;
  className?: string;
}

export default function ImageUrlInput({
  value,
  onChange,
  placeholder = "https://...",
  accept = "image/jpeg,image/png,image/gif,image/webp",
  className = "",
}: ImageUrlInputProps) {
  const [uploading, setUploading] = useState(false);
  const [uploadError, setUploadError] = useState("");

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    setUploadError("");
    setUploading(true);
    const formData = new FormData();
    formData.append("file", file);
    const res = await apiService.uploadFile<{ url?: string }>("v1/cms/upload-image", formData);
    setUploading(false);
    e.target.value = "";
    if (res.success && res.data?.url) {
      onChange(res.data.url);
    } else {
      setUploadError(res.message ?? "Gagal mengunggah");
    }
  };

  return (
    <div className={`space-y-2 ${className}`}>
      <div className="flex gap-2">
        <input
          type="text"
          value={value}
          onChange={(e) => {
            onChange(e.target.value);
            setUploadError("");
          }}
          placeholder={placeholder}
          className="flex-1 rounded-lg border border-gray-300 px-3 py-2 text-sm"
        />
        <label className="shrink-0 rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 cursor-pointer">
          {uploading ? "Mengunggah..." : "Unggah"}
          <input
            type="file"
            accept={accept}
            onChange={handleFileChange}
            disabled={uploading}
            className="hidden"
          />
        </label>
      </div>
      {uploadError && <p className="text-sm text-red-600">{uploadError}</p>}
    </div>
  );
}
