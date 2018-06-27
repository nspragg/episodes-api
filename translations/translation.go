package translation

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nspragg/episodes-api/translations/model"
)

var translations = make(map[string]interface{})
var translationTemplate = regexp.MustCompile(`#{(.*?)}`)

const (
	// TODO
	translationsDir = "./locales/"
)

func init() {
	files, _ := ioutil.ReadDir(translationsDir)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			abs, err := filepath.Abs(filepath.Join(translationsDir, f.Name()))
			if err != nil {
				panic("Failure deriving abs for: " + f.Name())
			}
			file, err := os.Open(abs)
			if err != nil {
				panic("Failure opening translation config: " + file.Name())
			}
			defer file.Close()

			lang := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			config := model.NewModel(lang)

			decoder := json.NewDecoder(file)
			if err := decoder.Decode(config); err != nil {
				panic("Failure decoding: " + f.Name())
			}

			translations[lang] = config
		}
	}
}

func Translate(data string, lang string) ([]string, error) {

	translated := make([]string, 0)
	matches := translationTemplate.FindAllStringSubmatch(data, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 0 {
				// TODO:
			}
		}
	}
	return translated, nil
}
