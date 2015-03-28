package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/pelletier/go-toml"
)

// A Set represents a set of defined configure flags. The zero value of a Set
// has no name and has ContinueOnError error handling.
type Set struct {
	*flag.FlagSet
}

// New returns a new config.Set with the given name and error handling
// policy. The three valid error handling policies are: ContinueOnError,
// ExitOnError, and PanicOnError.
func New(name string, errorHandling flag.ErrorHandling) *Set {
	return &Set{
		flag.NewFlagSet(name, errorHandling),
	}
}

// Parse takes a path to a TOML file and loads it. This must be called after
// all the config flags in the config.Set have been defined but before the flags
// are accessed by the program.
func (c *Set) Parse(path string) error {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	tomlTree, err := toml.Load(string(configBytes))
	if err != nil {
		errorString := fmt.Sprintf("%s is not a valid TOML file. See https://github.com/mojombo/toml", path)
		return errors.New(errorString)
	}

	err = c.loadTomlTree(tomlTree, []string{})
	if err != nil {
		return err
	}

	return nil
}

// ParseString takes a string representation of a TOML file and loads it. This
// must be called after all the config flags in the config.Set have been defined
// but before the flags are accessed by the program.
func (c *Set) ParseString(str string) error {
	tomlTree, err := toml.Load(str)
	if err != nil {
		errorString := fmt.Sprintf("Not a valid TOML. See https://github.com/mojombo/toml")
		return errors.New(errorString)
	}

	err = c.loadTomlTree(tomlTree, []string{})
	if err != nil {
		return err
	}

	return nil
}

// ParseArguments parses flag definitions from the argument list, which should
// not include the command name. Must be called after all the config flags in
// the config.Set have been defined but before the flags are accessed by the
// program. The return value will be flag.ErrHelp if -help or -h were set but
// not defined.
func (c *Set) ParseArguments(arguments []string) error {
	return c.FlagSet.Parse(arguments)
}

// PrintCurrentValues prints lines in format
//    flagName=flagCurrentValue
// to the os.Stderr. Useful for showing current configuration to the user.
// The output format is subject to change.
func (c *Set) PrintCurrentValues() {
	c.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "%s=%v\n", f.Name, f.Value.String())
	})
}

// loadTomlTree recursively loads a TomlTree into this config.Set's config
// variables.
func (c *Set) loadTomlTree(tree *toml.TomlTree, path []string) error {
	for _, key := range tree.Keys() {
		fullPath := append(path, key)
		value := tree.Get(key)
		if subtree, isTree := value.(*toml.TomlTree); isTree {
			err := c.loadTomlTree(subtree, fullPath)
			if err != nil {
				return err
			}
		} else {
			fullPath := strings.Join(append(path, key), ".")
			err := c.Set(fullPath, fmt.Sprintf("%v", value))
			if err != nil {
				return buildLoadError(fullPath, err)
			}
		}
	}
	return nil
}

// buildLoadError takes an error from flag.FlagSet#Set and makes it a bit more
// readable, if it recognizes the format.
func buildLoadError(path string, err error) error {
	missingFlag := regexp.MustCompile(`^no such flag -([^\s]+)`)
	invalidSyntax := regexp.MustCompile(`^.+ parsing "(.+)": invalid syntax$`)
	errorString := err.Error()

	if missingFlag.MatchString(errorString) {
		errorString = missingFlag.ReplaceAllString(errorString, "$1 is not a valid config setting")
	} else if invalidSyntax.MatchString(errorString) {
		errorString = "The value for " + path + " is invalid"
	}

	return errors.New(errorString)
}
