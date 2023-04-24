package cleaner

import (
	"fmt"
	"regexp"
	"strings"
)

var spaceRegexp *regexp.Regexp

func init() {
	spaceRegexp = regexp.MustCompile(`\s+`)
}

func Clean(str string) string {
	str = strings.ToLower(str)
	cleaned := ""

	for _, r := range str {
		switch r {
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57: //0 to 9 unchanged
			cleaned += string(r)
		case 98, 100, 102, 103, 104, 106, 107, 108, 109, 112, 113, 114, 115, 116, 118, 119, 120, 122: //letters unchanged
			cleaned += string(r)
		case 97, 224, 225, 226, 227, 228, 229, 257, 259, 261: //"a" variants
			cleaned += "a"
		case 99, 231: //"c" variants
			cleaned += "c"
		case 101, 232, 233, 234, 235: //"e" variants
			cleaned += "e"
		case 105, 236, 237, 238, 239: //"i" variants
			cleaned += "i"
		case 110, 241: //"n" variants
			cleaned += "n"
		case 111, 242, 243, 244, 245, 246: //"o" variants
			cleaned += "o"
		case 117, 249, 250, 251, 252: //"u" variants
			cleaned += "u"
		case 121, 255: //"y" variants
			cleaned += "y"
		default:
			cleaned += " "
		}
	}
	fmt.Println([]rune(cleaned))
	return strings.TrimSpace(spaceRegexp.ReplaceAllString(strings.ToUpper(string(cleaned)), " "))

}
