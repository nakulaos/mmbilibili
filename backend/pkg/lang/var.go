/**
 ******************************************************************************
 * @file           : var.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package lang

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"sync"
)

//go:embed *.toml
var LocaleFS embed.FS

var bundle *i18n.Bundle

var _once sync.Once

func GetBundle() *i18n.Bundle {
	_once.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.LoadMessageFileFS(LocaleFS, "active.en.toml")
		bundle.LoadMessageFileFS(LocaleFS, "active.zh.toml")
	})
	return bundle
}

func init() {
	_once.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.LoadMessageFileFS(LocaleFS, "active.en.toml")
		bundle.LoadMessageFileFS(LocaleFS, "active.zh.toml")
	})
}
