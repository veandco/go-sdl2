package sdl

// #include "sdl_wrapper.h"
import "C"

const K_SCANCODE_MASK = 1 << 30

// The SDL virtual key representation.
// (https://wiki.libsdl.org/SDL_Keycode)
// (https://wiki.libsdl.org/SDLKeycodeLookup)
type Keycode C.SDL_Keycode

const (
	K_UNKNOWN Keycode = C.SDLK_UNKNOWN // "" (no name, empty string)

	K_RETURN     Keycode = C.SDLK_RETURN     // "Return" (the Enter key (main keyboard))
	K_ESCAPE     Keycode = C.SDLK_ESCAPE     // "Escape" (the Esc key)
	K_BACKSPACE  Keycode = C.SDLK_BACKSPACE  // "Backspace"
	K_TAB        Keycode = C.SDLK_TAB        // "Tab" (the Tab key)
	K_SPACE      Keycode = C.SDLK_SPACE      // "Space" (the Space Bar key(s))
	K_EXCLAIM    Keycode = C.SDLK_EXCLAIM    // "!"
	K_QUOTEDBL   Keycode = C.SDLK_QUOTEDBL   // """
	K_HASH       Keycode = C.SDLK_HASH       // "#"
	K_PERCENT    Keycode = C.SDLK_PERCENT    // "%"
	K_DOLLAR     Keycode = C.SDLK_DOLLAR     // "$"
	K_AMPERSAND  Keycode = C.SDLK_AMPERSAND  // "&"
	K_QUOTE      Keycode = C.SDLK_QUOTE      // "'"
	K_LEFTPAREN  Keycode = C.SDLK_LEFTPAREN  // "("
	K_RIGHTPAREN Keycode = C.SDLK_RIGHTPAREN // ")"
	K_ASTERISK   Keycode = C.SDLK_ASTERISK   // "*"
	K_PLUS       Keycode = C.SDLK_PLUS       // "+"
	K_COMMA      Keycode = C.SDLK_COMMA      // ","
	K_MINUS      Keycode = C.SDLK_MINUS      // "-"
	K_PERIOD     Keycode = C.SDLK_PERIOD     // "."
	K_SLASH      Keycode = C.SDLK_SLASH      // "/"
	K_0          Keycode = C.SDLK_0          // "0"
	K_1          Keycode = C.SDLK_1          // "1"
	K_2          Keycode = C.SDLK_2          // "2"
	K_3          Keycode = C.SDLK_3          // "3"
	K_4          Keycode = C.SDLK_4          // "4"
	K_5          Keycode = C.SDLK_5          // "5"
	K_6          Keycode = C.SDLK_6          // "6"
	K_7          Keycode = C.SDLK_7          // "7"
	K_8          Keycode = C.SDLK_8          // "8"
	K_9          Keycode = C.SDLK_9          // "9"
	K_COLON      Keycode = C.SDLK_COLON      // ":"
	K_SEMICOLON  Keycode = C.SDLK_SEMICOLON  // ";"
	K_LESS       Keycode = C.SDLK_LESS       // "<"
	K_EQUALS     Keycode = C.SDLK_EQUALS     // "="
	K_GREATER    Keycode = C.SDLK_GREATER    // ">"
	K_QUESTION   Keycode = C.SDLK_QUESTION   // "?"
	K_AT         Keycode = C.SDLK_AT         // "@"
	/*
	   Skip uppercase letters
	*/
	K_LEFTBRACKET  Keycode = C.SDLK_LEFTBRACKET  // "["
	K_BACKSLASH    Keycode = C.SDLK_BACKSLASH    // "\"
	K_RIGHTBRACKET Keycode = C.SDLK_RIGHTBRACKET // "]"
	K_CARET        Keycode = C.SDLK_CARET        // "^"
	K_UNDERSCORE   Keycode = C.SDLK_UNDERSCORE   // "_"
	K_BACKQUOTE    Keycode = C.SDLK_BACKQUOTE    // "`"
	K_a            Keycode = C.SDLK_a            // "A"
	K_b            Keycode = C.SDLK_b            // "B"
	K_c            Keycode = C.SDLK_c            // "C"
	K_d            Keycode = C.SDLK_d            // "D"
	K_e            Keycode = C.SDLK_e            // "E"
	K_f            Keycode = C.SDLK_f            // "F"
	K_g            Keycode = C.SDLK_g            // "G"
	K_h            Keycode = C.SDLK_h            // "H"
	K_i            Keycode = C.SDLK_i            // "I"
	K_j            Keycode = C.SDLK_j            // "J"
	K_k            Keycode = C.SDLK_k            // "K"
	K_l            Keycode = C.SDLK_l            // "L"
	K_m            Keycode = C.SDLK_m            // "M"
	K_n            Keycode = C.SDLK_n            // "N"
	K_o            Keycode = C.SDLK_o            // "O"
	K_p            Keycode = C.SDLK_p            // "P"
	K_q            Keycode = C.SDLK_q            // "Q"
	K_r            Keycode = C.SDLK_r            // "R"
	K_s            Keycode = C.SDLK_s            // "S"
	K_t            Keycode = C.SDLK_t            // "T"
	K_u            Keycode = C.SDLK_u            // "U"
	K_v            Keycode = C.SDLK_v            // "V"
	K_w            Keycode = C.SDLK_w            // "W"
	K_x            Keycode = C.SDLK_x            // "X"
	K_y            Keycode = C.SDLK_y            // "Y"
	K_z            Keycode = C.SDLK_z            // "Z"

	K_CAPSLOCK Keycode = C.SDLK_CAPSLOCK // "CapsLock"

	K_F1  Keycode = C.SDLK_F1  // "F1"
	K_F2  Keycode = C.SDLK_F2  // "F2"
	K_F3  Keycode = C.SDLK_F3  // "F3"
	K_F4  Keycode = C.SDLK_F4  // "F4"
	K_F5  Keycode = C.SDLK_F5  // "F5"
	K_F6  Keycode = C.SDLK_F6  // "F6"
	K_F7  Keycode = C.SDLK_F7  // "F7"
	K_F8  Keycode = C.SDLK_F8  // "F8"
	K_F9  Keycode = C.SDLK_F9  // "F9"
	K_F10 Keycode = C.SDLK_F10 // "F10"
	K_F11 Keycode = C.SDLK_F11 // "F11"
	K_F12 Keycode = C.SDLK_F12 // "F12"

	K_PRINTSCREEN Keycode = C.SDLK_PRINTSCREEN // "PrintScreen"
	K_SCROLLLOCK  Keycode = C.SDLK_SCROLLLOCK  // "ScrollLock"
	K_PAUSE       Keycode = C.SDLK_PAUSE       // "Pause" (the Pause / Break key)
	K_INSERT      Keycode = C.SDLK_INSERT      // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	K_HOME        Keycode = C.SDLK_HOME        // "Home"
	K_PAGEUP      Keycode = C.SDLK_PAGEUP      // "PageUp"
	K_DELETE      Keycode = C.SDLK_DELETE      // "Delete"
	K_END         Keycode = C.SDLK_END         // "End"
	K_PAGEDOWN    Keycode = C.SDLK_PAGEDOWN    // "PageDown"
	K_RIGHT       Keycode = C.SDLK_RIGHT       // "Right" (the Right arrow key (navigation keypad))
	K_LEFT        Keycode = C.SDLK_LEFT        // "Left" (the Left arrow key (navigation keypad))
	K_DOWN        Keycode = C.SDLK_DOWN        // "Down" (the Down arrow key (navigation keypad))
	K_UP          Keycode = C.SDLK_UP          // "Up" (the Up arrow key (navigation keypad))

	K_NUMLOCKCLEAR Keycode = C.SDLK_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	K_KP_DIVIDE    Keycode = C.SDLK_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	K_KP_MULTIPLY  Keycode = C.SDLK_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	K_KP_MINUS     Keycode = C.SDLK_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	K_KP_PLUS      Keycode = C.SDLK_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	K_KP_ENTER     Keycode = C.SDLK_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	K_KP_1         Keycode = C.SDLK_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	K_KP_2         Keycode = C.SDLK_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	K_KP_3         Keycode = C.SDLK_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	K_KP_4         Keycode = C.SDLK_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	K_KP_5         Keycode = C.SDLK_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	K_KP_6         Keycode = C.SDLK_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	K_KP_7         Keycode = C.SDLK_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	K_KP_8         Keycode = C.SDLK_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	K_KP_9         Keycode = C.SDLK_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	K_KP_0         Keycode = C.SDLK_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	K_KP_PERIOD    Keycode = C.SDLK_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	K_APPLICATION    Keycode = C.SDLK_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	K_POWER          Keycode = C.SDLK_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	K_KP_EQUALS      Keycode = C.SDLK_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	K_F13            Keycode = C.SDLK_F13            // "F13"
	K_F14            Keycode = C.SDLK_F14            // "F14"
	K_F15            Keycode = C.SDLK_F15            // "F15"
	K_F16            Keycode = C.SDLK_F16            // "F16"
	K_F17            Keycode = C.SDLK_F17            // "F17"
	K_F18            Keycode = C.SDLK_F18            // "F18"
	K_F19            Keycode = C.SDLK_F19            // "F19"
	K_F20            Keycode = C.SDLK_F20            // "F20"
	K_F21            Keycode = C.SDLK_F21            // "F21"
	K_F22            Keycode = C.SDLK_F22            // "F22"
	K_F23            Keycode = C.SDLK_F23            // "F23"
	K_F24            Keycode = C.SDLK_F24            // "F24"
	K_EXECUTE        Keycode = C.SDLK_EXECUTE        // "Execute"
	K_HELP           Keycode = C.SDLK_HELP           // "Help"
	K_MENU           Keycode = C.SDLK_MENU           // "Menu"
	K_SELECT         Keycode = C.SDLK_SELECT         // "Select"
	K_STOP           Keycode = C.SDLK_STOP           // "Stop"
	K_AGAIN          Keycode = C.SDLK_AGAIN          // "Again" (the Again key (Redo))
	K_UNDO           Keycode = C.SDLK_UNDO           // "Undo"
	K_CUT            Keycode = C.SDLK_CUT            // "Cut"
	K_COPY           Keycode = C.SDLK_COPY           // "Copy"
	K_PASTE          Keycode = C.SDLK_PASTE          // "Paste"
	K_FIND           Keycode = C.SDLK_FIND           // "Find"
	K_MUTE           Keycode = C.SDLK_MUTE           // "Mute"
	K_VOLUMEUP       Keycode = C.SDLK_VOLUMEUP       // "VolumeUp"
	K_VOLUMEDOWN     Keycode = C.SDLK_VOLUMEDOWN     // "VolumeDown"
	K_KP_COMMA       Keycode = C.SDLK_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	K_KP_EQUALSAS400 Keycode = C.SDLK_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	K_ALTERASE   Keycode = C.SDLK_ALTERASE   // "AltErase" (Erase-Eaze)
	K_SYSREQ     Keycode = C.SDLK_SYSREQ     // "SysReq" (the SysReq key)
	K_CANCEL     Keycode = C.SDLK_CANCEL     // "Cancel"
	K_CLEAR      Keycode = C.SDLK_CLEAR      // "Clear"
	K_PRIOR      Keycode = C.SDLK_PRIOR      // "Prior"
	K_RETURN2    Keycode = C.SDLK_RETURN2    // "Return"
	K_SEPARATOR  Keycode = C.SDLK_SEPARATOR  // "Separator"
	K_OUT        Keycode = C.SDLK_OUT        // "Out"
	K_OPER       Keycode = C.SDLK_OPER       // "Oper"
	K_CLEARAGAIN Keycode = C.SDLK_CLEARAGAIN // "Clear / Again"
	K_CRSEL      Keycode = C.SDLK_CRSEL      // "CrSel"
	K_EXSEL      Keycode = C.SDLK_EXSEL      // "ExSel"

	K_KP_00              Keycode = C.SDLK_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	K_KP_000             Keycode = C.SDLK_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	K_THOUSANDSSEPARATOR Keycode = C.SDLK_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	K_DECIMALSEPARATOR   Keycode = C.SDLK_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	K_CURRENCYUNIT       Keycode = C.SDLK_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	K_CURRENCYSUBUNIT    Keycode = C.SDLK_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	K_KP_LEFTPAREN       Keycode = C.SDLK_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	K_KP_RIGHTPAREN      Keycode = C.SDLK_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	K_KP_LEFTBRACE       Keycode = C.SDLK_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	K_KP_RIGHTBRACE      Keycode = C.SDLK_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	K_KP_TAB             Keycode = C.SDLK_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	K_KP_BACKSPACE       Keycode = C.SDLK_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	K_KP_A               Keycode = C.SDLK_KP_A               // "Keypad A" (the A key (numeric keypad))
	K_KP_B               Keycode = C.SDLK_KP_B               // "Keypad B" (the B key (numeric keypad))
	K_KP_C               Keycode = C.SDLK_KP_C               // "Keypad C" (the C key (numeric keypad))
	K_KP_D               Keycode = C.SDLK_KP_D               // "Keypad D" (the D key (numeric keypad))
	K_KP_E               Keycode = C.SDLK_KP_E               // "Keypad E" (the E key (numeric keypad))
	K_KP_F               Keycode = C.SDLK_KP_F               // "Keypad F" (the F key (numeric keypad))
	K_KP_XOR             Keycode = C.SDLK_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	K_KP_POWER           Keycode = C.SDLK_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	K_KP_PERCENT         Keycode = C.SDLK_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	K_KP_LESS            Keycode = C.SDLK_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	K_KP_GREATER         Keycode = C.SDLK_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	K_KP_AMPERSAND       Keycode = C.SDLK_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	K_KP_DBLAMPERSAND    Keycode = C.SDLK_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	K_KP_VERTICALBAR     Keycode = C.SDLK_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	K_KP_DBLVERTICALBAR  Keycode = C.SDLK_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	K_KP_COLON           Keycode = C.SDLK_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	K_KP_HASH            Keycode = C.SDLK_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	K_KP_SPACE           Keycode = C.SDLK_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	K_KP_AT              Keycode = C.SDLK_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	K_KP_EXCLAM          Keycode = C.SDLK_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	K_KP_MEMSTORE        Keycode = C.SDLK_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	K_KP_MEMRECALL       Keycode = C.SDLK_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	K_KP_MEMCLEAR        Keycode = C.SDLK_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	K_KP_MEMADD          Keycode = C.SDLK_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	K_KP_MEMSUBTRACT     Keycode = C.SDLK_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	K_KP_MEMMULTIPLY     Keycode = C.SDLK_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	K_KP_MEMDIVIDE       Keycode = C.SDLK_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	K_KP_PLUSMINUS       Keycode = C.SDLK_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	K_KP_CLEAR           Keycode = C.SDLK_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	K_KP_CLEARENTRY      Keycode = C.SDLK_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	K_KP_BINARY          Keycode = C.SDLK_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	K_KP_OCTAL           Keycode = C.SDLK_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	K_KP_DECIMAL         Keycode = C.SDLK_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	K_KP_HEXADECIMAL     Keycode = C.SDLK_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	K_LCTRL  Keycode = C.SDLK_LCTRL  // "Left Ctrl"
	K_LSHIFT Keycode = C.SDLK_LSHIFT // "Left Shift"
	K_LALT   Keycode = C.SDLK_LALT   // "Left Alt" (alt, option)
	K_LGUI   Keycode = C.SDLK_LGUI   // "Left GUI" (windows, command (apple), meta)
	K_RCTRL  Keycode = C.SDLK_RCTRL  // "Right Ctrl"
	K_RSHIFT Keycode = C.SDLK_RSHIFT // "Right Shift"
	K_RALT   Keycode = C.SDLK_RALT   // "Right Alt" (alt, option)
	K_RGUI   Keycode = C.SDLK_RGUI   // "Right GUI" (windows, command (apple), meta)

	K_MODE Keycode = C.SDLK_MODE // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)

	K_AUDIONEXT    Keycode = C.SDLK_AUDIONEXT    // "AudioNext" (the Next Track media key)
	K_AUDIOPREV    Keycode = C.SDLK_AUDIOPREV    // "AudioPrev" (the Previous Track media key)
	K_AUDIOSTOP    Keycode = C.SDLK_AUDIOSTOP    // "AudioStop" (the Stop media key)
	K_AUDIOPLAY    Keycode = C.SDLK_AUDIOPLAY    // "AudioPlay" (the Play media key)
	K_AUDIOMUTE    Keycode = C.SDLK_AUDIOMUTE    // "AudioMute" (the Mute volume key)
	K_MEDIASELECT  Keycode = C.SDLK_MEDIASELECT  // "MediaSelect" (the Media Select key)
	K_WWW          Keycode = C.SDLK_WWW          // "WWW" (the WWW/World Wide Web key)
	K_MAIL         Keycode = C.SDLK_MAIL         // "Mail" (the Mail/eMail key)
	K_CALCULATOR   Keycode = C.SDLK_CALCULATOR   // "Calculator" (the Calculator key)
	K_COMPUTER     Keycode = C.SDLK_COMPUTER     // "Computer" (the My Computer key)
	K_AC_SEARCH    Keycode = C.SDLK_AC_SEARCH    // "AC Search" (the Search key (application control keypad))
	K_AC_HOME      Keycode = C.SDLK_AC_HOME      // "AC Home" (the Home key (application control keypad))
	K_AC_BACK      Keycode = C.SDLK_AC_BACK      // "AC Back" (the Back key (application control keypad))
	K_AC_FORWARD   Keycode = C.SDLK_AC_FORWARD   // "AC Forward" (the Forward key (application control keypad))
	K_AC_STOP      Keycode = C.SDLK_AC_STOP      // "AC Stop" (the Stop key (application control keypad))
	K_AC_REFRESH   Keycode = C.SDLK_AC_REFRESH   // "AC Refresh" (the Refresh key (application control keypad))
	K_AC_BOOKMARKS Keycode = C.SDLK_AC_BOOKMARKS // "AC Bookmarks" (the Bookmarks key (application control keypad))

	K_BRIGHTNESSDOWN Keycode = C.SDLK_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	K_BRIGHTNESSUP   Keycode = C.SDLK_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	K_DISPLAYSWITCH  Keycode = C.SDLK_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	K_KBDILLUMTOGGLE Keycode = C.SDLK_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	K_KBDILLUMDOWN   Keycode = C.SDLK_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	K_KBDILLUMUP     Keycode = C.SDLK_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	K_EJECT          Keycode = C.SDLK_EJECT          // "Eject" (the Eject key)
	K_SLEEP          Keycode = C.SDLK_SLEEP          // "Sleep" (the Sleep key)
)

// An enumeration of key modifier masks.
// (https://wiki.libsdl.org/SDL_Keymod)
type Keymod C.SDL_Keymod

const (
	KMOD_NONE     Keymod = C.KMOD_NONE     // 0 (no modifier is applicable)
	KMOD_LSHIFT   Keymod = C.KMOD_LSHIFT   // the left Shift key is down
	KMOD_RSHIFT   Keymod = C.KMOD_RSHIFT   // the right Shift key is down
	KMOD_LCTRL    Keymod = C.KMOD_LCTRL    // the left Ctrl (Control) key is down
	KMOD_RCTRL    Keymod = C.KMOD_RCTRL    // the right Ctrl (Control) key is down
	KMOD_LALT     Keymod = C.KMOD_LALT     // the left Alt key is down
	KMOD_RALT     Keymod = C.KMOD_RALT     // the right Alt key is down
	KMOD_LGUI     Keymod = C.KMOD_LGUI     // the left GUI key (often the Windows key) is down
	KMOD_RGUI     Keymod = C.KMOD_RGUI     // the right GUI key (often the Windows key) is down
	KMOD_NUM      Keymod = C.KMOD_NUM      // the Num Lock key (may be located on an extended keypad) is down
	KMOD_CAPS     Keymod = C.KMOD_CAPS     // the Caps Lock key is down
	KMOD_MODE     Keymod = C.KMOD_MODE     // the AltGr key is down
	KMOD_CTRL     Keymod = C.KMOD_CTRL     // (KMOD_LCTRL|KMOD_RCTRL)
	KMOD_SHIFT    Keymod = C.KMOD_SHIFT    // (KMOD_LSHIFT|KMOD_RSHIFT)
	KMOD_ALT      Keymod = C.KMOD_ALT      // (KMOD_LALT|KMOD_RALT)
	KMOD_GUI      Keymod = C.KMOD_GUI      // (KMOD_LGUI|KMOD_RGUI)
	KMOD_RESERVED Keymod = C.KMOD_RESERVED // reserved for future use
)

func (code Keycode) c() C.SDL_Keycode {
	return C.SDL_Keycode(code)
}

func (mod Keymod) c() C.SDL_Keymod {
	return C.SDL_Keymod(mod)
}
