// These characters are defined in curses.h
// They can't be constants because they go through a map

package curses

// #include "curses_chars.h"
import "C"

func ACS_ULCORNER() int32 {
	return int32(C.wrap_ACS_ULCORNER())
}

func ACS_LLCORNER() int32 {
	return int32(C.wrap_ACS_LLCORNER())
}

func ACS_URCORNER() int32 {
	return int32(C.wrap_ACS_URCORNER())
}

func ACS_LRCORNER() int32 {
	return int32(C.wrap_ACS_LRCORNER())
}

func ACS_LTEE() int32 {
	return int32(C.wrap_ACS_LTEE())
}

func ACS_RTEE() int32 {
	return int32(C.wrap_ACS_RTEE())
}

func ACS_BTEE() int32 {
	return int32(C.wrap_ACS_BTEE())
}

func ACS_TTEE() int32 {
	return int32(C.wrap_ACS_TTEE())
}

func ACS_HLINE() int32 {
	return int32(C.wrap_ACS_HLINE())
}

func ACS_VLINE() int32 {
	return int32(C.wrap_ACS_VLINE())
}

func ACS_PLUS() int32 {
	return int32(C.wrap_ACS_PLUS())
}

func ACS_S1() int32 {
	return int32(C.wrap_ACS_S1())
}

func ACS_S9() int32 {
	return int32(C.wrap_ACS_S9())
}

func ACS_DIAMOND() int32 {
	return int32(C.wrap_ACS_DIAMOND())
}

func ACS_CKBOARD() int32 {
	return int32(C.wrap_ACS_CKBOARD())
}

func ACS_DEGREE() int32 {
	return int32(C.wrap_ACS_DEGREE())
}

func ACS_PLMINUS() int32 {
	return int32(C.wrap_ACS_PLMINUS())
}

func ACS_BULLET() int32 {
	return int32(C.wrap_ACS_BULLET())
}

func ACS_LARROW() int32 {
	return int32(C.wrap_ACS_LARROW())
}

func ACS_RARROW() int32 {
	return int32(C.wrap_ACS_RARROW())
}

func ACS_DARROW() int32 {
	return int32(C.wrap_ACS_DARROW())
}

func ACS_UARROW() int32 {
	return int32(C.wrap_ACS_UARROW())
}

func ACS_BOARD() int32 {
	return int32(C.wrap_ACS_BOARD())
}

func ACS_LANTERN() int32 {
	return int32(C.wrap_ACS_LANTERN())
}

func ACS_BLOCK() int32 {
	return int32(C.wrap_ACS_BLOCK())
}
