package model

type Command struct {
	command     string
	description string
	subcommands []string
}

var CommandList = []*Command{
	{
		command:     "status",
		description: "Show desktop capability status.",
	},
	{
		command:     "doctor",
		description: "Run desktop environment checks",
	},
}
