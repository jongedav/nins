package cli

import (
	"errors"
	"strings"

	"github.com/jongedav/nins/internal/model"
)

func ParseArgs(args []string) (*model.Request, error) {
	request := new(model.Request)

	for i := 1; i < len(args); i++ {

		switch {
		case strings.HasPrefix(args[i], "--"):
			flag, err := findGlobalFlagByName(args[i][2:])
			if err != nil {
				return nil, err
			}
			if flag.NeedsValue {
				if i+1 == len(args) {
					return request, errors.New("too few values given for flag.")
				}
				if err := flag.Apply(request, args[i+1]); err != nil {
					return request, err
				}
				i++

			} else {
				if err := flag.Apply(request, ""); err != nil {
					return request, err
				}
			}

		case strings.HasPrefix(args[i], "-") && len(args[i]) > 1:
			for _, arg := range args[i][1:] {
				flag, err := findGlobalFlagByShortName(string(arg))
				if err != nil {
					return request, err
				}
				if flag.NeedsValue {
					return request, errors.New("short flag cannot take values.")
				}
				if err := flag.Apply(request, ""); err != nil {
					return request, err
				}
			}
		default:

		}

	}
	return request, nil
}

func findGlobalFlagByName(name string) (*model.GlobalFlag, error) {
	for _, flag := range model.GlobalFlagList {
		if name == flag.Name {
			return flag, nil
		}
	}
	return nil, errors.New("unknown flag given.")
}

func findGlobalFlagByShortName(name string) (*model.GlobalFlag, error) {
	for _, flag := range model.GlobalFlagList {
		if name == flag.ShortName {
			return flag, nil
		}
	}
	return nil, errors.New("unknown flag given.")
}

func isValidCommand(command string, previous_command string) (bool, error) {

}
