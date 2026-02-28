import { cn } from "@/lib/utils";
import React from "react";

interface CardProps {
  title?: string;
  subtitle?: string;
  children: React.ReactNode;
  className?: string;
  actions?: React.ReactNode;
}

export default function Card({ title, subtitle, children, className, actions }: CardProps) {
  return (
    <div className={cn("bg-white rounded-lg shadow-sm border border-gray-200", className)}>
      {(title || actions) && (
        <div className="flex items-center justify-between p-6 border-b border-gray-200">
          <div>
            {title && <h3 className="text-lg font-semibold text-gray-900">{title}</h3>}
            {subtitle && <p className="text-sm text-gray-500 mt-1">{subtitle}</p>}
          </div>
          {actions && <div className="flex items-center space-x-2">{actions}</div>}
        </div>
      )}
      <div className="p-6">{children}</div>
    </div>
  );
}
