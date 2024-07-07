package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func get_linux_lang() string {
	locale := os.Getenv("LANG")
	if locale == "" {
		locale = os.Getenv("LC_ALL")
	}

	if locale == "" {
		return language.English.String()
	}

	parts := strings.Split(locale, ".")
	if len(parts) > 0 {
		locale = parts[0]
	}

	localeParts := strings.Split(locale, "_")
	if len(localeParts) == 2 {
		language := localeParts[0]
		region := localeParts[1]
		return fmt.Sprintf("%s-%s", language, region)
	} else {
		return localeParts[0]
	}
}

func localization_init() {

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("i18n/en.json")
	bundle.LoadMessageFile("i18n/pl.json")
	localizer = i18n.NewLocalizer(bundle, get_linux_lang())
}

func localized_week() [7]string {
	monday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "monday"})
	tuesday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "tuesday"})
	wednesday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "wednesday"})
	thursday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "thursday"})
	friday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "friday"})
	saturday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "saturday"})
	sunday, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "sunday"})

	return [7]string{monday, tuesday, wednesday, thursday, friday, saturday, sunday}
}
