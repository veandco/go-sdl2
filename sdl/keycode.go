package sdl

// #include "sdl_wrapper.h"
import "C"

const K_SCANCODE_MASK = 1 << 30

// The SDL virtual key representation.
// (https://wiki.libsdl.org/SDL_Keycode)
// (https://wiki.libsdl.org/SDLKeycodeLookup)
const (
	K_UNKNOWN = C.SDLK_UNKNOWN // "" (no name, empty string)

	K_RETURN     = C.SDLK_RETURN     // "Return" (the Enter key (main keyboard))
	K_ESCAPE     = C.SDLK_ESCAPE     // "Escape" (the Esc key)
	K_BACKSPACE  = C.SDLK_BACKSPACE  // "Backspace"
	K_TAB        = C.SDLK_TAB        // "Tab" (the Tab key)
	K_SPACE      = C.SDLK_SPACE      // "Space" (the Space Bar key(s))
	K_EXCLAIM    = C.SDLK_EXCLAIM    // "!"
	K_QUOTEDBL   = C.SDLK_QUOTEDBL   // """
	K_HASH       = C.SDLK_HASH       // "#"
	K_PERCENT    = C.SDLK_PERCENT    // "%"
	K_DOLLAR     = C.SDLK_DOLLAR     // "$"
	K_AMPERSAND  = C.SDLK_AMPERSAND  // "&"
	K_QUOTE      = C.SDLK_QUOTE      // "'"
	K_LEFTPAREN  = C.SDLK_LEFTPAREN  // "("
	K_RIGHTPAREN = C.SDLK_RIGHTPAREN // ")"
	K_ASTERISK   = C.SDLK_ASTERISK   // "*"
	K_PLUS       = C.SDLK_PLUS       // "+"
	K_COMMA      = C.SDLK_COMMA      // ","
	K_MINUS      = C.SDLK_MINUS      // "-"
	K_PERIOD     = C.SDLK_PERIOD     // "."
	K_SLASH      = C.SDLK_SLASH      // "/"
	K_0          = C.SDLK_0          // "0"
	K_1          = C.SDLK_1          // "1"
	K_2          = C.SDLK_2          // "2"
	K_3          = C.SDLK_3          // "3"
	K_4          = C.SDLK_4          // "4"
	K_5          = C.SDLK_5          // "5"
	K_6          = C.SDLK_6          // "6"
	K_7          = C.SDLK_7          // "7"
	K_8          = C.SDLK_8          // "8"
	K_9          = C.SDLK_9          // "9"
	K_COLON      = C.SDLK_COLON      // ":"
	K_SEMICOLON  = C.SDLK_SEMICOLON  // ";"
	K_LESS       = C.SDLK_LESS       // "<"
	K_EQUALS     = C.SDLK_EQUALS     // "="
	K_GREATER    = C.SDLK_GREATER    // ">"
	K_QUESTION   = C.SDLK_QUESTION   // "?"
	K_AT         = C.SDLK_AT         // "@"
	/*
	   Skip uppercase letters
	*/
	K_LEFTBRACKET  = C.SDLK_LEFTBRACKET  // "["
	K_BACKSLASH    = C.SDLK_BACKSLASH    // "\"
	K_RIGHTBRACKET = C.SDLK_RIGHTBRACKET // "]"
	K_CARET        = C.SDLK_CARET        // "^"
	K_UNDERSCORE   = C.SDLK_UNDERSCORE   // "_"
	K_BACKQUOTE    = C.SDLK_BACKQUOTE    // "`"
	K_a            = C.SDLK_a            // "A"
	K_b            = C.SDLK_b            // "B"
	K_c            = C.SDLK_c            // "C"
	K_d            = C.SDLK_d            // "D"
	K_e            = C.SDLK_e            // "E"
	K_f            = C.SDLK_f            // "F"
	K_g            = C.SDLK_g            // "G"
	K_h            = C.SDLK_h            // "H"
	K_i            = C.SDLK_i            // "I"
	K_j            = C.SDLK_j            // "J"
	K_k            = C.SDLK_k            // "K"
	K_l            = C.SDLK_l            // "L"
	K_m            = C.SDLK_m            // "M"
	K_n            = C.SDLK_n            // "N"
	K_o            = C.SDLK_o            // "O"
	K_p            = C.SDLK_p            // "P"
	K_q            = C.SDLK_q            // "Q"
	K_r            = C.SDLK_r            // "R"
	K_s            = C.SDLK_s            // "S"
	K_t            = C.SDLK_t            // "T"
	K_u            = C.SDLK_u            // "U"
	K_v            = C.SDLK_v            // "V"
	K_w            = C.SDLK_w            // "W"
	K_x            = C.SDLK_x            // "X"
	K_y            = C.SDLK_y            // "Y"
	K_z            = C.SDLK_z            // "Z"

	K_CAPSLOCK = C.SDLK_CAPSLOCK // "CapsLock"

	K_F1  = C.SDLK_F1  // "F1"
	K_F2  = C.SDLK_F2  // "F2"
	K_F3  = C.SDLK_F3  // "F3"
	K_F4  = C.SDLK_F4  // "F4"
	K_F5  = C.SDLK_F5  // "F5"
	K_F6  = C.SDLK_F6  // "F6"
	K_F7  = C.SDLK_F7  // "F7"
	K_F8  = C.SDLK_F8  // "F8"
	K_F9  = C.SDLK_F9  // "F9"
	K_F10 = C.SDLK_F10 // "F10"
	K_F11 = C.SDLK_F11 // "F11"
	K_F12 = C.SDLK_F12 // "F12"

	K_PRINTSCREEN = C.SDLK_PRINTSCREEN // "PrintScreen"
	K_SCROLLLOCK  = C.SDLK_SCROLLLOCK  // "ScrollLock"
	K_PAUSE       = C.SDLK_PAUSE       // "Pause" (the Pause / Break key)
	K_INSERT      = C.SDLK_INSERT      // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	K_HOME        = C.SDLK_HOME        // "Home"
	K_PAGEUP      = C.SDLK_PAGEUP      // "PageUp"
	K_DELETE      = C.SDLK_DELETE      // "Delete"
	K_END         = C.SDLK_END         // "End"
	K_PAGEDOWN    = C.SDLK_PAGEDOWN    // "PageDown"
	K_RIGHT       = C.SDLK_RIGHT       // "Right" (the Right arrow key (navigation keypad))
	K_LEFT        = C.SDLK_LEFT        // "Left" (the Left arrow key (navigation keypad))
	K_DOWN        = C.SDLK_DOWN        // "Down" (the Down arrow key (navigation keypad))
	K_UP          = C.SDLK_UP          // "Up" (the Up arrow key (navigation keypad))

	K_NUMLOCKCLEAR = C.SDLK_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	K_KP_DIVIDE    = C.SDLK_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	K_KP_MULTIPLY  = C.SDLK_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	K_KP_MINUS     = C.SDLK_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	K_KP_PLUS      = C.SDLK_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	K_KP_ENTER     = C.SDLK_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	K_KP_1         = C.SDLK_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	K_KP_2         = C.SDLK_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	K_KP_3         = C.SDLK_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	K_KP_4         = C.SDLK_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	K_KP_5         = C.SDLK_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	K_KP_6         = C.SDLK_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	K_KP_7         = C.SDLK_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	K_KP_8         = C.SDLK_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	K_KP_9         = C.SDLK_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	K_KP_0         = C.SDLK_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	K_KP_PERIOD    = C.SDLK_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	K_APPLICATION    = C.SDLK_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	K_POWER          = C.SDLK_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	K_KP_EQUALS      = C.SDLK_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	K_F13            = C.SDLK_F13            // "F13"
	K_F14            = C.SDLK_F14            // "F14"
	K_F15            = C.SDLK_F15            // "F15"
	K_F16            = C.SDLK_F16            // "F16"
	K_F17            = C.SDLK_F17            // "F17"
	K_F18            = C.SDLK_F18            // "F18"
	K_F19            = C.SDLK_F19            // "F19"
	K_F20            = C.SDLK_F20            // "F20"
	K_F21            = C.SDLK_F21            // "F21"
	K_F22            = C.SDLK_F22            // "F22"
	K_F23            = C.SDLK_F23            // "F23"
	K_F24            = C.SDLK_F24            // "F24"
	K_EXECUTE        = C.SDLK_EXECUTE        // "Execute"
	K_HELP           = C.SDLK_HELP           // "Help"
	K_MENU           = C.SDLK_MENU           // "Menu"
	K_SELECT         = C.SDLK_SELECT         // "Select"
	K_STOP           = C.SDLK_STOP           // "Stop"
	K_AGAIN          = C.SDLK_AGAIN          // "Again" (the Again key (Redo))
	K_UNDO           = C.SDLK_UNDO           // "Undo"
	K_CUT            = C.SDLK_CUT            // "Cut"
	K_COPY           = C.SDLK_COPY           // "Copy"
	K_PASTE          = C.SDLK_PASTE          // "Paste"
	K_FIND           = C.SDLK_FIND           // "Find"
	K_MUTE           = C.SDLK_MUTE           // "Mute"
	K_VOLUMEUP       = C.SDLK_VOLUMEUP       // "VolumeUp"
	K_VOLUMEDOWN     = C.SDLK_VOLUMEDOWN     // "VolumeDown"
	K_KP_COMMA       = C.SDLK_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	K_KP_EQUALSAS400 = C.SDLK_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	K_ALTERASE   = C.SDLK_ALTERASE   // "AltErase" (Erase-Eaze)
	K_SYSREQ     = C.SDLK_SYSREQ     // "SysReq" (the SysReq key)
	K_CANCEL     = C.SDLK_CANCEL     // "Cancel"
	K_CLEAR      = C.SDLK_CLEAR      // "Clear"
	K_PRIOR      = C.SDLK_PRIOR      // "Prior"
	K_RETURN2    = C.SDLK_RETURN2    // "Return"
	K_SEPARATOR  = C.SDLK_SEPARATOR  // "Separator"
	K_OUT        = C.SDLK_OUT        // "Out"
	K_OPER       = C.SDLK_OPER       // "Oper"
	K_CLEARAGAIN = C.SDLK_CLEARAGAIN // "Clear / Again"
	K_CRSEL      = C.SDLK_CRSEL      // "CrSel"
	K_EXSEL      = C.SDLK_EXSEL      // "ExSel"

	K_KP_00              = C.SDLK_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	K_KP_000             = C.SDLK_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	K_THOUSANDSSEPARATOR = C.SDLK_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	K_DECIMALSEPARATOR   = C.SDLK_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	K_CURRENCYUNIT       = C.SDLK_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	K_CURRENCYSUBUNIT    = C.SDLK_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	K_KP_LEFTPAREN       = C.SDLK_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	K_KP_RIGHTPAREN      = C.SDLK_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	K_KP_LEFTBRACE       = C.SDLK_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	K_KP_RIGHTBRACE      = C.SDLK_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	K_KP_TAB             = C.SDLK_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	K_KP_BACKSPACE       = C.SDLK_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	K_KP_A               = C.SDLK_KP_A               // "Keypad A" (the A key (numeric keypad))
	K_KP_B               = C.SDLK_KP_B               // "Keypad B" (the B key (numeric keypad))
	K_KP_C               = C.SDLK_KP_C               // "Keypad C" (the C key (numeric keypad))
	K_KP_D               = C.SDLK_KP_D               // "Keypad D" (the D key (numeric keypad))
	K_KP_E               = C.SDLK_KP_E               // "Keypad E" (the E key (numeric keypad))
	K_KP_F               = C.SDLK_KP_F               // "Keypad F" (the F key (numeric keypad))
	K_KP_XOR             = C.SDLK_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	K_KP_POWER           = C.SDLK_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	K_KP_PERCENT         = C.SDLK_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	K_KP_LESS            = C.SDLK_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	K_KP_GREATER         = C.SDLK_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	K_KP_AMPERSAND       = C.SDLK_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	K_KP_DBLAMPERSAND    = C.SDLK_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	K_KP_VERTICALBAR     = C.SDLK_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	K_KP_DBLVERTICALBAR  = C.SDLK_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	K_KP_COLON           = C.SDLK_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	K_KP_HASH            = C.SDLK_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	K_KP_SPACE           = C.SDLK_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	K_KP_AT              = C.SDLK_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	K_KP_EXCLAM          = C.SDLK_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	K_KP_MEMSTORE        = C.SDLK_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	K_KP_MEMRECALL       = C.SDLK_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	K_KP_MEMCLEAR        = C.SDLK_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	K_KP_MEMADD          = C.SDLK_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	K_KP_MEMSUBTRACT     = C.SDLK_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	K_KP_MEMMULTIPLY     = C.SDLK_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	K_KP_MEMDIVIDE       = C.SDLK_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	K_KP_PLUSMINUS       = C.SDLK_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	K_KP_CLEAR           = C.SDLK_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	K_KP_CLEARENTRY      = C.SDLK_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	K_KP_BINARY          = C.SDLK_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	K_KP_OCTAL           = C.SDLK_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	K_KP_DECIMAL         = C.SDLK_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	K_KP_HEXADECIMAL     = C.SDLK_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	K_LCTRL  = C.SDLK_LCTRL  // "Left Ctrl"
	K_LSHIFT = C.SDLK_LSHIFT // "Left Shift"
	K_LALT   = C.SDLK_LALT   // "Left Alt" (alt, option)
	K_LGUI   = C.SDLK_LGUI   // "Left GUI" (windows, command (apple), meta)
	K_RCTRL  = C.SDLK_RCTRL  // "Right Ctrl"
	K_RSHIFT = C.SDLK_RSHIFT // "Right Shift"
	K_RALT   = C.SDLK_RALT   // "Right Alt" (alt, option)
	K_RGUI   = C.SDLK_RGUI   // "Right GUI" (windows, command (apple), meta)

	K_MODE = C.SDLK_MODE // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)

	K_AUDIONEXT    = C.SDLK_AUDIONEXT    // "AudioNext" (the Next Track media key)
	K_AUDIOPREV    = C.SDLK_AUDIOPREV    // "AudioPrev" (the Previous Track media key)
	K_AUDIOSTOP    = C.SDLK_AUDIOSTOP    // "AudioStop" (the Stop media key)
	K_AUDIOPLAY    = C.SDLK_AUDIOPLAY    // "AudioPlay" (the Play media key)
	K_AUDIOMUTE    = C.SDLK_AUDIOMUTE    // "AudioMute" (the Mute volume key)
	K_MEDIASELECT  = C.SDLK_MEDIASELECT  // "MediaSelect" (the Media Select key)
	K_WWW          = C.SDLK_WWW          // "WWW" (the WWW/World Wide Web key)
	K_MAIL         = C.SDLK_MAIL         // "Mail" (the Mail/eMail key)
	K_CALCULATOR   = C.SDLK_CALCULATOR   // "Calculator" (the Calculator key)
	K_COMPUTER     = C.SDLK_COMPUTER     // "Computer" (the My Computer key)
	K_AC_SEARCH    = C.SDLK_AC_SEARCH    // "AC Search" (the Search key (application control keypad))
	K_AC_HOME      = C.SDLK_AC_HOME      // "AC Home" (the Home key (application control keypad))
	K_AC_BACK      = C.SDLK_AC_BACK      // "AC Back" (the Back key (application control keypad))
	K_AC_FORWARD   = C.SDLK_AC_FORWARD   // "AC Forward" (the Forward key (application control keypad))
	K_AC_STOP      = C.SDLK_AC_STOP      // "AC Stop" (the Stop key (application control keypad))
	K_AC_REFRESH   = C.SDLK_AC_REFRESH   // "AC Refresh" (the Refresh key (application control keypad))
	K_AC_BOOKMARKS = C.SDLK_AC_BOOKMARKS // "AC Bookmarks" (the Bookmarks key (application control keypad))

	K_BRIGHTNESSDOWN = C.SDLK_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	K_BRIGHTNESSUP   = C.SDLK_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	K_DISPLAYSWITCH  = C.SDLK_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	K_KBDILLUMTOGGLE = C.SDLK_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	K_KBDILLUMDOWN   = C.SDLK_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	K_KBDILLUMUP     = C.SDLK_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	K_EJECT          = C.SDLK_EJECT          // "Eject" (the Eject key)
	K_SLEEP          = C.SDLK_SLEEP          // "Sleep" (the Sleep key)
)

// An enumeration of key modifier masks.
// (https://wiki.libsdl.org/SDL_Keymod)
const (
	KMOD_NONE     = C.KMOD_NONE     // 0 (no modifier is applicable)
	KMOD_LSHIFT   = C.KMOD_LSHIFT   // the left Shift key is down
	KMOD_RSHIFT   = C.KMOD_RSHIFT   // the right Shift key is down
	KMOD_LCTRL    = C.KMOD_LCTRL    // the left Ctrl (Control) key is down
	KMOD_RCTRL    = C.KMOD_RCTRL    // the right Ctrl (Control) key is down
	KMOD_LALT     = C.KMOD_LALT     // the left Alt key is down
	KMOD_RALT     = C.KMOD_RALT     // the right Alt key is down
	KMOD_LGUI     = C.KMOD_LGUI     // the left GUI key (often the Windows key) is down
	KMOD_RGUI     = C.KMOD_RGUI     // the right GUI key (often the Windows key) is down
	KMOD_NUM      = C.KMOD_NUM      // the Num Lock key (may be located on an extended keypad) is down
	KMOD_CAPS     = C.KMOD_CAPS     // the Caps Lock key is down
	KMOD_MODE     = C.KMOD_MODE     // the AltGr key is down
	KMOD_CTRL     = C.KMOD_CTRL     // (KMOD_LCTRL|KMOD_RCTRL)
	KMOD_SHIFT    = C.KMOD_SHIFT    // (KMOD_LSHIFT|KMOD_RSHIFT)
	KMOD_ALT      = C.KMOD_ALT      // (KMOD_LALT|KMOD_RALT)
	KMOD_GUI      = C.KMOD_GUI      // (KMOD_LGUI|KMOD_RGUI)
	KMOD_RESERVED = C.KMOD_RESERVED // reserved for future use
)

// Keycode is the SDL virtual key representation.
// (https://wiki.libsdl.org/SDL_Keycode)
type Keycode C.SDL_Keycode

// Keymod is a key modifier masks.
// (https://wiki.libsdl.org/SDL_Keymod)
type Keymod C.SDL_Keymod

func (code Keycode) c() C.SDL_Keycode {
	return C.SDL_Keycode(code)
}

func (mod Keymod) c() C.SDL_Keymod {
	return C.SDL_Keymod(mod)
}
