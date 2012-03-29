#include <curses.h>

enum {
	$ERR = ERR,
	$OK = OK,

	$A_ALTCHARSET = A_ALTCHARSET,
	$A_BLINK = A_BLINK,
	$A_BOLD = A_BOLD,
	$A_DIM = A_DIM,
	$A_INVIS = A_INVIS,
	$A_NORMAL = A_NORMAL,
	$A_PROTECT = A_PROTECT,
	$A_REVERSE = A_REVERSE,
	$A_STANDOUT = A_STANDOUT,
	$A_UNDERLINE = A_UNDERLINE,
	$A_ATTRIBUTES = A_ATTRIBUTES,
	$A_CHARTEXT = A_CHARTEXT,
	$A_COLOR = A_COLOR,

	$WA_ALTCHARSET = WA_ALTCHARSET,
	$WA_BLINK = WA_BLINK,
	$WA_BOLD = WA_BOLD,
	$WA_DIM = WA_DIM,
	$WA_INVIS = WA_INVIS,
	$WA_LEFT = WA_LEFT,
	$WA_NORMAL = WA_NORMAL,
	$WA_PROTECT = WA_PROTECT,
	$WA_REVERSE = WA_REVERSE,
	$WA_RIGHT = WA_RIGHT,
	$WA_STANDOUT = WA_STANDOUT,
	$WA_UNDERLINE = WA_UNDERLINE,

	$COLOR_BLACK = COLOR_BLACK,
	$COLOR_BLUE = COLOR_BLUE,
	$COLOR_GREEN = COLOR_GREEN,
	$COLOR_CYAN = COLOR_CYAN,
	$COLOR_RED = COLOR_RED,
	$COLOR_MAGENTA = COLOR_MAGENTA,
	$COLOR_YELLOW = COLOR_YELLOW,
	$COLOR_WHITE = COLOR_WHITE,

	$KEY_BREAK = KEY_BREAK,
	$KEY_DOWN = KEY_DOWN,
	$KEY_UP = KEY_UP,
	$KEY_LEFT = KEY_LEFT,
	$KEY_RIGHT = KEY_RIGHT,
	$KEY_HOME = KEY_HOME,
	$KEY_BACKSPACE = KEY_BACKSPACE,
	$KEY_F0 = KEY_F0,
	$KEY_DL = KEY_DL,
	$KEY_IL = KEY_IL,
	$KEY_DC = KEY_DC,
	$KEY_IC = KEY_IC,
	$KEY_EIC = KEY_EIC,
	$KEY_CLEAR = KEY_CLEAR,
	$KEY_EOS = KEY_EOS,
	$KEY_EOL = KEY_EOL,
	$KEY_SF = KEY_SF,
	$KEY_SR = KEY_SR,
	$KEY_NPAGE = KEY_NPAGE,
	$KEY_PPAGE = KEY_PPAGE,
	$KEY_STAB = KEY_STAB,
	$KEY_CTAB = KEY_CTAB,
	$KEY_CATAB = KEY_CATAB,
	$KEY_ENTER = KEY_ENTER,
	$KEY_SRESET = KEY_SRESET,
	$KEY_RESET = KEY_RESET,
	$KEY_PRINT = KEY_PRINT,
	$KEY_LL = KEY_LL,
	$KEY_A1 = KEY_A1,
	$KEY_A3 = KEY_A3,
	$KEY_B2 = KEY_B2,
	$KEY_C1 = KEY_C1,
	$KEY_C3 = KEY_C3,
	$KEY_BTAB = KEY_BTAB,
	$KEY_BEG = KEY_BEG,
	$KEY_CANCEL = KEY_CANCEL,
	$KEY_CLOSE = KEY_CLOSE,
	$KEY_COMMAND = KEY_COMMAND,
	$KEY_COPY = KEY_COPY,
	$KEY_CREATE = KEY_CREATE,
	$KEY_END = KEY_END,
	$KEY_EXIT = KEY_EXIT,
	$KEY_FIND = KEY_FIND,
	$KEY_HELP = KEY_HELP,
	$KEY_MARK = KEY_MARK,
	$KEY_MESSAGE = KEY_MESSAGE,
	$KEY_MOVE = KEY_MOVE,
	$KEY_NEXT = KEY_NEXT,
	$KEY_OPEN = KEY_OPEN,
	$KEY_OPTIONS = KEY_OPTIONS,
	$KEY_PREVIOUS = KEY_PREVIOUS,
	$KEY_REDO = KEY_REDO,
	$KEY_REFERENCE = KEY_REFERENCE,
	$KEY_REFRESH = KEY_REFRESH,
	$KEY_REPLACE = KEY_REPLACE,
	$KEY_RESTART = KEY_RESTART,
	$KEY_RESUME = KEY_RESUME,
	$KEY_SAVE = KEY_SAVE,
	$KEY_SBEG = KEY_SBEG,
	$KEY_SCANCEL = KEY_SCANCEL,
	$KEY_SCOMMAND = KEY_SCOMMAND,
	$KEY_SCOPY = KEY_SCOPY,
	$KEY_SCREATE = KEY_SCREATE,
	$KEY_SDC = KEY_SDC,
	$KEY_SDL = KEY_SDL,
	$KEY_SELECT = KEY_SELECT,
	$KEY_SEND = KEY_SEND,
	$KEY_SEOL = KEY_SEOL,
	$KEY_SEXIT = KEY_SEXIT,
	$KEY_SFIND = KEY_SFIND,
	$KEY_SHELP = KEY_SHELP,
	$KEY_SHOME = KEY_SHOME,
	$KEY_SIC = KEY_SIC,
	$KEY_SLEFT = KEY_SLEFT,
	$KEY_SMESSAGE = KEY_SMESSAGE,
	$KEY_SMOVE = KEY_SMOVE,
	$KEY_SNEXT = KEY_SNEXT,
	$KEY_SOPTIONS = KEY_SOPTIONS,
	$KEY_SPREVIOUS = KEY_SPREVIOUS,
	$KEY_SPRINT = KEY_SPRINT,
	$KEY_SREDO = KEY_SREDO,
	$KEY_SREPLACE = KEY_SREPLACE,
	$KEY_SRIGHT = KEY_SRIGHT,
	$KEY_SRSUME = KEY_SRSUME,
	$KEY_SSAVE = KEY_SSAVE,
	$KEY_SSUSPEND = KEY_SSUSPEND,
	$KEY_SUNDO = KEY_SUNDO,
	$KEY_SUSPEND = KEY_SUSPEND,
	$KEY_UNDO = KEY_UNDO
};
