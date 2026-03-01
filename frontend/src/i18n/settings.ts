export const fallbackLng = "id";
export const defaultNS = "common";
export const supportedLngs = ["id", "en"] as const;
export type SupportedLng = (typeof supportedLngs)[number];

export const getOptions = (lng = fallbackLng, ns = defaultNS) => ({
  supportedLngs,
  fallbackLng,
  fallbackNS: defaultNS,
  defaultNS,
  lng,
  ns,
});
