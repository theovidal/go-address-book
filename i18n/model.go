package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type I18n struct {
	Localizer *i18n.Localizer
}

func NewInstance(locale string) (instance I18n, err error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if _, err := bundle.LoadMessageFile("i18n/cli/en.toml"); err != nil {
		return instance, err
	}
	if _, err := bundle.LoadMessageFile("i18n/cli/fr.toml"); err != nil {
		return instance, err
	}

	localizer := i18n.NewLocalizer(bundle, locale)
	instance.Localizer = localizer

	return
}

func (instance I18n) T(key string, params map[string]interface{}) (message string, err error) {
	message, err = instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: params,
	})
	return
}
