package sdl

// #include "sdl_wrapper.h"
import "C"

// The SDL keyboard scancode representation.
// (https://wiki.libsdl.org/SDL_Scancode)
// (https://wiki.libsdl.org/SDLScancodeLookup)
type Scancode uint32

const (
	SCANCODE_UNKNOWN Scancode = C.SDL_SCANCODE_UNKNOWN // "" (no name, empty string)

	SCANCODE_A Scancode = C.SDL_SCANCODE_A // "A"
	SCANCODE_B Scancode = C.SDL_SCANCODE_B // "B"
	SCANCODE_C Scancode = C.SDL_SCANCODE_C // "C"
	SCANCODE_D Scancode = C.SDL_SCANCODE_D // "D"
	SCANCODE_E Scancode = C.SDL_SCANCODE_E // "E"
	SCANCODE_F Scancode = C.SDL_SCANCODE_F // "F"
	SCANCODE_G Scancode = C.SDL_SCANCODE_G // "G"
	SCANCODE_H Scancode = C.SDL_SCANCODE_H // "H"
	SCANCODE_I Scancode = C.SDL_SCANCODE_I // "I"
	SCANCODE_J Scancode = C.SDL_SCANCODE_J // "J"
	SCANCODE_K Scancode = C.SDL_SCANCODE_K // "K"
	SCANCODE_L Scancode = C.SDL_SCANCODE_L // "L"
	SCANCODE_M Scancode = C.SDL_SCANCODE_M // "M"
	SCANCODE_N Scancode = C.SDL_SCANCODE_N // "N"
	SCANCODE_O Scancode = C.SDL_SCANCODE_O // "O"
	SCANCODE_P Scancode = C.SDL_SCANCODE_P // "P"
	SCANCODE_Q Scancode = C.SDL_SCANCODE_Q // "Q"
	SCANCODE_R Scancode = C.SDL_SCANCODE_R // "R"
	SCANCODE_S Scancode = C.SDL_SCANCODE_S // "S"
	SCANCODE_T Scancode = C.SDL_SCANCODE_T // "T"
	SCANCODE_U Scancode = C.SDL_SCANCODE_U // "U"
	SCANCODE_V Scancode = C.SDL_SCANCODE_V // "V"
	SCANCODE_W Scancode = C.SDL_SCANCODE_W // "W"
	SCANCODE_X Scancode = C.SDL_SCANCODE_X // "X"
	SCANCODE_Y Scancode = C.SDL_SCANCODE_Y // "Y"
	SCANCODE_Z Scancode = C.SDL_SCANCODE_Z // "Z"

	SCANCODE_1 Scancode = C.SDL_SCANCODE_1 // "1"
	SCANCODE_2 Scancode = C.SDL_SCANCODE_2 // "2"
	SCANCODE_3 Scancode = C.SDL_SCANCODE_3 // "3"
	SCANCODE_4 Scancode = C.SDL_SCANCODE_4 // "4"
	SCANCODE_5 Scancode = C.SDL_SCANCODE_5 // "5"
	SCANCODE_6 Scancode = C.SDL_SCANCODE_6 // "6"
	SCANCODE_7 Scancode = C.SDL_SCANCODE_7 // "7"
	SCANCODE_8 Scancode = C.SDL_SCANCODE_8 // "8"
	SCANCODE_9 Scancode = C.SDL_SCANCODE_9 // "9"
	SCANCODE_0 Scancode = C.SDL_SCANCODE_0 // "0"

	SCANCODE_RETURN    Scancode = C.SDL_SCANCODE_RETURN    // "Return"
	SCANCODE_ESCAPE    Scancode = C.SDL_SCANCODE_ESCAPE    // "Escape" (the Esc key)
	SCANCODE_BACKSPACE Scancode = C.SDL_SCANCODE_BACKSPACE // "Backspace"
	SCANCODE_TAB       Scancode = C.SDL_SCANCODE_TAB       // "Tab" (the Tab key)
	SCANCODE_SPACE     Scancode = C.SDL_SCANCODE_SPACE     // "Space" (the Space Bar key(s))

	SCANCODE_MINUS        Scancode = C.SDL_SCANCODE_MINUS        // "-"
	SCANCODE_EQUALS       Scancode = C.SDL_SCANCODE_EQUALS       // "="
	SCANCODE_LEFTBRACKET  Scancode = C.SDL_SCANCODE_LEFTBRACKET  // "["
	SCANCODE_RIGHTBRACKET Scancode = C.SDL_SCANCODE_RIGHTBRACKET // "]"
	SCANCODE_BACKSLASH    Scancode = C.SDL_SCANCODE_BACKSLASH    // "\"
	SCANCODE_NONUSHASH    Scancode = C.SDL_SCANCODE_NONUSHASH    // "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
	SCANCODE_SEMICOLON    Scancode = C.SDL_SCANCODE_SEMICOLON    // ";"
	SCANCODE_APOSTROPHE   Scancode = C.SDL_SCANCODE_APOSTROPHE   // "'"
	SCANCODE_GRAVE        Scancode = C.SDL_SCANCODE_GRAVE        // "`"
	SCANCODE_COMMA        Scancode = C.SDL_SCANCODE_COMMA        // ","
	SCANCODE_PERIOD       Scancode = C.SDL_SCANCODE_PERIOD       // "."
	SCANCODE_SLASH        Scancode = C.SDL_SCANCODE_SLASH        // "/"
	SCANCODE_CAPSLOCK     Scancode = C.SDL_SCANCODE_CAPSLOCK     // "CapsLock"
	SCANCODE_F1           Scancode = C.SDL_SCANCODE_F1           // "F1"
	SCANCODE_F2           Scancode = C.SDL_SCANCODE_F2           // "F2"
	SCANCODE_F3           Scancode = C.SDL_SCANCODE_F3           // "F3"
	SCANCODE_F4           Scancode = C.SDL_SCANCODE_F4           // "F4"
	SCANCODE_F5           Scancode = C.SDL_SCANCODE_F5           // "F5"
	SCANCODE_F6           Scancode = C.SDL_SCANCODE_F6           // "F6"
	SCANCODE_F7           Scancode = C.SDL_SCANCODE_F7           // "F7"
	SCANCODE_F8           Scancode = C.SDL_SCANCODE_F8           // "F8"
	SCANCODE_F9           Scancode = C.SDL_SCANCODE_F9           // "F9"
	SCANCODE_F10          Scancode = C.SDL_SCANCODE_F10          // "F10"
	SCANCODE_F11          Scancode = C.SDL_SCANCODE_F11          // "F11"
	SCANCODE_F12          Scancode = C.SDL_SCANCODE_F12          // "F12"
	SCANCODE_PRINTSCREEN  Scancode = C.SDL_SCANCODE_PRINTSCREEN  // "PrintScreen"
	SCANCODE_SCROLLLOCK   Scancode = C.SDL_SCANCODE_SCROLLLOCK   // "ScrollLock"
	SCANCODE_PAUSE        Scancode = C.SDL_SCANCODE_PAUSE        // "Pause" (the Pause / Break key)
	SCANCODE_INSERT       Scancode = C.SDL_SCANCODE_INSERT       // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	SCANCODE_HOME         Scancode = C.SDL_SCANCODE_HOME         // "Home"
	SCANCODE_PAGEUP       Scancode = C.SDL_SCANCODE_PAGEUP       // "PageUp"
	SCANCODE_DELETE       Scancode = C.SDL_SCANCODE_DELETE       // "Delete"
	SCANCODE_END          Scancode = C.SDL_SCANCODE_END          // "End"
	SCANCODE_PAGEDOWN     Scancode = C.SDL_SCANCODE_PAGEDOWN     // "PageDown"
	SCANCODE_RIGHT        Scancode = C.SDL_SCANCODE_RIGHT        // "Right" (the Right arrow key (navigation keypad))
	SCANCODE_LEFT         Scancode = C.SDL_SCANCODE_LEFT         // "Left" (the Left arrow key (navigation keypad))
	SCANCODE_DOWN         Scancode = C.SDL_SCANCODE_DOWN         // "Down" (the Down arrow key (navigation keypad))
	SCANCODE_UP           Scancode = C.SDL_SCANCODE_UP           // "Up" (the Up arrow key (navigation keypad))

	SCANCODE_NUMLOCKCLEAR Scancode = C.SDL_SCANCODE_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	SCANCODE_KP_DIVIDE    Scancode = C.SDL_SCANCODE_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	SCANCODE_KP_MULTIPLY  Scancode = C.SDL_SCANCODE_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	SCANCODE_KP_MINUS     Scancode = C.SDL_SCANCODE_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	SCANCODE_KP_PLUS      Scancode = C.SDL_SCANCODE_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	SCANCODE_KP_ENTER     Scancode = C.SDL_SCANCODE_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	SCANCODE_KP_1         Scancode = C.SDL_SCANCODE_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	SCANCODE_KP_2         Scancode = C.SDL_SCANCODE_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	SCANCODE_KP_3         Scancode = C.SDL_SCANCODE_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	SCANCODE_KP_4         Scancode = C.SDL_SCANCODE_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	SCANCODE_KP_5         Scancode = C.SDL_SCANCODE_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	SCANCODE_KP_6         Scancode = C.SDL_SCANCODE_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	SCANCODE_KP_7         Scancode = C.SDL_SCANCODE_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	SCANCODE_KP_8         Scancode = C.SDL_SCANCODE_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	SCANCODE_KP_9         Scancode = C.SDL_SCANCODE_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	SCANCODE_KP_0         Scancode = C.SDL_SCANCODE_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	SCANCODE_KP_PERIOD    Scancode = C.SDL_SCANCODE_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	SCANCODE_NONUSBACKSLASH Scancode = C.SDL_SCANCODE_NONUSBACKSLASH // "" (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.)
	SCANCODE_APPLICATION    Scancode = C.SDL_SCANCODE_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	SCANCODE_POWER          Scancode = C.SDL_SCANCODE_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	SCANCODE_KP_EQUALS      Scancode = C.SDL_SCANCODE_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	SCANCODE_F13            Scancode = C.SDL_SCANCODE_F13            // "F13"
	SCANCODE_F14            Scancode = C.SDL_SCANCODE_F14            // "F14"
	SCANCODE_F15            Scancode = C.SDL_SCANCODE_F15            // "F15"
	SCANCODE_F16            Scancode = C.SDL_SCANCODE_F16            // "F16"
	SCANCODE_F17            Scancode = C.SDL_SCANCODE_F17            // "F17"
	SCANCODE_F18            Scancode = C.SDL_SCANCODE_F18            // "F18"
	SCANCODE_F19            Scancode = C.SDL_SCANCODE_F19            // "F19"
	SCANCODE_F20            Scancode = C.SDL_SCANCODE_F20            // "F20"
	SCANCODE_F21            Scancode = C.SDL_SCANCODE_F21            // "F21"
	SCANCODE_F22            Scancode = C.SDL_SCANCODE_F22            // "F22"
	SCANCODE_F23            Scancode = C.SDL_SCANCODE_F23            // "F23"
	SCANCODE_F24            Scancode = C.SDL_SCANCODE_F24            // "F24"
	SCANCODE_EXECUTE        Scancode = C.SDL_SCANCODE_EXECUTE        // "Execute"
	SCANCODE_HELP           Scancode = C.SDL_SCANCODE_HELP           // "Help"
	SCANCODE_MENU           Scancode = C.SDL_SCANCODE_MENU           // "Menu"
	SCANCODE_SELECT         Scancode = C.SDL_SCANCODE_SELECT         // "Select"
	SCANCODE_STOP           Scancode = C.SDL_SCANCODE_STOP           // "Stop"
	SCANCODE_AGAIN          Scancode = C.SDL_SCANCODE_AGAIN          // "Again" (the Again key (Redo))
	SCANCODE_UNDO           Scancode = C.SDL_SCANCODE_UNDO           // "Undo"
	SCANCODE_CUT            Scancode = C.SDL_SCANCODE_CUT            // "Cut"
	SCANCODE_COPY           Scancode = C.SDL_SCANCODE_COPY           // "Copy"
	SCANCODE_PASTE          Scancode = C.SDL_SCANCODE_PASTE          // "Paste"
	SCANCODE_FIND           Scancode = C.SDL_SCANCODE_FIND           // "Find"
	SCANCODE_MUTE           Scancode = C.SDL_SCANCODE_MUTE           // "Mute"
	SCANCODE_VOLUMEUP       Scancode = C.SDL_SCANCODE_VOLUMEUP       // "VolumeUp"
	SCANCODE_VOLUMEDOWN     Scancode = C.SDL_SCANCODE_VOLUMEDOWN     // "VolumeDown"
	SCANCODE_KP_COMMA       Scancode = C.SDL_SCANCODE_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	SCANCODE_KP_EQUALSAS400 Scancode = C.SDL_SCANCODE_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	SCANCODE_INTERNATIONAL1 Scancode = C.SDL_SCANCODE_INTERNATIONAL1 // "" (no name, empty string; used on Asian keyboards, see footnotes in USB doc)
	SCANCODE_INTERNATIONAL2 Scancode = C.SDL_SCANCODE_INTERNATIONAL2 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL3 Scancode = C.SDL_SCANCODE_INTERNATIONAL3 // "" (no name, empty string; Yen)
	SCANCODE_INTERNATIONAL4 Scancode = C.SDL_SCANCODE_INTERNATIONAL4 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL5 Scancode = C.SDL_SCANCODE_INTERNATIONAL5 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL6 Scancode = C.SDL_SCANCODE_INTERNATIONAL6 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL7 Scancode = C.SDL_SCANCODE_INTERNATIONAL7 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL8 Scancode = C.SDL_SCANCODE_INTERNATIONAL8 // "" (no name, empty string)
	SCANCODE_INTERNATIONAL9 Scancode = C.SDL_SCANCODE_INTERNATIONAL9 // "" (no name, empty string)
	SCANCODE_LANG1          Scancode = C.SDL_SCANCODE_LANG1          // "" (no name, empty string; Hangul/English toggle)
	SCANCODE_LANG2          Scancode = C.SDL_SCANCODE_LANG2          // "" (no name, empty string; Hanja conversion)
	SCANCODE_LANG3          Scancode = C.SDL_SCANCODE_LANG3          // "" (no name, empty string; Katakana)
	SCANCODE_LANG4          Scancode = C.SDL_SCANCODE_LANG4          // "" (no name, empty string; Hiragana)
	SCANCODE_LANG5          Scancode = C.SDL_SCANCODE_LANG5          // "" (no name, empty string; Zenkaku/Hankaku)
	SCANCODE_LANG6          Scancode = C.SDL_SCANCODE_LANG6          // "" (no name, empty string; reserved)
	SCANCODE_LANG7          Scancode = C.SDL_SCANCODE_LANG7          // "" (no name, empty string; reserved)
	SCANCODE_LANG8          Scancode = C.SDL_SCANCODE_LANG8          // "" (no name, empty string; reserved)
	SCANCODE_LANG9          Scancode = C.SDL_SCANCODE_LANG9          // "" (no name, empty string; reserved)

	SCANCODE_ALTERASE   Scancode = C.SDL_SCANCODE_ALTERASE   // "AltErase" (Erase-Eaze)
	SCANCODE_SYSREQ     Scancode = C.SDL_SCANCODE_SYSREQ     // "SysReq" (the SysReq key)
	SCANCODE_CANCEL     Scancode = C.SDL_SCANCODE_CANCEL     // "Cancel"
	SCANCODE_CLEAR      Scancode = C.SDL_SCANCODE_CLEAR      // "Clear"
	SCANCODE_PRIOR      Scancode = C.SDL_SCANCODE_PRIOR      // "Prior"
	SCANCODE_RETURN2    Scancode = C.SDL_SCANCODE_RETURN2    // "Return"
	SCANCODE_SEPARATOR  Scancode = C.SDL_SCANCODE_SEPARATOR  // "Separator"
	SCANCODE_OUT        Scancode = C.SDL_SCANCODE_OUT        // "Out"
	SCANCODE_OPER       Scancode = C.SDL_SCANCODE_OPER       // "Oper"
	SCANCODE_CLEARAGAIN Scancode = C.SDL_SCANCODE_CLEARAGAIN // "Clear / Again"
	SCANCODE_CRSEL      Scancode = C.SDL_SCANCODE_CRSEL      // "CrSel"
	SCANCODE_EXSEL      Scancode = C.SDL_SCANCODE_EXSEL      // "ExSel"

	SCANCODE_KP_00              Scancode = C.SDL_SCANCODE_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	SCANCODE_KP_000             Scancode = C.SDL_SCANCODE_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	SCANCODE_THOUSANDSSEPARATOR Scancode = C.SDL_SCANCODE_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	SCANCODE_DECIMALSEPARATOR   Scancode = C.SDL_SCANCODE_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	SCANCODE_CURRENCYUNIT       Scancode = C.SDL_SCANCODE_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	SCANCODE_CURRENCYSUBUNIT    Scancode = C.SDL_SCANCODE_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	SCANCODE_KP_LEFTPAREN       Scancode = C.SDL_SCANCODE_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	SCANCODE_KP_RIGHTPAREN      Scancode = C.SDL_SCANCODE_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	SCANCODE_KP_LEFTBRACE       Scancode = C.SDL_SCANCODE_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	SCANCODE_KP_RIGHTBRACE      Scancode = C.SDL_SCANCODE_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	SCANCODE_KP_TAB             Scancode = C.SDL_SCANCODE_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	SCANCODE_KP_BACKSPACE       Scancode = C.SDL_SCANCODE_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	SCANCODE_KP_A               Scancode = C.SDL_SCANCODE_KP_A               // "Keypad A" (the A key (numeric keypad))
	SCANCODE_KP_B               Scancode = C.SDL_SCANCODE_KP_B               // "Keypad B" (the B key (numeric keypad))
	SCANCODE_KP_C               Scancode = C.SDL_SCANCODE_KP_C               // "Keypad C" (the C key (numeric keypad))
	SCANCODE_KP_D               Scancode = C.SDL_SCANCODE_KP_D               // "Keypad D" (the D key (numeric keypad))
	SCANCODE_KP_E               Scancode = C.SDL_SCANCODE_KP_E               // "Keypad E" (the E key (numeric keypad))
	SCANCODE_KP_F               Scancode = C.SDL_SCANCODE_KP_F               // "Keypad F" (the F key (numeric keypad))
	SCANCODE_KP_XOR             Scancode = C.SDL_SCANCODE_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	SCANCODE_KP_POWER           Scancode = C.SDL_SCANCODE_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	SCANCODE_KP_PERCENT         Scancode = C.SDL_SCANCODE_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	SCANCODE_KP_LESS            Scancode = C.SDL_SCANCODE_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	SCANCODE_KP_GREATER         Scancode = C.SDL_SCANCODE_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	SCANCODE_KP_AMPERSAND       Scancode = C.SDL_SCANCODE_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	SCANCODE_KP_DBLAMPERSAND    Scancode = C.SDL_SCANCODE_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	SCANCODE_KP_VERTICALBAR     Scancode = C.SDL_SCANCODE_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	SCANCODE_KP_DBLVERTICALBAR  Scancode = C.SDL_SCANCODE_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	SCANCODE_KP_COLON           Scancode = C.SDL_SCANCODE_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	SCANCODE_KP_HASH            Scancode = C.SDL_SCANCODE_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	SCANCODE_KP_SPACE           Scancode = C.SDL_SCANCODE_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	SCANCODE_KP_AT              Scancode = C.SDL_SCANCODE_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	SCANCODE_KP_EXCLAM          Scancode = C.SDL_SCANCODE_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	SCANCODE_KP_MEMSTORE        Scancode = C.SDL_SCANCODE_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	SCANCODE_KP_MEMRECALL       Scancode = C.SDL_SCANCODE_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	SCANCODE_KP_MEMCLEAR        Scancode = C.SDL_SCANCODE_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	SCANCODE_KP_MEMADD          Scancode = C.SDL_SCANCODE_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	SCANCODE_KP_MEMSUBTRACT     Scancode = C.SDL_SCANCODE_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	SCANCODE_KP_MEMMULTIPLY     Scancode = C.SDL_SCANCODE_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	SCANCODE_KP_MEMDIVIDE       Scancode = C.SDL_SCANCODE_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	SCANCODE_KP_PLUSMINUS       Scancode = C.SDL_SCANCODE_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	SCANCODE_KP_CLEAR           Scancode = C.SDL_SCANCODE_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	SCANCODE_KP_CLEARENTRY      Scancode = C.SDL_SCANCODE_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	SCANCODE_KP_BINARY          Scancode = C.SDL_SCANCODE_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	SCANCODE_KP_OCTAL           Scancode = C.SDL_SCANCODE_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	SCANCODE_KP_DECIMAL         Scancode = C.SDL_SCANCODE_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	SCANCODE_KP_HEXADECIMAL     Scancode = C.SDL_SCANCODE_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	SCANCODE_LCTRL          Scancode = C.SDL_SCANCODE_LCTRL          // "Left Ctrl"
	SCANCODE_LSHIFT         Scancode = C.SDL_SCANCODE_LSHIFT         // "Left Shift"
	SCANCODE_LALT           Scancode = C.SDL_SCANCODE_LALT           // "Left Alt" (alt, option)
	SCANCODE_LGUI           Scancode = C.SDL_SCANCODE_LGUI           // "Left GUI" (windows, command (apple), meta)
	SCANCODE_RCTRL          Scancode = C.SDL_SCANCODE_RCTRL          // "Right Ctrl"
	SCANCODE_RSHIFT         Scancode = C.SDL_SCANCODE_RSHIFT         // "Right Shift"
	SCANCODE_RALT           Scancode = C.SDL_SCANCODE_RALT           // "Right Alt" (alt gr, option)
	SCANCODE_RGUI           Scancode = C.SDL_SCANCODE_RGUI           // "Right GUI" (windows, command (apple), meta)
	SCANCODE_MODE           Scancode = C.SDL_SCANCODE_MODE           // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	SCANCODE_AUDIONEXT      Scancode = C.SDL_SCANCODE_AUDIONEXT      // "AudioNext" (the Next Track media key)
	SCANCODE_AUDIOPREV      Scancode = C.SDL_SCANCODE_AUDIOPREV      // "AudioPrev" (the Previous Track media key)
	SCANCODE_AUDIOSTOP      Scancode = C.SDL_SCANCODE_AUDIOSTOP      // "AudioStop" (the Stop media key)
	SCANCODE_AUDIOPLAY      Scancode = C.SDL_SCANCODE_AUDIOPLAY      // "AudioPlay" (the Play media key)
	SCANCODE_AUDIOMUTE      Scancode = C.SDL_SCANCODE_AUDIOMUTE      // "AudioMute" (the Mute volume key)
	SCANCODE_MEDIASELECT    Scancode = C.SDL_SCANCODE_MEDIASELECT    // "MediaSelect" (the Media Select key)
	SCANCODE_WWW            Scancode = C.SDL_SCANCODE_WWW            // "WWW" (the WWW/World Wide Web key)
	SCANCODE_MAIL           Scancode = C.SDL_SCANCODE_MAIL           // "Mail" (the Mail/eMail key)
	SCANCODE_CALCULATOR     Scancode = C.SDL_SCANCODE_CALCULATOR     // "Calculator" (the Calculator key)
	SCANCODE_COMPUTER       Scancode = C.SDL_SCANCODE_COMPUTER       // "Computer" (the My Computer key)
	SCANCODE_AC_SEARCH      Scancode = C.SDL_SCANCODE_AC_SEARCH      // "AC Search" (the Search key (application control keypad))
	SCANCODE_AC_HOME        Scancode = C.SDL_SCANCODE_AC_HOME        // "AC Home" (the Home key (application control keypad))
	SCANCODE_AC_BACK        Scancode = C.SDL_SCANCODE_AC_BACK        // "AC Back" (the Back key (application control keypad))
	SCANCODE_AC_FORWARD     Scancode = C.SDL_SCANCODE_AC_FORWARD     // "AC Forward" (the Forward key (application control keypad))
	SCANCODE_AC_STOP        Scancode = C.SDL_SCANCODE_AC_STOP        // "AC Stop" (the Stop key (application control keypad))
	SCANCODE_AC_REFRESH     Scancode = C.SDL_SCANCODE_AC_REFRESH     // "AC Refresh" (the Refresh key (application control keypad))
	SCANCODE_AC_BOOKMARKS   Scancode = C.SDL_SCANCODE_AC_BOOKMARKS   // "AC Bookmarks" (the Bookmarks key (application control keypad))
	SCANCODE_BRIGHTNESSDOWN Scancode = C.SDL_SCANCODE_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	SCANCODE_BRIGHTNESSUP   Scancode = C.SDL_SCANCODE_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	SCANCODE_DISPLAYSWITCH  Scancode = C.SDL_SCANCODE_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	SCANCODE_KBDILLUMTOGGLE Scancode = C.SDL_SCANCODE_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	SCANCODE_KBDILLUMDOWN   Scancode = C.SDL_SCANCODE_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	SCANCODE_KBDILLUMUP     Scancode = C.SDL_SCANCODE_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	SCANCODE_EJECT          Scancode = C.SDL_SCANCODE_EJECT          // "Eject" (the Eject key)
	SCANCODE_SLEEP          Scancode = C.SDL_SCANCODE_SLEEP          // "Sleep" (the Sleep key)
	SCANCODE_APP1           Scancode = C.SDL_SCANCODE_APP1
	SCANCODE_APP2           Scancode = C.SDL_SCANCODE_APP2
	NUM_SCANCODES           Scancode = C.SDL_NUM_SCANCODES
)

func (code Scancode) c() C.SDL_Scancode {
	return C.SDL_Scancode(code)
}
