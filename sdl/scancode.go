package sdl

// #include "sdl_wrapper.h"
import "C"

// The SDL keyboard scancode representation.
// (https://wiki.libsdl.org/SDL_Scancode)
// (https://wiki.libsdl.org/SDLScancodeLookup)
const (
	SCANCODE_UNKNOWN = 0 // "" (no name, empty string)

	SCANCODE_A = C.SDL_SCANCODE_A // "A"
	SCANCODE_B = C.SDL_SCANCODE_B // "B"
	SCANCODE_C = C.SDL_SCANCODE_C // "C"
	SCANCODE_D = C.SDL_SCANCODE_D // "D"
	SCANCODE_E = C.SDL_SCANCODE_E // "E"
	SCANCODE_F = C.SDL_SCANCODE_F // "F"
	SCANCODE_G = C.SDL_SCANCODE_G // "G"
	SCANCODE_H = C.SDL_SCANCODE_H // "H"
	SCANCODE_I = C.SDL_SCANCODE_I // "I"
	SCANCODE_J = C.SDL_SCANCODE_J // "J"
	SCANCODE_K = C.SDL_SCANCODE_K // "K"
	SCANCODE_L = C.SDL_SCANCODE_L // "L"
	SCANCODE_M = C.SDL_SCANCODE_M // "M"
	SCANCODE_N = C.SDL_SCANCODE_N // "N"
	SCANCODE_O = C.SDL_SCANCODE_O // "O"
	SCANCODE_P = C.SDL_SCANCODE_P // "P"
	SCANCODE_Q = C.SDL_SCANCODE_Q // "Q"
	SCANCODE_R = C.SDL_SCANCODE_R // "R"
	SCANCODE_S = C.SDL_SCANCODE_S // "S"
	SCANCODE_T = C.SDL_SCANCODE_T // "T"
	SCANCODE_U = C.SDL_SCANCODE_U // "U"
	SCANCODE_V = C.SDL_SCANCODE_V // "V"
	SCANCODE_W = C.SDL_SCANCODE_W // "W"
	SCANCODE_X = C.SDL_SCANCODE_X // "X"
	SCANCODE_Y = C.SDL_SCANCODE_Y // "Y"
	SCANCODE_Z = C.SDL_SCANCODE_Z // "Z"

	SCANCODE_1 = C.SDL_SCANCODE_1 // "1"
	SCANCODE_2 = C.SDL_SCANCODE_2 // "2"
	SCANCODE_3 = C.SDL_SCANCODE_3 // "3"
	SCANCODE_4 = C.SDL_SCANCODE_4 // "4"
	SCANCODE_5 = C.SDL_SCANCODE_5 // "5"
	SCANCODE_6 = C.SDL_SCANCODE_6 // "6"
	SCANCODE_7 = C.SDL_SCANCODE_7 // "7"
	SCANCODE_8 = C.SDL_SCANCODE_8 // "8"
	SCANCODE_9 = C.SDL_SCANCODE_9 // "9"
	SCANCODE_0 = C.SDL_SCANCODE_0 // "0"

	SCANCODE_RETURN    = C.SDL_SCANCODE_RETURN    // "Return"
	SCANCODE_ESCAPE    = C.SDL_SCANCODE_ESCAPE    // "Escape" (the Esc key)
	SCANCODE_BACKSPACE = C.SDL_SCANCODE_BACKSPACE // "Backspace"
	SCANCODE_TAB       = C.SDL_SCANCODE_TAB       // "Tab" (the Tab key)
	SCANCODE_SPACE     = C.SDL_SCANCODE_SPACE     // "Space" (the Space Bar key(s))

	SCANCODE_MINUS        = C.SDL_SCANCODE_MINUS        // "-"
	SCANCODE_EQUALS       = C.SDL_SCANCODE_EQUALS       // "="
	SCANCODE_LEFTBRACKET  = C.SDL_SCANCODE_LEFTBRACKET  // "["
	SCANCODE_RIGHTBRACKET = C.SDL_SCANCODE_RIGHTBRACKET // "]"
	SCANCODE_BACKSLASH    = C.SDL_SCANCODE_BACKSLASH    // "\"
	SCANCODE_NONUSHASH    = C.SDL_SCANCODE_NONUSHASH    // "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
	SCANCODE_SEMICOLON    = C.SDL_SCANCODE_SEMICOLON    // ";"
	SCANCODE_APOSTROPHE   = C.SDL_SCANCODE_APOSTROPHE   // "'"
	SCANCODE_GRAVE        = C.SDL_SCANCODE_GRAVE        // "`"
	SCANCODE_COMMA        = C.SDL_SCANCODE_COMMA        // ","
	SCANCODE_PERIOD       = C.SDL_SCANCODE_PERIOD       // "."
	SCANCODE_SLASH        = C.SDL_SCANCODE_SLASH        // "/"
	SCANCODE_CAPSLOCK     = C.SDL_SCANCODE_CAPSLOCK     // "CapsLock"
	SCANCODE_F1           = C.SDL_SCANCODE_F1           // "F1"
	SCANCODE_F2           = C.SDL_SCANCODE_F2           // "F2"
	SCANCODE_F3           = C.SDL_SCANCODE_F3           // "F3"
	SCANCODE_F4           = C.SDL_SCANCODE_F4           // "F4"
	SCANCODE_F5           = C.SDL_SCANCODE_F5           // "F5"
	SCANCODE_F6           = C.SDL_SCANCODE_F6           // "F6"
	SCANCODE_F7           = C.SDL_SCANCODE_F7           // "F7"
	SCANCODE_F8           = C.SDL_SCANCODE_F8           // "F8"
	SCANCODE_F9           = C.SDL_SCANCODE_F9           // "F9"
	SCANCODE_F10          = C.SDL_SCANCODE_F10          // "F10"
	SCANCODE_F11          = C.SDL_SCANCODE_F11          // "F11"
	SCANCODE_F12          = C.SDL_SCANCODE_F12          // "F12"
	SCANCODE_PRINTSCREEN  = C.SDL_SCANCODE_PRINTSCREEN  // "PrintScreen"
	SCANCODE_SCROLLLOCK   = C.SDL_SCANCODE_SCROLLLOCK   // "ScrollLock"
	SCANCODE_PAUSE        = C.SDL_SCANCODE_PAUSE        // "Pause" (the Pause / Break key)
	SCANCODE_INSERT       = C.SDL_SCANCODE_INSERT       // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	SCANCODE_HOME         = C.SDL_SCANCODE_HOME         // "Home"
	SCANCODE_PAGEUP       = C.SDL_SCANCODE_PAGEUP       // "PageUp"
	SCANCODE_DELETE       = C.SDL_SCANCODE_DELETE       // "Delete"
	SCANCODE_END          = C.SDL_SCANCODE_END          // "End"
	SCANCODE_PAGEDOWN     = C.SDL_SCANCODE_PAGEDOWN     // "PageDown"
	SCANCODE_RIGHT        = C.SDL_SCANCODE_RIGHT        // "Right" (the Right arrow key (navigation keypad))
	SCANCODE_LEFT         = C.SDL_SCANCODE_LEFT         // "Left" (the Left arrow key (navigation keypad))
	SCANCODE_DOWN         = C.SDL_SCANCODE_DOWN         // "Down" (the Down arrow key (navigation keypad))
	SCANCODE_UP           = C.SDL_SCANCODE_UP           // "Up" (the Up arrow key (navigation keypad))

	SCANCODE_NUMLOCKCLEAR = C.SDL_SCANCODE_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	SCANCODE_KP_DIVIDE    = C.SDL_SCANCODE_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	SCANCODE_KP_MULTIPLY  = C.SDL_SCANCODE_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	SCANCODE_KP_MINUS     = C.SDL_SCANCODE_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	SCANCODE_KP_PLUS      = C.SDL_SCANCODE_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	SCANCODE_KP_ENTER     = C.SDL_SCANCODE_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	SCANCODE_KP_1         = C.SDL_SCANCODE_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	SCANCODE_KP_2         = C.SDL_SCANCODE_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	SCANCODE_KP_3         = C.SDL_SCANCODE_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	SCANCODE_KP_4         = C.SDL_SCANCODE_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	SCANCODE_KP_5         = C.SDL_SCANCODE_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	SCANCODE_KP_6         = C.SDL_SCANCODE_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	SCANCODE_KP_7         = C.SDL_SCANCODE_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	SCANCODE_KP_8         = C.SDL_SCANCODE_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	SCANCODE_KP_9         = C.SDL_SCANCODE_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	SCANCODE_KP_0         = C.SDL_SCANCODE_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	SCANCODE_KP_PERIOD    = C.SDL_SCANCODE_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	SCANCODE_NONUSBACKSLASH = C.SDL_SCANCODE_NONUSBACKSLASH // "" (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.)
	SCANCODE_APPLICATION    = C.SDL_SCANCODE_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	SCANCODE_POWER          = C.SDL_SCANCODE_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	SCANCODE_KP_EQUALS      = C.SDL_SCANCODE_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	SCANCODE_F13            = C.SDL_SCANCODE_F13            // "F13"
	SCANCODE_F14            = C.SDL_SCANCODE_F14            // "F14"
	SCANCODE_F15            = C.SDL_SCANCODE_F15            // "F15"
	SCANCODE_F16            = C.SDL_SCANCODE_F16            // "F16"
	SCANCODE_F17            = C.SDL_SCANCODE_F17            // "F17"
	SCANCODE_F18            = C.SDL_SCANCODE_F18            // "F18"
	SCANCODE_F19            = C.SDL_SCANCODE_F19            // "F19"
	SCANCODE_F20            = C.SDL_SCANCODE_F20            // "F20"
	SCANCODE_F21            = C.SDL_SCANCODE_F21            // "F21"
	SCANCODE_F22            = C.SDL_SCANCODE_F22            // "F22"
	SCANCODE_F23            = C.SDL_SCANCODE_F23            // "F23"
	SCANCODE_F24            = C.SDL_SCANCODE_F24            // "F24"
	SCANCODE_EXECUTE        = C.SDL_SCANCODE_EXECUTE        // "Execute"
	SCANCODE_HELP           = C.SDL_SCANCODE_HELP           // "Help"
	SCANCODE_MENU           = C.SDL_SCANCODE_MENU           // "Menu"
	SCANCODE_SELECT         = C.SDL_SCANCODE_SELECT         // "Select"
	SCANCODE_STOP           = C.SDL_SCANCODE_STOP           // "Stop"
	SCANCODE_AGAIN          = C.SDL_SCANCODE_AGAIN          // "Again" (the Again key (Redo))
	SCANCODE_UNDO           = C.SDL_SCANCODE_UNDO           // "Undo"
	SCANCODE_CUT            = C.SDL_SCANCODE_CUT            // "Cut"
	SCANCODE_COPY           = C.SDL_SCANCODE_COPY           // "Copy"
	SCANCODE_PASTE          = C.SDL_SCANCODE_PASTE          // "Paste"
	SCANCODE_FIND           = C.SDL_SCANCODE_FIND           // "Find"
	SCANCODE_MUTE           = C.SDL_SCANCODE_MUTE           // "Mute"
	SCANCODE_VOLUMEUP       = C.SDL_SCANCODE_VOLUMEUP       // "VolumeUp"
	SCANCODE_VOLUMEDOWN     = C.SDL_SCANCODE_VOLUMEDOWN     // "VolumeDown"
	SCANCODE_KP_COMMA       = C.SDL_SCANCODE_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	SCANCODE_KP_EQUALSAS400 = C.SDL_SCANCODE_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	SCANCODE_INTERNATIONAL1 = C.SDL_SCANCODE_INTERNATIONAL1 // "" (no name, empty string; used on Asian keyboards, see footnotes in USB doc)
	SCANCODE_INTERNATIONAL2 = C.SDL_SCANCODE_INTERNATIONAL2 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL3 = C.SDL_SCANCODE_INTERNATIONAL3 // "" (no name, empty string; Yen)
	SCANCODE_INTERNATIONAL4 = C.SDL_SCANCODE_INTERNATIONAL4 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL5 = C.SDL_SCANCODE_INTERNATIONAL5 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL6 = C.SDL_SCANCODE_INTERNATIONAL6 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL7 = C.SDL_SCANCODE_INTERNATIONAL7 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL8 = C.SDL_SCANCODE_INTERNATIONAL8 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL9 = C.SDL_SCANCODE_INTERNATIONAL9 // "" (no name, empty string)
	SCANCODE_LANG1          = C.SDL_SCANCODE_LANG1          // "" (no name, empty string; Hangul/English toggle)
	SCANCODE_LANG2          = C.SDL_SCANCODE_LANG2          // "" (no name, empty string; Hanja conversion)
	SCANCODE_LANG3          = C.SDL_SCANCODE_LANG3          // "" (no name, empty string; Katakana)
	SCANCODE_LANG4          = C.SDL_SCANCODE_LANG4          // "" (no name, empty string; Hiragana)
	SCANCODE_LANG5          = C.SDL_SCANCODE_LANG5          // "" (no name, empty string; Zenkaku/Hankaku)
	SCANCODE_LANG6          = C.SDL_SCANCODE_LANG6          // "" (no name, empty string; reserved)
	SCANCODE_LANG7          = C.SDL_SCANCODE_LANG7          // "" (no name, empty string; reserved)
	SCANCODE_LANG8          = C.SDL_SCANCODE_LANG8          // "" (no name, empty string; reserved)
	SCANCODE_LANG9          = C.SDL_SCANCODE_LANG9          // "" (no name, empty string; reserved)

	SCANCODE_ALTERASE   = C.SDL_SCANCODE_ALTERASE   // "AltErase" (Erase-Eaze)
	SCANCODE_SYSREQ     = C.SDL_SCANCODE_SYSREQ     // "SysReq" (the SysReq key)
	SCANCODE_CANCEL     = C.SDL_SCANCODE_CANCEL     // "Cancel"
	SCANCODE_CLEAR      = C.SDL_SCANCODE_CLEAR      // "Clear"
	SCANCODE_PRIOR      = C.SDL_SCANCODE_PRIOR      // "Prior"
	SCANCODE_RETURN2    = C.SDL_SCANCODE_RETURN2    // "Return"
	SCANCODE_SEPARATOR  = C.SDL_SCANCODE_SEPARATOR  // "Separator"
	SCANCODE_OUT        = C.SDL_SCANCODE_OUT        // "Out"
	SCANCODE_OPER       = C.SDL_SCANCODE_OPER       // "Oper"
	SCANCODE_CLEARAGAIN = C.SDL_SCANCODE_CLEARAGAIN // "Clear / Again"
	SCANCODE_CRSEL      = C.SDL_SCANCODE_CRSEL      // "CrSel"
	SCANCODE_EXSEL      = C.SDL_SCANCODE_EXSEL      // "ExSel"

	SCANCODE_KP_00              = C.SDL_SCANCODE_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	SCANCODE_KP_000             = C.SDL_SCANCODE_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	SCANCODE_THOUSANDSSEPARATOR = C.SDL_SCANCODE_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	SCANCODE_DECIMALSEPARATOR   = C.SDL_SCANCODE_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	SCANCODE_CURRENCYUNIT       = C.SDL_SCANCODE_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	SCANCODE_CURRENCYSUBUNIT    = C.SDL_SCANCODE_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	SCANCODE_KP_LEFTPAREN       = C.SDL_SCANCODE_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	SCANCODE_KP_RIGHTPAREN      = C.SDL_SCANCODE_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	SCANCODE_KP_LEFTBRACE       = C.SDL_SCANCODE_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	SCANCODE_KP_RIGHTBRACE      = C.SDL_SCANCODE_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	SCANCODE_KP_TAB             = C.SDL_SCANCODE_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	SCANCODE_KP_BACKSPACE       = C.SDL_SCANCODE_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	SCANCODE_KP_A               = C.SDL_SCANCODE_KP_A               // "Keypad A" (the A key (numeric keypad))
	SCANCODE_KP_B               = C.SDL_SCANCODE_KP_B               // "Keypad B" (the B key (numeric keypad))
	SCANCODE_KP_C               = C.SDL_SCANCODE_KP_C               // "Keypad C" (the C key (numeric keypad))
	SCANCODE_KP_D               = C.SDL_SCANCODE_KP_D               // "Keypad D" (the D key (numeric keypad))
	SCANCODE_KP_E               = C.SDL_SCANCODE_KP_E               // "Keypad E" (the E key (numeric keypad))
	SCANCODE_KP_F               = C.SDL_SCANCODE_KP_F               // "Keypad F" (the F key (numeric keypad))
	SCANCODE_KP_XOR             = C.SDL_SCANCODE_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	SCANCODE_KP_POWER           = C.SDL_SCANCODE_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	SCANCODE_KP_PERCENT         = C.SDL_SCANCODE_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	SCANCODE_KP_LESS            = C.SDL_SCANCODE_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	SCANCODE_KP_GREATER         = C.SDL_SCANCODE_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	SCANCODE_KP_AMPERSAND       = C.SDL_SCANCODE_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	SCANCODE_KP_DBLAMPERSAND    = C.SDL_SCANCODE_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	SCANCODE_KP_VERTICALBAR     = C.SDL_SCANCODE_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	SCANCODE_KP_DBLVERTICALBAR  = C.SDL_SCANCODE_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	SCANCODE_KP_COLON           = C.SDL_SCANCODE_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	SCANCODE_KP_HASH            = C.SDL_SCANCODE_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	SCANCODE_KP_SPACE           = C.SDL_SCANCODE_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	SCANCODE_KP_AT              = C.SDL_SCANCODE_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	SCANCODE_KP_EXCLAM          = C.SDL_SCANCODE_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	SCANCODE_KP_MEMSTORE        = C.SDL_SCANCODE_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	SCANCODE_KP_MEMRECALL       = C.SDL_SCANCODE_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	SCANCODE_KP_MEMCLEAR        = C.SDL_SCANCODE_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	SCANCODE_KP_MEMADD          = C.SDL_SCANCODE_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	SCANCODE_KP_MEMSUBTRACT     = C.SDL_SCANCODE_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	SCANCODE_KP_MEMMULTIPLY     = C.SDL_SCANCODE_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	SCANCODE_KP_MEMDIVIDE       = C.SDL_SCANCODE_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	SCANCODE_KP_PLUSMINUS       = C.SDL_SCANCODE_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	SCANCODE_KP_CLEAR           = C.SDL_SCANCODE_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	SCANCODE_KP_CLEARENTRY      = C.SDL_SCANCODE_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	SCANCODE_KP_BINARY          = C.SDL_SCANCODE_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	SCANCODE_KP_OCTAL           = C.SDL_SCANCODE_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	SCANCODE_KP_DECIMAL         = C.SDL_SCANCODE_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	SCANCODE_KP_HEXADECIMAL     = C.SDL_SCANCODE_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	SCANCODE_LCTRL          = C.SDL_SCANCODE_LCTRL          // "Left Ctrl"
	SCANCODE_LSHIFT         = C.SDL_SCANCODE_LSHIFT         // "Left Shift"
	SCANCODE_LALT           = C.SDL_SCANCODE_LALT           // "Left Alt" (alt, option)
	SCANCODE_LGUI           = C.SDL_SCANCODE_LGUI           // "Left GUI" (windows, command (apple), meta)
	SCANCODE_RCTRL          = C.SDL_SCANCODE_RCTRL          // "Right Ctrl"
	SCANCODE_RSHIFT         = C.SDL_SCANCODE_RSHIFT         // "Right Shift"
	SCANCODE_RALT           = C.SDL_SCANCODE_RALT           // "Right Alt" (alt gr, option)
	SCANCODE_RGUI           = C.SDL_SCANCODE_RGUI           // "Right GUI" (windows, command (apple), meta)
	SCANCODE_MODE           = C.SDL_SCANCODE_MODE           // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	SCANCODE_AUDIONEXT      = C.SDL_SCANCODE_AUDIONEXT      // "AudioNext" (the Next Track media key)
	SCANCODE_AUDIOPREV      = C.SDL_SCANCODE_AUDIOPREV      // "AudioPrev" (the Previous Track media key)
	SCANCODE_AUDIOSTOP      = C.SDL_SCANCODE_AUDIOSTOP      // "AudioStop" (the Stop media key)
	SCANCODE_AUDIOPLAY      = C.SDL_SCANCODE_AUDIOPLAY      // "AudioPlay" (the Play media key)
	SCANCODE_AUDIOMUTE      = C.SDL_SCANCODE_AUDIOMUTE      // "AudioMute" (the Mute volume key)
	SCANCODE_MEDIASELECT    = C.SDL_SCANCODE_MEDIASELECT    // "MediaSelect" (the Media Select key)
	SCANCODE_WWW            = C.SDL_SCANCODE_WWW            // "WWW" (the WWW/World Wide Web key)
	SCANCODE_MAIL           = C.SDL_SCANCODE_MAIL           // "Mail" (the Mail/eMail key)
	SCANCODE_CALCULATOR     = C.SDL_SCANCODE_CALCULATOR     // "Calculator" (the Calculator key)
	SCANCODE_COMPUTER       = C.SDL_SCANCODE_COMPUTER       // "Computer" (the My Computer key)
	SCANCODE_AC_SEARCH      = C.SDL_SCANCODE_AC_SEARCH      // "AC Search" (the Search key (application control keypad))
	SCANCODE_AC_HOME        = C.SDL_SCANCODE_AC_HOME        // "AC Home" (the Home key (application control keypad))
	SCANCODE_AC_BACK        = C.SDL_SCANCODE_AC_BACK        // "AC Back" (the Back key (application control keypad))
	SCANCODE_AC_FORWARD     = C.SDL_SCANCODE_AC_FORWARD     // "AC Forward" (the Forward key (application control keypad))
	SCANCODE_AC_STOP        = C.SDL_SCANCODE_AC_STOP        // "AC Stop" (the Stop key (application control keypad))
	SCANCODE_AC_REFRESH     = C.SDL_SCANCODE_AC_REFRESH     // "AC Refresh" (the Refresh key (application control keypad))
	SCANCODE_AC_BOOKMARKS   = C.SDL_SCANCODE_AC_BOOKMARKS   // "AC Bookmarks" (the Bookmarks key (application control keypad))
	SCANCODE_BRIGHTNESSDOWN = C.SDL_SCANCODE_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	SCANCODE_BRIGHTNESSUP   = C.SDL_SCANCODE_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	SCANCODE_DISPLAYSWITCH  = C.SDL_SCANCODE_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	SCANCODE_KBDILLUMTOGGLE = C.SDL_SCANCODE_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	SCANCODE_KBDILLUMDOWN   = C.SDL_SCANCODE_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	SCANCODE_KBDILLUMUP     = C.SDL_SCANCODE_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	SCANCODE_EJECT          = C.SDL_SCANCODE_EJECT          // "Eject" (the Eject key)
	SCANCODE_SLEEP          = C.SDL_SCANCODE_SLEEP          // "Sleep" (the Sleep key)
	SCANCODE_APP1           = C.SDL_SCANCODE_APP1
	SCANCODE_APP2           = C.SDL_SCANCODE_APP2
	NUM_SCANCODES           = C.SDL_NUM_SCANCODES
)

// Scancode is an SDL keyboard scancode representation.
// (https://wiki.libsdl.org/SDL_Scancode)
type Scancode uint32

func (code Scancode) c() C.SDL_Scancode {
	return C.SDL_Scancode(code)
}
