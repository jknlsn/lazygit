package context

import (
	"testing"

	"github.com/gookit/color"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/stretchr/testify/assert"
	"github.com/xo/terminfo"
)

func TestGetBufferLength(t *testing.T) {
	oldColorLevel := color.ForceSetColorLevel(terminfo.ColorLevelMillions)
	defer color.ForceSetColorLevel(oldColorLevel)

	scenarios := []struct {
		name          string
		subject       string
		autoWrapWidth int
		expected      string
	}{
		{
			name:          "under warning width",
			subject:       "hello",
			autoWrapWidth: 72,
			expected:      " 5 ",
		},
		{
			name:          "over warning width (70% of wrap width)",
			subject:       "this subject is definitely longer than fifty characters in total",
			autoWrapWidth: 72,
			expected:      style.FgYellow.SetBold().Sprint(" 64 "),
		},
		{
			name:          "over wrap width",
			subject:       "this subject is so long that it sails past seventy-two characters without any trouble",
			autoWrapWidth: 72,
			expected:      style.FgRed.SetBold().Sprint(" 85 "),
		},
		{
			name:          "thresholds follow the wrap width",
			subject:       "this subject is definitely longer than fifty characters in total",
			autoWrapWidth: 40,
			expected:      style.FgRed.SetBold().Sprint(" 64 "),
		},
		{
			name:          "zero wrap width disables colouring",
			subject:       "any length at all",
			autoWrapWidth: 0,
			expected:      " 17 ",
		},
		{
			name:          "length counts runes, not bytes",
			subject:       "héllo wörld",
			autoWrapWidth: 72,
			expected:      " 11 ",
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			assert.Equal(t, s.expected, getBufferLength(s.subject, s.autoWrapWidth))
		})
	}
}
