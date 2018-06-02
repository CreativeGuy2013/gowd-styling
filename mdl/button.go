package bootstrap

import (
	"strings"

	"github.com/dtylman/gowd"
)

const (
	//ButtonColored makes the button the primary color.
	ButtonColored = "mdl-button--colored"

	//ButtonColored makes the button the have a ripple effect.
	ButtonRippled  = "mdl-js-ripple-effect"
)

//NewButton creates a new mdl <button> element
func NewButtonDefault(buttontype string, caption string) *gowd.Element {
	btn := NewElement("button", "btn "+buttontype)
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

func NewFab(icon string, enabled bool, buttontype ...string) *gowd.Element {
	btn := NewElement("button", "mdl-button mdl-js-button mdl-button--fab"+strings.Join(buttontype, ", "))

	if icon != "" {
		btnIcon := NewElement("i", "material-icons")
		btnIcon.SetText(icon)
		btn.AddElement(btnIcon)
	} else {
		return nil
	}
	if !enabled {
		btn.Disable()
	}
	return btn
}

//NewLinkButton creates a new bootstrap link button (<a>)
func NewLinkButton(caption string) *gowd.Element {
	linkBtn := gowd.NewElement("a")
	linkBtn.SetAttribute("href", "#")
	if caption != "" {
		linkBtn.AddElement(gowd.NewText(caption))
	}
	return linkBtn
}
