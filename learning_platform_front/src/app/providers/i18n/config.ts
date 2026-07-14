import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";

import ru from "@/shared/locales/ru/translations.json";
import en from "@/shared/locales/en/translations.json";

i18n
    .use(LanguageDetector)
    .use(initReactI18next)
    .init({
        resources: {
            ru: {
                translation: ru,
            },
            en: {
                translation: en,
            },
        },

        fallbackLng: "ru",

        interpolation: {
            escapeValue: false,
        },
    });

export default i18n;