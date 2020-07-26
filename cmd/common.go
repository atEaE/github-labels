package cmd

// -------------------------------------------------------------------------------
// This file is not a command file.
// It manages structs and functions that are commonly used by all commands.
// -------------------------------------------------------------------------------

// FlagString : manages flag information.(for string)
type FlagString struct {
	Name     string
	Short    string
	DefValue string
	Usage    string
}

// common flags
// user flag
var commonFlagUser = &FlagString{
	Name:     "user",
	Short:    "u",
	DefValue: "",
	Usage:    "set your github account.",
}

// repository
var commonFlagRepo = &FlagString{
	Name:     "repo",
	Short:    "r",
	DefValue: "",
	Usage:    "set your github repository.",
}

// access token
var commonFlagToken = &FlagString{
	Name:     "token",
	Short:    "t",
	DefValue: "",
	Usage:    "set your github access token.",
}
