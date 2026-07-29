import { useEffect, useRef } from "react"
import { useTranslation } from "react-i18next"
import { useTheme } from "@teispace/next-themes/client"
import { useAppSelector } from "@/app/providers/storeProvider/hooks/hooks"
import { getUserFullData } from "@/entities/user"
import { UserThemeEnum } from "@/shared/enums/user"

export function UserPreferencesSync() {
    const userData = useAppSelector(getUserFullData)
    const { setTheme } = useTheme()
    const { i18n } = useTranslation()
    const appliedKey = useRef<string | null>(null)

    useEffect(() => {
        if (!userData) {
            appliedKey.current = null
            return
        }

        const theme = userData.userSettings.theme
        const language = userData.userSettings.language
        const key = `${theme}:${language}`

        if (appliedKey.current === key) return
        appliedKey.current = key

        setTheme(theme === UserThemeEnum.DARK ? "dark" : "light")

        const nextLang = language.toLowerCase()
        const currentLang = (i18n.resolvedLanguage ?? i18n.language).toLowerCase()
        if (!currentLang.startsWith(nextLang)) {
            void i18n.changeLanguage(nextLang)
        }
    }, [userData, setTheme, i18n])

    return null
}
