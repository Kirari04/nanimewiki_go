package templates

import (
	"ch/kirari/animeApi/helpers"
	"ch/kirari/animeApi/models"
	"errors"
	"strings"
)

var Langs = []string{
	"en",
}
var Data = models.TemplateBlocks{
	EmailRegister: map[string]models.Template{
		"en": {
			Exists: true,
			Data: map[string]string{
				"head": helpers.FileToString("./templates/email/register/en.head.txt"),
				"body": helpers.FileToString("./templates/email/register/en.body.html"),
				"txt":  helpers.FileToString("./templates/email/register/en.body.txt"),
			},
		},
	},
}

func Get(from map[string]models.Template, lang string) (map[string]string, error) {
	defaultLang := "en"
	if val, ok := from[lang]; ok {
		if val.Exists {
			return val.Data, nil
		}
	}
	if val, ok := from[defaultLang]; ok {
		if val.Exists {
			return val.Data, nil
		}
	}

	return nil, errors.New("Can't find response string to work with.")
}

func Prepare(value string, vars []models.TemplateVars) string {
	for _, v := range vars {
		value = strings.ReplaceAll(value, "{{"+v.Variable+"}}", v.Value)
	}
	return value
}
