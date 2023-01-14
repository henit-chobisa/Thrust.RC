package UIAssets

import (
	"github.com/briandowns/spinner"
	"time"
)

func Spinner(text string, color string, line string) *spinner.Spinner {
	s := spinner.New([]string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}, 600*time.Millisecond)
	s.Suffix = color + text + line + "\n"
	s.Prefix = "\n"
	return s
}
