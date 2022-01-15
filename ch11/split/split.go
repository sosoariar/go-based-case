package split

import "strings"

func Split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+1:]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
