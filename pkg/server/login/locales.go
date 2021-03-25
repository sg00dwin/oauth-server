package login

import "golang.org/x/text/language"

func getLocale(langBase string) map[string]string {
	switch langBase {
	case language.English.String():
		return locale_en
	case language.Chinese.String():
		return locale_zh
	case language.Japanese.String():
		return locale_jp
	case language.Korean.String():
		return locale_ko
	default:
		return locale_en
	}
}

func getPrefferedLang(acceptLangHeader string) string {
	matcher := language.NewMatcher(supportedLangs)
	userPrefs, _, err := language.ParseAcceptLanguage(acceptLangHeader)
	if err != nil {
		// if error occurs, fallback to English
		return language.English.String()
	}
	tag, _, _ := matcher.Match(userPrefs...)
	base, _ := tag.Base()
	return base.String()
}

var supportedLangs = []language.Tag{
	language.English,  // en - first language is fallback
	language.Chinese,  // zh
	language.Japanese, // jp
	language.Korean,   // ko
}

var locale_en = map[string]string{
	"": "",
}

var locale_zh = map[string]string{
	"": "",
}

var locale_jp = map[string]string{
	"": "",
}

var locale_ko = map[string]string{
	"": "",
}
