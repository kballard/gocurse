package curses

// struct ldat{};
// struct _win_st{};
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
import "C"

import (
	"os"
	"unsafe"
)

type void unsafe.Pointer

type Window C.WINDOW

type CursesError struct {
	message string
}

func (ce CursesError) String() string {
	return ce.message
}

// Cursor options.
const (
	CURS_HIDE = iota
	CURS_NORM
	CURS_HIGH
)

// Pointers to the values in curses, which may change values.
var Cols *int = nil
var Rows *int = nil

var Colors *int = nil
var ColorPairs *int = nil

var Tabsize *int = nil

// The window returned from C.initscr()
var Stdwin *Window = nil

// Initializes gocurses
func init() {
	Cols = (*int)(void(&C.COLS))
	Rows = (*int)(void(&C.LINES))

	Colors = (*int)(void(&C.COLORS))
	ColorPairs = (*int)(void(&C.COLOR_PAIRS))

	Tabsize = (*int)(void(&C.TABSIZE))
}

func Initscr() (*Window, os.Error) {
	Stdwin = (*Window)(C.initscr())

	if Stdwin == nil {
		return nil, CursesError{"Initscr failed"}
	}

	return Stdwin, nil
}

func Newwin(rows int, cols int, starty int, startx int) (*Window, os.Error) {
	nw := (*Window)(C.newwin(C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

	if nw == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return nw, nil
}

func (win *Window) Subwin(rows int, cols int, starty int, startx int) (*Window, os.Error) {
	sw := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

	if sw == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return sw, nil
}

func (win *Window) Derwin(rows int, cols int, starty int, startx int) (*Window, os.Error) {
	dw := (*Window)(C.derwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

	if dw == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return dw, nil
}

func Start_color() os.Error {
	C.start_color()
	if int(C.has_colors()) == C.FALSE {
		return CursesError{"terminal does not support color"}
	}

	return nil
}

func Init_pair(pair int, fg int, bg int) os.Error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) != OK {
		return CursesError{"Init_pair failed"}
	}
	return nil
}

func Color_pair(pair int) int32 {
	return int32(C.COLOR_PAIR(C.int(pair)))
}

func Pair_number(flags int32) int {
	return int(C.PAIR_NUMBER(C.int(flags)))
}

func Noecho() os.Error {
	if int(C.noecho()) != OK {
		return CursesError{"Noecho failed"}
	}
	return nil
}

func Echo() os.Error {
	if int(C.noecho()) != OK {
		return CursesError{"Echo failed"}
	}
	return nil
}

func Curs_set(c int) os.Error {
	if C.curs_set(C.int(c)) != OK {
		return CursesError{"Curs_set failed"}
	}
	return nil
}

func Nocbreak() os.Error {
	if C.nocbreak() != OK {
		return CursesError{"Nocbreak failed"}
	}
	return nil
}

func Cbreak() os.Error {
	if C.cbreak() != OK {
		return CursesError{"Cbreak failed"}
	}
	return nil
}

func Endwin() os.Error {
	if C.endwin() != OK {
		return CursesError{"Endwin failed"}
	}
	return nil
}

func (win *Window) Color_set(color int16) os.Error {
	if C.color_set(C.short(color), nil) != OK {
		return CursesError{"Color_set failed"}
	}
	return nil
}

// Attr_on() and related functions use the WA_ attributes
func (win *Window) Attr_on(flags int32) {
	// manpage says return value is irrelevant
	C.wattr_on((*C.WINDOW)(win), C.attr_t(flags), nil)
}

func (win *Window) Attr_off(flags int32) {
	C.wattr_off((*C.WINDOW)(win), C.attr_t(flags), nil)
}

// sets both attributes and color
func (win *Window) Attr_set(flags int32, color int16) {
	C.wattr_set((*C.WINDOW)(win), C.attr_t(flags), C.short(color), nil)
}

func (win *Window) Attr_get() (flags int32, color int16) {
	var cflags C.attr_t
	var cpair C.short
	C.wattr_get((*C.WINDOW)(win), &cflags, &cpair, nil)
	return int32(cflags), int16(cpair)
}

func (win *Window) Chgat(n int, flags int32, color int16) {
	C.wchgat((*C.WINDOW)(win), C.int(n), C.attr_t(flags), C.short(color), nil)
}

func (win *Window) Mvchgat(y, x, n int, flags int32, color int16) {
	C.mvwchgat((*C.WINDOW)(win), C.int(y), C.int(x), C.int(n), C.attr_t(flags), C.short(color), nil)
}

func (win *Window) Getyx() (int, int) {
	return int(C.getcury((*C.WINDOW)(win))), int(C.getcurx((*C.WINDOW)(win)))
}

func (win *Window) Getparyx() (int, int) {
	return int(C.getpary((*C.WINDOW)(win))), int(C.getparx((*C.WINDOW)(win)))
}

func (win *Window) Getbegyx() (int, int) {
	return int(C.getbegy((*C.WINDOW)(win))), int(C.getbegx((*C.WINDOW)(win)))
}

func (win *Window) Getmaxyx() (int, int) {
	return int(C.getmaxy((*C.WINDOW)(win))), int(C.getmaxx((*C.WINDOW)(win)))
}

func (win *Window) Getch() int {
	return int(C.wgetch((*C.WINDOW)(win)))
}

func (win *Window) Inch() int {
	return int(C.winch((*C.WINDOW)(win)))
}

func (win *Window) Mvinch(y, x int) int {
	return int(C.mvwinch((*C.WINDOW)(win), C.int(y), C.int(x)))
}

func (win *Window) Addch(c int32, flags int32) {
	C.waddch((*C.WINDOW)(win), C.chtype(c)|C.chtype(flags))
}

func (win *Window) Mvaddch(y, x int, c int32, flags int32) {
	C.mvwaddch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)|C.chtype(flags))
}

func (win *Window) Insch(c int32, flags int32) {
	C.winsch((*C.WINDOW)(win), C.chtype(c)|C.chtype(flags))
}

func (win *Window) Mvinsch(y, x int, c int32, flags int32) {
	C.mvwinsch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)|C.chtype(flags))
}

func (win *Window) Delch() {
	C.wdelch((*C.WINDOW)(win))
}

func (win *Window) Mvdelch(y, x int) {
	C.mvwdelch((*C.WINDOW)(win), C.int(y), C.int(x))
}

// waddstr() doesn't support attributes, so mimic it by using waddch instead
func (win *Window) Addstr(str string, flags int32) {
	for i := 0; i < len(str); i++ {
		C.waddch((*C.WINDOW)(win), C.chtype(str[i])|C.chtype(flags))
	}
}

func (win *Window) Mvaddstr(y, x int, str string, flags int32) {
	win.Move(y, x)
	win.Addstr(str, flags)
}

// Normally Y is the first parameter passed in curses.
func (win *Window) Move(y, x int) {
	C.wmove((*C.WINDOW)(win), C.int(y), C.int(x))
}

func (w *Window) Keypad(tf bool) os.Error {
	var outint int
	if tf == true {
		outint = 1
	}
	if tf == false {
		outint = 0
	}
	if C.keypad((*C.WINDOW)(w), C.int(outint)) != OK {
		return CursesError{"Keypad failed"}
	}
	return nil
}

func (win *Window) Refresh() os.Error {
	if C.wrefresh((*C.WINDOW)(win)) != OK {
		return CursesError{"refresh failed"}
	}
	return nil
}

func (win *Window) Redrawln(beg_line, num_lines int) {
	C.wredrawln((*C.WINDOW)(win), C.int(beg_line), C.int(num_lines))
}

func (win *Window) Redraw() {
	C.redrawwin((*C.WINDOW)(win))
}

func (win *Window) Clear() {
	C.wclear((*C.WINDOW)(win))
}

func (win *Window) Erase() {
	C.werase((*C.WINDOW)(win))
}

func (win *Window) Clrtobot() {
	C.wclrtobot((*C.WINDOW)(win))
}

func (win *Window) Clrtoeol() {
	C.wclrtoeol((*C.WINDOW)(win))
}

func (win *Window) Box(verch, horch int) {
	C.box((*C.WINDOW)(win), C.chtype(verch), C.chtype(horch))
}

func (win *Window) Background(colour int32) {
	C.wbkgd((*C.WINDOW)(win), C.chtype(colour))
}
