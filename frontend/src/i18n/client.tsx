"use client";

import i18n from "i18next";
import resourcesToBackend from "i18next-resources-to-backend";
import { I18nextProvider, initReactI18next } from "react-i18next";
import { getOptions } from "./settings";

const initI18n = () => {
  if (i18n.isInitialized) return i18n;

  i18n
    .use(
      resourcesToBackend(
        (lng: string, ns: string) => import(`@/locales/${lng}/${ns}.json`)
      )
    )
    .use(initReactI18next)
    .init({
      ...getOptions(),
      react: { useSuspense: false },
      interpolation: { escapeValue: false },
    });

  return i18n;
};

initI18n();

export function I18nProvider({
  children,
  lng = "id",
}: {
  children: React.ReactNode;
  lng?: string;
}) {
  if (typeof window !== "undefined" && i18n.language !== lng) {
    i18n.changeLanguage(lng);
  }
  return <I18nextProvider i18n={i18n}>{children}</I18nextProvider>;
}
