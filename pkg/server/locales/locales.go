package locales

import "golang.org/x/text/language"

func GetLocale(acceptLangHeader string) map[string]string {
	langBase := getPreferredLang(acceptLangHeader)
	switch langBase {
	case language.English.String():
		return locale_en
	case language.Chinese.String():
		return locale_zh
	case language.Japanese.String():
		return locale_ja
	case language.Korean.String():
		return locale_ko
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
	language.English,  // en - first language is fallback
	language.Chinese,  // zh
	language.Japanese, // ja
	language.Korean,   // ko
}

var locale_en = map[string]string{
	"LogInToYourAccount": "Log in to your account",
	"Username":           "Username",
	"Password":           "Password",
	"LogIn":              "Log in",
	"WelcomeTo":          "Welcome to",
	"LogInWith":          "Log in with",
	"Error":              "Error",
}

var locale_zh = map[string]string{
	"LogInToYourAccount": "登录到您的帐户",
	"Username":           "用户名",
	"Password":           "密码",
	"LogIn":              "登录",
	"WelcomeTo":          "欢迎使用",
	"LogInWith":          "登录使用",
	"Error":              "错误",
}

var locale_ja = map[string]string{
	"LogInToYourAccount": "アカウントにログイン",
	"Username":           "ユーザー名",
	"Password":           "パスワード",
	"LogIn":              "ログイン",
	"WelcomeTo":          "ようこそ:",
	"LogInWith":          "ログイン:",
	"Error":              "エラー",
}

var locale_ko = map[string]string{
	"LogInToYourAccount": "귀하의 계정에 로그인하십시오",
	"Username":           "사용자 이름",
	"Password":           "암호",
	"LogIn":              "로그인",
	"WelcomeTo":          "환영합니다",
	"LogInWith":          "로그인",
	"Error":              "오류",
}
