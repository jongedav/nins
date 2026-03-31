package model

type GlobalFlags struct {
	Help  bool
	Greet string
}

type Request struct {
	GlobalFlags GlobalFlags
	Path        []string
	Flags       map[string]string
	Args        []string
}
