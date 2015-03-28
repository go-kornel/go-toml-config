package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

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

// BoolVar defines a bool config with a given name and default value for a config.Set.
// The argument p points to a bool variable in which to store the value of the config.
func (c *Set) BoolVar(p *bool, name string, value bool, usage string) {
	c.FlagSet.BoolVar(p, name, value, usage)
}

// Bool defines a bool config variable with a given name and default value for
// a config.Set.
func (c *Set) Bool(name string, value bool, usage string) *bool {
	return c.FlagSet.Bool(name, value, usage)
}

// IntVar defines a int config with a given name and default value for a config.Set.
// The argument p points to a int variable in which to store the value of the config.
func (c *Set) IntVar(p *int, name string, value int, usage string) {
	c.FlagSet.IntVar(p, name, value, usage)
}

// Int defines a int config variable with a given name and default value for a
// config.Set.
func (c *Set) Int(name string, value int, usage string) *int {
	return c.FlagSet.Int(name, value, usage)
}

// Int64Var defines a int64 config with a given name and default value for a config.Set.
// The argument p points to a int64 variable in which to store the value of the config.
func (c *Set) Int64Var(p *int64, name string, value int64, usage string) {
	c.FlagSet.Int64Var(p, name, value, usage)
}

// Int64 defines a int64 config variable with a given name and default value
// for a config.Set.
func (c *Set) Int64(name string, value int64, usage string) *int64 {
	return c.FlagSet.Int64(name, value, usage)
}

// UintVar defines a uint config with a given name and default value for a config.Set.
// The argument p points to a uint variable in which to store the value of the config.
func (c *Set) UintVar(p *uint, name string, value uint, usage string) {
	c.FlagSet.UintVar(p, name, value, usage)
}

// Uint defines a uint config variable with a given name and default value for
// a config.Set.
func (c *Set) Uint(name string, value uint, usage string) *uint {
	return c.FlagSet.Uint(name, value, usage)
}

// Uint64Var defines a uint64 config with a given name and default value for a config.Set.
// The argument p points to a uint64 variable in which to store the value of the config.
func (c *Set) Uint64Var(p *uint64, name string, value uint64, usage string) {
	c.FlagSet.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 config variable with a given name and default value
// for a config.Set.
func (c *Set) Uint64(name string, value uint64, usage string) *uint64 {
	return c.FlagSet.Uint64(name, value, usage)
}

// StringVar defines a string config with a given name and default value for a config.Set.
// The argument p points to a string variable in which to store the value of the config.
func (c *Set) StringVar(p *string, name string, value string, usage string) {
	c.FlagSet.StringVar(p, name, value, usage)
}

// String defines a string config variable with a given name and default value
// for a config.Set.
func (c *Set) String(name string, value string, usage string) *string {
	return c.FlagSet.String(name, value, usage)
}

// Float64Var defines a float64 config with a given name and default value for a config.Set.
// The argument p points to a float64 variable in which to store the value of the config.
func (c *Set) Float64Var(p *float64, name string, value float64, usage string) {
	c.FlagSet.Float64Var(p, name, value, usage)
}

// Float64 defines a float64 config variable with a given name and default
// value for a config.Set.
func (c *Set) Float64(name string, value float64, usage string) *float64 {
	return c.FlagSet.Float64(name, value, usage)
}

// DurationVar defines a time.Duration config with a given name and default value for a config.Set.
// The argument p points to a time.Duration variable in which to store the value of the config.
func (c *Set) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	c.FlagSet.DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration config variable with a given name and
// default value.
func (c *Set) Duration(name string, value time.Duration, usage string) *time.Duration {
	return globalConfig.FlagSet.Duration(name, value, usage)
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
