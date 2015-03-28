go-toml-config [![GoDoc](https://godoc.org/gopkg.in/go-kornel/go-toml-config.v0?status.svg)](https://godoc.org/gopkg.in/go-kornel/go-toml-config.v0) [![GoCover](http://gocover.io/_badge/gopkg.in/go-kornel/go-toml-config.v0)](http://gocover.io/gopkg.in/go-kornel/go-toml-config.v0)
==============


go-toml-config is a simple [TOML](https://github.com/mojombo/toml)-based
configuration package for Golang apps that allows you to easily load
configuration files and set defaults. It's a simple wrapper around
[`flag.FlagSet`](http://golang.org/pkg/flag/), so you can use it in pretty much
the same exact way.

[API documentation](http://godoc.org/gopkg.in/go-kornel/go-toml-config.v0)

Example
--------

With `my_app.conf`:

```toml
country = "USA"

[atlanta]
enabled = true
population = 432427
temperature = 99.6
```

Use:

```go
import "github.com/stvp/go-toml-config"

var (
  country            = config.String("country", "Unknown")
  atlantaEnabled     = config.Bool("atlanta.enabled", false)
  alantaPopulation   = config.Int("atlanta.population", 0)
  atlantaTemperature = config.Float("atlanta.temperature", 0)
)

func main() {
  config.Parse("/path/to/my_app.conf")
}
```

You can also create different ConfigSets to manage different logical groupings
of config variables:

```go
networkConfig = config.New("network settings", config.ExitOnError)
networkConfig.String("host", "localhost")
networkConfig.Int("port", 8080)
networkConfig.Parse("/path/to/network.conf")
```

About
-----
This package has been forked from https://github.com/stvp/go-toml-config.

* Some stutter has been avoided.
* New ParseString functions have beed added that enable you to e.g. use a couple
  of different sources for your configuration.
* New PareArguments functions have been added that enable you to get your
  configuration from command line arguments.
* Added PrintCurrentValues functions.

License
-------

Copyright (c) 2015 Kornel Maczyński.
Copyright (c) 2013 Stovepipe Studios, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

