package telnetrus

import (
	"errors"
	"golang.org/x/text/encoding/charmap"
	"strings"
)

const (
	Windows1251 = "1"
	KOI8        = "2"
)

func ToUTF8(encodingName string, bytes []byte) (result string, err error) {
	characterMap, err := GetCharacterMap(encodingName)

	if err != nil {
		return
	}

	result, _ = characterMap.NewDecoder().String(string(bytes))

	return
}

func FromUTF8(encodingName string, string string) (resultBytes []byte, err error) {
	characterMap, err := GetCharacterMap(encodingName)

	if err != nil {
		return
	}

	// doubling "я" letter for CP1251 encoding
	// @see http://citforum.ru/nets/semenov/4/45/tlnt_453.shtml
	if encodingName == Windows1251 {
		string = strings.Replace(string, "я", "яя", -1)
	}

	result, _ := characterMap.NewEncoder().String(string)
	resultBytes = []byte(result)

	return
}

func GetCharacterMap(characterMapName string) (characterMap *charmap.Charmap, err error) {
	characterMaps := map[string]*charmap.Charmap{
		Windows1251: charmap.Windows1251,
		KOI8:        charmap.KOI8R,
	}

	if charMapName, ok := characterMaps[characterMapName]; ok {
		characterMap = charMapName
	} else {
		err = errors.New("encoding is not supported: " + characterMapName)
	}

	return
}
