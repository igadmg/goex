package stringsex

import (
	"iter"
	"slices"
	"strings"
)

func JoinSeq(x iter.Seq[string], sep string) string {
	return strings.Join(slices.Collect(x), sep)
}
