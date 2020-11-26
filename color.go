package chalk

import "fmt"

const (
	// 8-color color set terminal (allowed by all terminal)
	black = "\u001b[30m"
	red = "\u001b[31m"
	green = "\u001b[32m"
	yellow = "\u001b[33m"
	blue = "\u001b[34m"
	magenta = "\u001b[35m"
	cyan = "\u001b[36m"
	white = "\u001b[37m"
	Reset = "\u001b[0m"
	// 8-color background color set terminal
	bgBlack = "\u001b[40m"
	bgRed = "\u001b[41m"
	bgGreen = "\u001b[42m"
	bgYellow = "\u001b[43m"
	bgBlue = "\u001b[44m"
	bgMagenta = "\u001b[45m"
	bgCyan = "\u001b[46m"
	bgWhite = "\u001b[47m"
	// Decoration
	Bold = "\u001b[1m"
	Underline = "\u001b[4m"
	Reversed = "\u001b[7m"
)

/*
	Color the string and using the default formats for its operands and returns the resulting string.
	Spaces are added between operands when neither is a string.
*/

func Black(a ...interface{}) string {
	return black + fmt.Sprint(a...) + Reset
}

func Red(a ...interface{}) string {
	return red + fmt.Sprint(a...) + Reset
}

func Green(a ...interface{}) string {
	return green + fmt.Sprint(a...) + Reset
}

func Yellow(a ...interface{}) string {
	return yellow + fmt.Sprint(a...) + Reset
}

func Blue(a ...interface{}) string {
	return blue + fmt.Sprint(a...) + Reset
}

func Magenta(a ...interface{}) string {
	return magenta + fmt.Sprint(a...) + Reset
}

func Cyan(a ...interface{}) string {
	return cyan + fmt.Sprint(a...) + Reset
}

func White(a ...interface{}) string {
	return white + fmt.Sprint(a...) + Reset
}

/*
	Color the string and formats according to a format specifier and returns the resulting string.
*/

func Blackf(format string, a ...interface{}) string {
	return black + fmt.Sprintf(format, a...) + Reset
}
func Redf(format string, a ...interface{}) string {
	return red + fmt.Sprintf(format, a...) + Reset
}
func Greenf(format string, a ...interface{}) string {
	return green + fmt.Sprintf(format, a...) + Reset
}
func Yellowf(format string, a ...interface{}) string {
	return yellow + fmt.Sprintf(format, a...) + Reset
}
func Bluef(format string, a ...interface{}) string {
	return blue + fmt.Sprintf(format, a...) + Reset
}
func Magentaf(format string, a ...interface{}) string {
	return magenta + fmt.Sprintf(format, a...) + Reset
}
func Cyanf(format string, a ...interface{}) string {
	return cyan + fmt.Sprintf(format, a...) + Reset
}
func Whitef(format string, a ...interface{}) string {
	return white + fmt.Sprintf(format, a...) + Reset
}

/*
	Color the background string and using the default formats for its operands and returns the resulting string.
	Spaces are added between operands when neither is a string.
*/

func BgBlack(a ...interface{}) string {
	return bgBlack + fmt.Sprint(a...) + Reset
}
func BgRed(a ...interface{}) string {
	return bgRed + fmt.Sprint(a...) + Reset
}
func BgGreen(a ...interface{}) string {
	return bgGreen + fmt.Sprint(a...) + Reset
}
func BgYellow(a ...interface{}) string {
	return bgYellow + fmt.Sprint(a...) + Reset
}
func BgBlue(a ...interface{}) string {
	return bgBlue + fmt.Sprint(a...) + Reset
}
func BgMagenta(a ...interface{}) string {
	return bgMagenta + fmt.Sprint(a...) + Reset
}
func BgCyan(a ...interface{}) string {
	return bgCyan + fmt.Sprint(a...) + Reset
}
func BgWhite(a ...interface{}) string {
	return bgWhite + fmt.Sprint(a...) + Reset
}

/*
	Color the background string and formats according to a format specifier and returns the resulting string.
*/

func BgBlackf(format string, a ...interface{}) string {
	return bgBlack + fmt.Sprintf(format, a...) + Reset
}
func BgRedf(format string, a ...interface{}) string {
	return bgRed + fmt.Sprintf(format, a...) + Reset
}
func BgGreenf(format string, a ...interface{}) string {
	return bgGreen + fmt.Sprintf(format, a...) + Reset
}
func BgYellowf(format string, a ...interface{}) string {
	return bgYellow + fmt.Sprintf(format, a...) + Reset
}
func BgBluef(format string, a ...interface{}) string {
	return bgBlue + fmt.Sprintf(format, a...) + Reset
}
func BgMagentaf(format string, a ...interface{}) string {
	return bgMagenta + fmt.Sprintf(format, a...) + Reset
}
func BgCyanf(format string, a ...interface{}) string {
	return bgCyan + fmt.Sprintf(format, a...) + Reset
}
func BgWhitef(format string, a ...interface{}) string {
	return bgWhite + fmt.Sprintf(format, a...) + Reset
}



// Bright is supported by 16-color extended color set terminal (may not work on some terminal)
func Bright(color string) string {
	return color[:len(color)-2] + ";1m"
}

// Todo : Actually not available
// Hex is supported by 256-color extended color set terminal (may not work on some terminal)
func Hex(hex string) string {
	//return color[:len(color)-2] + ";5;1m"
	return ""
}

// Todo : Actually not available
// RGB is supported by 256-color extended color set terminal (may not work on some terminal)
func RGB(r, g, b string) string {
	//return color[:len(color)-2] + ";5;1m"
	return ""
}