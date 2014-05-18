package sdl

// #include <SDL2/SDL.h>
import "C"

const (
	SCANCODE_UNKNOWN = 0

	/**
	 *  \name Usage page 0x07
	 *
	 *  These values are from usage page 0x07 (USB keyboard page).
	 */
	/*@{*/

	SCANCODE_A = 4
	SCANCODE_B = 5
	SCANCODE_C = 6
	SCANCODE_D = 7
	SCANCODE_E = 8
	SCANCODE_F = 9
	SCANCODE_G = 10
	SCANCODE_H = 11
	SCANCODE_I = 12
	SCANCODE_J = 13
	SCANCODE_K = 14
	SCANCODE_L = 15
	SCANCODE_M = 16
	SCANCODE_N = 17
	SCANCODE_O = 18
	SCANCODE_P = 19
	SCANCODE_Q = 20
	SCANCODE_R = 21
	SCANCODE_S = 22
	SCANCODE_T = 23
	SCANCODE_U = 24
	SCANCODE_V = 25
	SCANCODE_W = 26
	SCANCODE_X = 27
	SCANCODE_Y = 28
	SCANCODE_Z = 29

	SCANCODE_1 = 30
	SCANCODE_2 = 31
	SCANCODE_3 = 32
	SCANCODE_4 = 33
	SCANCODE_5 = 34
	SCANCODE_6 = 35
	SCANCODE_7 = 36
	SCANCODE_8 = 37
	SCANCODE_9 = 38
	SCANCODE_0 = 39

	SCANCODE_RETURN    = 40
	SCANCODE_ESCAPE    = 41
	SCANCODE_BACKSPACE = 42
	SCANCODE_TAB       = 43
	SCANCODE_SPACE     = 44

	SCANCODE_MINUS        = 45
	SCANCODE_EQUALS       = 46
	SCANCODE_LEFTBRACKET  = 47
	SCANCODE_RIGHTBRACKET = 48
	SCANCODE_BACKSLASH    = 49 /**< Located at the lower left of the return
	 *   key on ISO keyboards and at the right end
	 *   of the QWERTY row on ANSI keyboards.
	 *   Produces REVERSE SOLIDUS (backslash) and
	 *   VERTICAL LINE in a US layout REVERSE
	 *   SOLIDUS and VERTICAL LINE in a UK Mac
	 *   layout NUMBER SIGN and TILDE in a UK
	 *   Windows layout DOLLAR SIGN and POUND SIGN
	 *   in a Swiss German layout NUMBER SIGN and
	 *   APOSTROPHE in a German layout GRAVE
	 *   ACCENT and POUND SIGN in a French Mac
	 *   layout and ASTERISK and MICRO SIGN in a
	 *   French Windows layout.
	 */
	SDL_SCANCODE_NONUSHASH = 50 /**< ISO USB keyboards actually use this code
	 *   instead of 49 for the same key but all
	 *   OSes I've seen treat the two codes
	 *   identically. So as an implementor unless
	 *   your keyboard generates both of those
	 *   codes and your OS treats them differently
	 *   you should generate SDL_SCANCODE_BACKSLASH
	 *   instead of this code. As a user you
	 *   should not rely on this code because SDL
	 *   will never generate it with most (all?)
	 *   keyboards.
	 */
	SDL_SCANCODE_SEMICOLON  = 51
	SDL_SCANCODE_APOSTROPHE = 52
	SDL_SCANCODE_GRAVE      = 53 /**< Located in the top left corner (on both ANSI
	 *   and ISO keyboards). Produces GRAVE ACCENT and
	 *   TILDE in a US Windows layout and in US and UK
	 *   Mac layouts on ANSI keyboards GRAVE ACCENT
	 *   and NOT SIGN in a UK Windows layout SECTION
	 *   SIGN and PLUS-MINUS SIGN in US and UK Mac
	 *   layouts on ISO keyboards SECTION SIGN and
	 *   DEGREE SIGN in a Swiss German layout (Mac:
	 *   only on ISO keyboards) CIRCUMFLEX ACCENT and
	 *   DEGREE SIGN in a German layout (Mac: only on
	 *   ISO keyboards) SUPERSCRIPT TWO and TILDE in a
	 *   French Windows layout COMMERCIAL AT and
	 *   NUMBER SIGN in a French Mac layout on ISO
	 *   keyboards and LESS-THAN SIGN and GREATER-THAN
	 *   SIGN in a Swiss German German or French Mac
	 *   layout on ANSI keyboards.
	 */
	SCANCODE_COMMA  = 54
	SCANCODE_PERIOD = 55
	SCANCODE_SLASH  = 56

	SCANCODE_CAPSLOCK = 57

	SCANCODE_F1  = 58
	SCANCODE_F2  = 59
	SCANCODE_F3  = 60
	SCANCODE_F4  = 61
	SCANCODE_F5  = 62
	SCANCODE_F6  = 63
	SCANCODE_F7  = 64
	SCANCODE_F8  = 65
	SCANCODE_F9  = 66
	SCANCODE_F10 = 67
	SCANCODE_F11 = 68
	SCANCODE_F12 = 69

	SCANCODE_PRINTSCREEN = 70
	SCANCODE_SCROLLLOCK  = 71
	SCANCODE_PAUSE       = 72
	SCANCODE_INSERT      = 73 /**< insert on PC help on some Mac keyboards (but
	  does send code 73 not 117) */
	SCANCODE_HOME     = 74
	SCANCODE_PAGEUP   = 75
	SCANCODE_DELETE   = 76
	SCANCODE_END      = 77
	SCANCODE_PAGEDOWN = 78
	SCANCODE_RIGHT    = 79
	SCANCODE_LEFT     = 80
	SCANCODE_DOWN     = 81
	SCANCODE_UP       = 82

	SCANCODE_NUMLOCKCLEAR = 83 /**< num lock on PC clear on Mac keyboards
	 */
	SCANCODE_KP_DIVIDE   = 84
	SCANCODE_KP_MULTIPLY = 85
	SCANCODE_KP_MINUS    = 86
	SCANCODE_KP_PLUS     = 87
	SCANCODE_KP_ENTER    = 88
	SCANCODE_KP_1        = 89
	SCANCODE_KP_2        = 90
	SCANCODE_KP_3        = 91
	SCANCODE_KP_4        = 92
	SCANCODE_KP_5        = 93
	SCANCODE_KP_6        = 94
	SCANCODE_KP_7        = 95
	SCANCODE_KP_8        = 96
	SCANCODE_KP_9        = 97
	SCANCODE_KP_0        = 98
	SCANCODE_KP_PERIOD   = 99

	SCANCODE_NONUSBACKSLASH = 100 /**< This is the additional key that ISO
	*   keyboards have over ANSI ones
	*   located between left shift and Y.
	*   Produces GRAVE ACCENT and TILDE in a
	*   US or UK Mac layout REVERSE SOLIDUS
	*   (backslash) and VERTICAL LINE in a
	*   US or UK Windows layout and
	*   LESS-THAN SIGN and GREATER-THAN SIGN
	*   in a Swiss German German or French
	*   layout. */
	SCANCODE_APPLICATION = 101 /**< windows contextual menu compose */
	SCANCODE_POWER       = 102 /**< The USB document says this is a status flag
	 *   not a physical key - but some Mac keyboards
	 *   do have a power key. */
	SCANCODE_KP_EQUALS  = 103
	SCANCODE_F13        = 104
	SCANCODE_F14        = 105
	SCANCODE_F15        = 106
	SCANCODE_F16        = 107
	SCANCODE_F17        = 108
	SCANCODE_F18        = 109
	SCANCODE_F19        = 110
	SCANCODE_F20        = 111
	SCANCODE_F21        = 112
	SCANCODE_F22        = 113
	SCANCODE_F23        = 114
	SCANCODE_F24        = 115
	SCANCODE_EXECUTE    = 116
	SCANCODE_HELP       = 117
	SCANCODE_MENU       = 118
	SCANCODE_SELECT     = 119
	SCANCODE_STOP       = 120
	SCANCODE_AGAIN      = 121 /**< redo */
	SCANCODE_UNDO       = 122
	SCANCODE_CUT        = 123
	SCANCODE_COPY       = 124
	SCANCODE_PASTE      = 125
	SCANCODE_FIND       = 126
	SCANCODE_MUTE       = 127
	SCANCODE_VOLUMEUP   = 128
	SCANCODE_VOLUMEDOWN = 129
	/* not sure whether there's a reason to enable these */
	/*     SDL_SCANCODE_LOCKINGCAPSLOCK = 130  */
	/*     SDL_SCANCODE_LOCKINGNUMLOCK = 131 */
	/*     SDL_SCANCODE_LOCKINGSCROLLLOCK = 132 */
	SCANCODE_KP_COMMA       = 133
	SCANCODE_KP_EQUALSAS400 = 134

	SCANCODE_INTERNATIONAL1 = 135 /**< used on Asian keyboards see
	  footnotes in USB doc */
	SCANCODE_INTERNATIONAL2 = 136
	SCANCODE_INTERNATIONAL3 = 137 /**< Yen */
	SCANCODE_INTERNATIONAL4 = 138
	SCANCODE_INTERNATIONAL5 = 139
	SCANCODE_INTERNATIONAL6 = 140
	SCANCODE_INTERNATIONAL7 = 141
	SCANCODE_INTERNATIONAL8 = 142
	SCANCODE_INTERNATIONAL9 = 143
	SCANCODE_LANG1          = 144 /**< Hangul/English toggle */
	SCANCODE_LANG2          = 145 /**< Hanja conversion */
	SCANCODE_LANG3          = 146 /**< Katakana */
	SCANCODE_LANG4          = 147 /**< Hiragana */
	SCANCODE_LANG5          = 148 /**< Zenkaku/Hankaku */
	SCANCODE_LANG6          = 149 /**< reserved */
	SCANCODE_LANG7          = 150 /**< reserved */
	SCANCODE_LANG8          = 151 /**< reserved */
	SCANCODE_LANG9          = 152 /**< reserved */

	SCANCODE_ALTERASE   = 153 /**< Erase-Eaze */
	SCANCODE_SYSREQ     = 154
	SCANCODE_CANCEL     = 155
	SCANCODE_CLEAR      = 156
	SCANCODE_PRIOR      = 157
	SCANCODE_RETURN2    = 158
	SCANCODE_SEPARATOR  = 159
	SCANCODE_OUT        = 160
	SCANCODE_OPER       = 161
	SCANCODE_CLEARAGAIN = 162
	SCANCODE_CRSEL      = 163
	SCANCODE_EXSEL      = 164

	SCANCODE_KP_00              = 176
	SCANCODE_KP_000             = 177
	SCANCODE_THOUSANDSSEPARATOR = 178
	SCANCODE_DECIMALSEPARATOR   = 179
	SCANCODE_CURRENCYUNIT       = 180
	SCANCODE_CURRENCYSUBUNIT    = 181
	SCANCODE_KP_LEFTPAREN       = 182
	SCANCODE_KP_RIGHTPAREN      = 183
	SCANCODE_KP_LEFTBRACE       = 184
	SCANCODE_KP_RIGHTBRACE      = 185
	SCANCODE_KP_TAB             = 186
	SCANCODE_KP_BACKSPACE       = 187
	SCANCODE_KP_A               = 188
	SCANCODE_KP_B               = 189
	SCANCODE_KP_C               = 190
	SCANCODE_KP_D               = 191
	SCANCODE_KP_E               = 192
	SCANCODE_KP_F               = 193
	SCANCODE_KP_XOR             = 194
	SCANCODE_KP_POWER           = 195
	SCANCODE_KP_PERCENT         = 196
	SCANCODE_KP_LESS            = 197
	SCANCODE_KP_GREATER         = 198
	SCANCODE_KP_AMPERSAND       = 199
	SCANCODE_KP_DBLAMPERSAND    = 200
	SCANCODE_KP_VERTICALBAR     = 201
	SCANCODE_KP_DBLVERTICALBAR  = 202
	SCANCODE_KP_COLON           = 203
	SCANCODE_KP_HASH            = 204
	SCANCODE_KP_SPACE           = 205
	SCANCODE_KP_AT              = 206
	SCANCODE_KP_EXCLAM          = 207
	SCANCODE_KP_MEMSTORE        = 208
	SCANCODE_KP_MEMRECALL       = 209
	SCANCODE_KP_MEMCLEAR        = 210
	SCANCODE_KP_MEMADD          = 211
	SCANCODE_KP_MEMSUBTRACT     = 212
	SCANCODE_KP_MEMMULTIPLY     = 213
	SCANCODE_KP_MEMDIVIDE       = 214
	SCANCODE_KP_PLUSMINUS       = 215
	SCANCODE_KP_CLEAR           = 216
	SCANCODE_KP_CLEARENTRY      = 217
	SCANCODE_KP_BINARY          = 218
	SCANCODE_KP_OCTAL           = 219
	SCANCODE_KP_DECIMAL         = 220
	SCANCODE_KP_HEXADECIMAL     = 221

	SCANCODE_LCTRL  = 224
	SCANCODE_LSHIFT = 225
	SCANCODE_LALT   = 226 /**< alt option */
	SCANCODE_LGUI   = 227 /**< windows command (apple) meta */
	SCANCODE_RCTRL  = 228
	SCANCODE_RSHIFT = 229
	SCANCODE_RALT   = 230 /**< alt gr option */
	SCANCODE_RGUI   = 231 /**< windows command (apple) meta */

	SCANCODE_MODE = 257 /**< I'm not sure if this is really not covered
	 *   by any of the above but since there's a
	 *   special KMOD_MODE for it I'm adding it here
	 */

	/*@}*/ /*Usage page 0x07*/

	/**
	 *  \name Usage page 0x0C
	 *
	 *  These values are mapped from usage page 0x0C (USB consumer page).
	 */
	/*@{*/

	SCANCODE_AUDIONEXT    = 258
	SCANCODE_AUDIOPREV    = 259
	SCANCODE_AUDIOSTOP    = 260
	SCANCODE_AUDIOPLAY    = 261
	SCANCODE_AUDIOMUTE    = 262
	SCANCODE_MEDIASELECT  = 263
	SCANCODE_WWW          = 264
	SCANCODE_MAIL         = 265
	SCANCODE_CALCULATOR   = 266
	SCANCODE_COMPUTER     = 267
	SCANCODE_AC_SEARCH    = 268
	SCANCODE_AC_HOME      = 269
	SCANCODE_AC_BACK      = 270
	SCANCODE_AC_FORWARD   = 271
	SCANCODE_AC_STOP      = 272
	SCANCODE_AC_REFRESH   = 273
	SCANCODE_AC_BOOKMARKS = 274

	/*@}*/ /*Usage page 0x0C*/

	/**
	 *  \name Walther keys
	 *
	 *  These are values that Christian Walther added (for mac keyboard?).
	 */
	/*@{*/

	SCANCODE_BRIGHTNESSDOWN = 275
	SCANCODE_BRIGHTNESSUP   = 276
	SCANCODE_DISPLAYSWITCH  = 277 /**< display mirroring/dual display
	  switch video mode switch */
	SCANCODE_KBDILLUMTOGGLE = 278
	SCANCODE_KBDILLUMDOWN   = 279
	SCANCODE_KBDILLUMUP     = 280
	SCANCODE_EJECT          = 281
	SCANCODE_SLEEP          = 282

	SCANCODE_APP1 = 283
	SCANCODE_APP2 = 284

	/*@}*/ /*Walther keys*/

	/* Add any other keys here. */

	NUM_SCANCODES = 512 /**< not a key just marks the number of scancodes
	  for array bounds */
)

type Scancode uint32

func (code Scancode) c() C.SDL_Scancode {
    return C.SDL_Scancode(code)
}
