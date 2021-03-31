package login

import "golang.org/x/text/language"

func getLocale(langBase string) map[string]string {
	switch langBase {
	case language.English.String():
		return locale_en
	// Uncomment once given locale is available
	// case language.Chinese.String():
	// 	return locale_zh
	// case language.Japanese.String():
	// 	return locale_ja
	// case language.Korean.String():
	// 	return locale_ko
	default:
		return locale_en
	}
}

func getPreferredLang(acceptLangHeader string) string {
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
	language.English, // en - first language is fallback
	// Uncomment once given locale is available
	// language.Chinese,  // zh
	// language.Japanese, // ja
	// language.Korean,   // ko
}

var locale_en = map[string]string{
	"LogInToYourAccount": "Log in to your account",
	"Username":           "Username",
	"Password":           "Password",
	"LogIn":              "Log in",
	"LoginTitle":         "Login",
	"WelcomeTo":          "Welcome to",
	"LogInWith":          "Log in with",
	"Error":              "Error",
}

// Uncomment once given locale is available
// var locale_zh = map[string]string{
// 	"": "",
// }

// var locale_ja = map[string]string{
// 	"": "",
// }

// var locale_ko = map[string]string{
// 	"": "",
// }
