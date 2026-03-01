"use client";

import { ui } from "@/config";
import { HomeIcon, QuestionMarkCircleIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import { useTranslation } from "react-i18next";

export default function NotFound() {
  const { t } = useTranslation("common");
  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="flex justify-center">
          <img src={ui.logo.src} alt={ui.logo.alt} className="h-16 w-auto" />
        </div>
        <h2 className="mt-6 text-center text-3xl font-bold text-gray-900">
          {t("notFound.title")}
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          {t("notFound.message")}
        </p>
      </div>
      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="bg-white py-8 px-4 shadow-sm rounded-2xl sm:px-10">
          <div className="flex flex-col items-center">
            <div className="mx-auto flex h-24 w-24 items-center justify-center rounded-full bg-blue-100 mb-6">
              <QuestionMarkCircleIcon className="h-12 w-12 text-blue-600" />
            </div>
            <h1 className="text-6xl font-bold text-gray-300 mb-2">404</h1>
            <p className="text-lg text-gray-600 mb-8">{t("notFound.hint")}</p>
            <div className="flex gap-3 w-full sm:w-auto">
              <Link
                href="/"
                className="flex-1 sm:flex-initial flex justify-center items-center px-4 py-3 rounded-lg text-white bg-blue-600 hover:bg-blue-700 font-medium text-sm"
              >
                <HomeIcon className="h-5 w-5 mr-2" />
                {t("nav.home")}
              </Link>
              <button
                type="button"
                onClick={() => window.history.back()}
                className="flex-1 sm:flex-initial flex justify-center items-center px-4 py-3 rounded-lg border border-gray-300 text-gray-700 bg-white hover:bg-gray-50 font-medium text-sm"
              >
                {t("action.back")}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
