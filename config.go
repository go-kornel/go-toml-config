/*
	Package config implements simple TOML-based configuration variables, based on
	the flag package in the standard Go library (In fact, it's just a simple
	wrapper around flag.FlagSet). It is used in a similar manner, minus the usage
	strings and other command-line specific bits.

	Usage:

	Given the following TOML file:

		country = "USA"

		[atlanta]
		enabled = true
		population = 432427
		temperature = 99.6

	Define your config variables and give them defaults:

		import "gopkg.in/go-kornel/go-toml-config.v0"

		var (
			country            = config.String("country", "Unknown")
			atlantaEnabled     = config.Bool("atlanta.enabled", false)
			alantaPopulation   = config.Int("atlanta.population", 0)
			atlantaTemperature = config.Float("atlanta.temperature", 0)
		)

	After all the config variables are defined, load the config file to overwrite
	the default values with the user-supplied config settings:

		if err := config.Parse("/path/to/myconfig.conf"); err != nil {
			panic(err)
		}

	You can also create separate config.Sets for different config files:

		networkConfig = config.New("network settings", config.ExitOnError)
		networkConfig.String("host", "localhost")
		networkConfig.Int("port", 8080)
		networkConfig.Parse("/path/to/network.conf")
*/
package config // import "gopkg.in/go-kornel/go-toml-config.v0"

import (
	"flag"
	"os"
	"time"
)

// flag.ErrorHandling defines how to handle flag parsing errors.
const (
	ContinueOnError flag.ErrorHandling = flag.ContinueOnError
	ExitOnError     flag.ErrorHandling = flag.ExitOnError
	PanicOnError    flag.ErrorHandling = flag.PanicOnError
)

// globalConfig is a package global default config.Set.
var globalConfig = New(os.Args[0], flag.ExitOnError)

// BoolVar defines a bool config with a given name and default value.
// The argument p points to a bool variable in which to store the value of the config.
func BoolVar(p *bool, name string, value bool, usage string) {
	globalConfig.BoolVar(p, name, value, usage)
}

// Bool defines a bool config variable with a given name and default value.
func Bool(name string, value bool, usage string) *bool {
	return globalConfig.Bool(name, value, usage)
}

// IntVar defines a int config with a given name and default value.
// The argument p points to a int variable in which to store the value of the config.
func IntVar(p *int, name string, value int, usage string) {
	globalConfig.IntVar(p, name, value, usage)
}

// Int defines a int config variable with a given name and default value.
func Int(name string, value int, usage string) *int {
	return globalConfig.Int(name, value, usage)
}

// Int64Var defines a int64 config with a given name and default value.
// The argument p points to a int64 variable in which to store the value of the config.
func Int64Var(p *int64, name string, value int64, usage string) {
	globalConfig.Int64Var(p, name, value, usage)
}

// Int64 defines a int64 config variable with a given name and default value.
func Int64(name string, value int64, usage string) *int64 {
	return globalConfig.Int64(name, value, usage)
}

// UintVar defines a uint config with a given name and default value.
// The argument p points to a uint variable in which to store the value of the config.
func UintVar(p *uint, name string, value uint, usage string) {
	globalConfig.UintVar(p, name, value, usage)
}

// Uint defines a uint config variable with a given name and default value.
func Uint(name string, value uint, usage string) *uint {
	return globalConfig.Uint(name, value, usage)
}

// Uint64Var defines a uint64 config with a given name and default value.
// The argument p points to a uint64 variable in which to store the value of the config.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	globalConfig.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 config variable with a given name and default value.
func Uint64(name string, value uint64, usage string) *uint64 {
	return globalConfig.Uint64(name, value, usage)
}

// StringVar defines a string config with a given name and default value.
// The argument p points to a string variable in which to store the value of the config.
func StringVar(p *string, name string, value string, usage string) {
	globalConfig.StringVar(p, name, value, usage)
}

// String defines a string config variable with a given name and default value.
func String(name string, value string, usage string) *string {
	return globalConfig.String(name, value, usage)
}

// Float64Var defines a float64 config with a given name and default value.
// The argument p points to a float64 variable in which to store the value of the config.
func Float64Var(p *float64, name string, value float64, usage string) {
	globalConfig.Float64Var(p, name, value, usage)
}

// Float64 defines a float64 config variable with a given name and default
// value.
func Float64(name string, value float64, usage string) *float64 {
	return globalConfig.Float64(name, value, usage)
}

// DurationVar defines a time.Duration config with a given name and default value.
// The argument p points to a time.Duration variable in which to store the value of the config.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	globalConfig.DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration config variable with a given name and
// default value.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return globalConfig.Duration(name, value, usage)
}

// Parse takes a path to a TOML file and loads it into the global config.Set.
// This must be called after all config flags have been defined but before the
// flags are accessed by the program.
func Parse(path string) error {
	return globalConfig.Parse(path)
}

// ParseString takes a string representing a TOML file and loads it into the
// global config.Set. This must be called after all config flags have been
// defined but before the flags are accessed by the program.
func ParseString(str string) error {
	return globalConfig.ParseString(str)
}

// ParseArgs parses the command-line flags from os.Args[1:]. Must be called after
// all config flags are defined and before flags are accessed by the program.
func ParseArgs() error {
	return globalConfig.ParseArguments(os.Args[1:])
}
