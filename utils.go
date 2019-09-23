package inforusms

import "strconv"

func strToInt(str string, bits int) int64 {
	i, err := strconv.ParseInt(str, 10, bits)
	if err != nil {
		return 0
	}

	return i
}

func strToUint(str string, bits int) uint64 {
	i, err := strconv.ParseUint(str, 10, bits)
	if err != nil {
		return 0
	}

	return i
}
