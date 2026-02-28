"use client";

import { ExclamationTriangleIcon, XMarkIcon } from "@heroicons/react/24/outline";
import React, { useEffect } from "react";

interface ErrorDialogProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  message: string | React.ReactNode;
  type?: "error" | "warning" | "info";
  disableClose?: boolean;
}

export default function ErrorDialog({
  isOpen,
  onClose,
  title,
  message,
  type = "error",
  disableClose = false,
}: ErrorDialogProps) {
  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === "Escape" && !disableClose) onClose();
    };
    if (isOpen) {
      document.addEventListener("keydown", handleEscape);
      document.body.style.overflow = "hidden";
    }
    return () => {
      document.removeEventListener("keydown", handleEscape);
      document.body.style.overflow = "";
    };
  }, [isOpen, onClose, disableClose]);

  if (!isOpen) return null;

  const iconColor =
    type === "error" ? "text-red-500" : type === "warning" ? "text-yellow-500" : "text-blue-500";
  const bgColor =
    type === "error" ? "bg-red-50 border-red-200" : type === "warning" ? "bg-yellow-50 border-yellow-200" : "bg-blue-50 border-blue-200";

  return (
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div
        className="fixed inset-0 bg-black/50"
        onClick={disableClose ? undefined : onClose}
        aria-hidden
      />
      <div className="flex min-h-full items-center justify-center p-4">
        <div className={`relative w-full max-w-md rounded-2xl border p-6 shadow-xl ${bgColor}`}>
          {!disableClose && (
            <button
              type="button"
              onClick={onClose}
              className="absolute right-4 top-4 rounded-full p-1 text-gray-400 hover:bg-gray-100 hover:text-gray-600"
            >
              <XMarkIcon className="h-5 w-5" />
            </button>
          )}
          <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-white shadow-lg mb-4">
            <ExclamationTriangleIcon className={`h-6 w-6 ${iconColor}`} />
          </div>
          <div className="text-center">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">{title}</h3>
            <div className="text-sm text-gray-600 mb-6">{message}</div>
            {!disableClose && (
              <button
                type="button"
                onClick={onClose}
                className="rounded-lg bg-gray-600 px-4 py-2 text-sm font-medium text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              >
                Tutup
              </button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
