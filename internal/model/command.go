package model

type Command struct {
	Command     string
	Description string
	Subcommands []*Command
	Runnable    bool
}

var CommandList = []*Command{
	{
		Command:     "status",
		Description: "Show desktop capability status.",
		Runnable:    true,
	},
	{
		Command:     "doctor",
		Description: "Run desktop environment checks",
		Runnable:    true,
	},
	{
		Command:     "audio",
		Description: "Manage audio settings.",
		Subcommands: []*Command{
			{
				Command:     "input",
				Description: "Manage audio input settings.",
				Subcommands: []*Command{
					{
						Command:     "switch",
						Description: "Switch audio input device.",
						Runnable:    true,
					},
				},
			},
			{
				Command:     "output",
				Description: "Manage audio output settings.",
				Subcommands: []*Command{
					{
						Command:     "switch",
						Description: "Switch audio output device.",
						Runnable:    true,
					},
				},
			},
		},
	},
}

func findCommand(path []string) (*Command, int) {
	currentList := CommandList
	var current *Command
	for i, segment := range path {
		var next *Command
		for _, cmd := range currentList {
			if cmd.Command == segment {
				next = cmd
				break
			}
		}
		if next == nil {
			return current, i
		}
		current = next
		currentList = next.Subcommands
	}
	return current, len(path)
}
