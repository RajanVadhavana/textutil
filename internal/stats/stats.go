package stats

import (
	"strings"
	"github.com/RajanVadhavana/textutil/internal/clean"
)

func WordCount(s string) (int) {
	return len(strings.Fields(clean.Normalize(s)))
}