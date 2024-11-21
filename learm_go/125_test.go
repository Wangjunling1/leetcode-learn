package learn_go

import (
	"strings"
	"testing"
)

func Test125(t *testing.T) {
	t.Log(f125("A man, a plan, a canal: Panama"))
	t.Log(f125("0P"))
	t.Log(f125("9"))
	t.Log("_"[0])
	t.Log("z"[0])
	t.Log("a"[0])
	t.Log("Z"[0])

}

func f125(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i <= j {
		if (57 < s[i] && s[i] < 65) || s[i] < 48 || s[i] > 122 || (90 < s[i] && s[i] < 97) {
			i++
			continue
		}
		if (57 < s[j] && s[j] < 65) || s[j] < 48 || s[j] > 122 || (90 < s[j] && s[j] < 97) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
