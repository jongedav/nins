package model

type GlobalFlag struct {
	Name        string
	ShortName   string
	Description string
	NeedsValue  bool
	Apply       func(*Request, string) error
}

var GlobalFlagList = []*GlobalFlag{
	{
		Name:        "help",
		ShortName:   "h",
		Description: "Show information.",
		NeedsValue:  false,
		Apply: func(r *Request, _ string) error {
			r.GlobalFlags.Help = true
			return nil
		},
	},
	{
		Name:        "greet",
		Description: "Greet the user using their name.",
		NeedsValue:  true,
		Apply: func(r *Request, name string) error {
			r.GlobalFlags.Greet = name
			return nil
		},
	},
}
