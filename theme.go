package pterm

var (
	// ThemeDefault is the default theme used by PTerm.
	// If this variable is overwritten, the new value is used as default theme.
	ThemeDefault = Theme{
		PrimaryStyle:            Style{FgCyan},
		SecondaryStyle:          Style{FgLightMagenta},
		HighlightStyle:          Style{Bold, FgYellow},
		InfoMessageStyle:        Style{FgLightCyan},
		InfoPrefixStyle:         Style{FgBlack, BgCyan},
		SuccessMessageStyle:     Style{FgGreen},
		SuccessPrefixStyle:      Style{FgBlack, BgGreen},
		WarningMessageStyle:     Style{FgYellow},
		WarningPrefixStyle:      Style{FgBlack, BgYellow},
		ErrorMessageStyle:       Style{FgLightRed},
		ErrorPrefixStyle:        Style{FgBlack, BgLightRed},
		FatalMessageStyle:       Style{FgLightRed},
		FatalPrefixStyle:        Style{FgBlack, BgLightRed},
		DescriptionMessageStyle: Style{FgWhite},
		DescriptionPrefixStyle:  Style{FgLightWhite, BgDarkGray},
		ScopeStyle:              Style{FgGray},
		ProgressbarBarStyle:     Style{FgLightCyan},
		ProgressbarTitleStyle:   Style{FgLightCyan},
		HeaderTextStyle:         Style{FgLightWhite, Bold},
		HeaderBackgroundStyle:   Style{BgGray},
		SpinnerStyle:            Style{FgLightCyan},
		SpinnerTextStyle:        Style{FgLightWhite},
		TableStyle:              Style{FgWhite},
		TableHeaderStyle:        Style{FgLightCyan},
		TableSeparatorStyle:     Style{FgGray},
		SectionStyle:            Style{Bold, FgYellow},
		ListTextStyle:           Style{FgWhite},
		ListBulletStyle:         Style{FgGray},
		LetterStyle:             Style{FgDefault},
	}
)

// Theme for PTerm.
// Theme contains every Style used in PTerm. You can create own themes for your application or use one
// of the existing themes.
type Theme struct {
	PrimaryStyle            Style
	SecondaryStyle          Style
	HighlightStyle          Style
	InfoMessageStyle        Style
	InfoPrefixStyle         Style
	SuccessMessageStyle     Style
	SuccessPrefixStyle      Style
	WarningMessageStyle     Style
	WarningPrefixStyle      Style
	ErrorMessageStyle       Style
	ErrorPrefixStyle        Style
	FatalMessageStyle       Style
	FatalPrefixStyle        Style
	DescriptionMessageStyle Style
	DescriptionPrefixStyle  Style
	ScopeStyle              Style
	ProgressbarBarStyle     Style
	ProgressbarTitleStyle   Style
	HeaderTextStyle         Style
	HeaderBackgroundStyle   Style
	SpinnerStyle            Style
	SpinnerTextStyle        Style
	TableStyle              Style
	TableHeaderStyle        Style
	TableSeparatorStyle     Style
	SectionStyle            Style
	ListTextStyle           Style
	ListBulletStyle         Style
	LetterStyle             Style
}

// WithPrimaryStyle returns a new theme with overridden value.
func (t Theme) WithPrimaryStyle(style Style) Theme {
	t.PrimaryStyle = style
	return t
}

// WithSecondaryStyle returns a new theme with overridden value.
func (t Theme) WithSecondaryStyle(style Style) Theme {
	t.SecondaryStyle = style
	return t
}

// WithHighlightStyle returns a new theme with overridden value.
func (t Theme) WithHighlightStyle(style Style) Theme {
	t.HighlightStyle = style
	return t
}

// WithInfoMessageStyle returns a new theme with overridden value.
func (t Theme) WithInfoMessageStyle(style Style) Theme {
	t.InfoMessageStyle = style
	return t
}

// WithInfoPrefixStyle returns a new theme with overridden value.
func (t Theme) WithInfoPrefixStyle(style Style) Theme {
	t.InfoPrefixStyle = style
	return t
}

// WithSuccessMessageStyle returns a new theme with overridden value.
func (t Theme) WithSuccessMessageStyle(style Style) Theme {
	t.SuccessMessageStyle = style
	return t
}

// WithSuccessPrefixStyle returns a new theme with overridden value.
func (t Theme) WithSuccessPrefixStyle(style Style) Theme {
	t.SuccessPrefixStyle = style
	return t
}

// WithWarningMessageStyle returns a new theme with overridden value.
func (t Theme) WithWarningMessageStyle(style Style) Theme {
	t.WarningMessageStyle = style
	return t
}

// WithWarningPrefixStyle returns a new theme with overridden value.
func (t Theme) WithWarningPrefixStyle(style Style) Theme {
	t.WarningPrefixStyle = style
	return t
}

// WithErrorMessageStyle returns a new theme with overridden value.
func (t Theme) WithErrorMessageStyle(style Style) Theme {
	t.ErrorMessageStyle = style
	return t
}

// WithErrorPrefixStyle returns a new theme with overridden value.
func (t Theme) WithErrorPrefixStyle(style Style) Theme {
	t.ErrorPrefixStyle = style
	return t
}

// WithFatalMessageStyle returns a new theme with overridden value.
func (t Theme) WithFatalMessageStyle(style Style) Theme {
	t.FatalMessageStyle = style
	return t
}

// WithFatalPrefixStyle returns a new theme with overridden value.
func (t Theme) WithFatalPrefixStyle(style Style) Theme {
	t.FatalPrefixStyle = style
	return t
}

// WithDescriptionMessageStyle returns a new theme with overridden value.
func (t Theme) WithDescriptionMessageStyle(style Style) Theme {
	t.DescriptionMessageStyle = style
	return t
}

// WithDescriptionPrefixStyle returns a new theme with overridden value.
func (t Theme) WithDescriptionPrefixStyle(style Style) Theme {
	t.DescriptionPrefixStyle = style
	return t
}

// WithListTextStyle returns a new theme with overridden value.
func (t Theme) WithListTextStyle(style Style) Theme {
	t.ListTextStyle = style
	return t
}

// WithListBulletStyle returns a new theme with overridden value.
func (t Theme) WithListBulletStyle(style Style) Theme {
	t.ListBulletStyle = style
	return t
}

// WithLetterStyle returns a new theme with overridden value.
func (t Theme) WithLetterStyle(style Style) Theme {
	t.LetterStyle = style
	return t
}
