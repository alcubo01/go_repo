package module_golors

import (
	"fmt"
)

func init() {
	fmt.Println("Módulo golors activado ...")
}

func Bold(s string) string {
	return fmt.Sprintf("\033[22m%s\033[0m", s)
}

func Italic(s string) string {
	return fmt.Sprintf("\033[23m%s\033[0m", s)
}

func Rgb(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColoredText(color string, s string) string {
	return fmt.Sprintf("\033[3%s%s\033[0m", color, s)
}
