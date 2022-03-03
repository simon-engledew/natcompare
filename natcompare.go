package natcompare

import (
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(\d+|\D+)`)

// Less can be used to sort string values that contain numbers.
func Less(a, b string) bool {
	for {
		if a == "" || b == "" {
			return len(a) < len(b)
		}

		posA := pattern.FindStringIndex(a)
		posB := pattern.FindStringIndex(b)

		chunkA := a[posA[0]:posA[1]]
		chunkB := b[posB[0]:posB[1]]

		a = a[posA[1]:]
		b = b[posB[1]:]

		if strings.Compare(chunkA, chunkB) == 0 {
			continue
		}

		aInt, aErr := strconv.Atoi(chunkA)
		bInt, bErr := strconv.Atoi(chunkB)

		isNumericA := aErr == nil
		isNumericB := bErr == nil

		if isNumericA || isNumericB {
			if aInt == bInt {
				continue
			}
			if isNumericA && isNumericB {
				return aInt < bInt
			}
			return isNumericA
		}

		return strings.Compare(chunkA, chunkB) < 0
	}
}
