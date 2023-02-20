package sort

import "strings"

func DirectionFromQuery(d string) (Direction, bool) {
	direction := Direction(strings.ToUpper(d))
	ok := true

	if direction != Desc && direction != Asc {
		ok = false
	}

	return direction, ok
}
