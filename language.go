package carbon

import (
	"strconv"
	"strings"

	"github.com/golang-module/carbon/translation"
)

var (
	// 默认区域
	defaultLocale = "en"
)

// Language 定义 Language 结构体
type Language struct {
	locale    string            // 区域
	resources map[string]string // 资源
}

// NewLanguage 初始化 Language 结构体
func NewLanguage() *Language {
	return &Language{
		locale:    defaultLocale,
		resources: make(map[string]string),
	}
}

// SetLocale 设置区域
func (lang *Language) SetLocale(locale string) error {
	if len(lang.resources) != 0 {
		return nil
	}

	resources, ok := translation.Translations[locale]
	if !ok {
		return invalidLocaleError(locale)
	}

	lang.resources = resources
	lang.locale = locale
	return nil
}

// SetResources 设置资源
func (lang *Language) SetResources(resources map[string]string) {
	if len(lang.resources) == 0 {
		lang.resources = resources
		return
	}
	for k, v := range resources {
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		}
	}
}

// translate 翻译转换
func (lang *Language) translate(unit string, diff int64) string {
	if len(lang.resources) == 0 {
		lang.SetLocale(defaultLocale)
	}
	slice := strings.Split(lang.resources[unit], "|")
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(diff, 10), 1)
	}
	if diff > 1 {
		return strings.Replace(slice[1], "%d", strconv.FormatInt(diff, 10), 1)
	}
	return slice[0]
}
