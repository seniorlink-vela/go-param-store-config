# PSConfig [![Go Param Store Config Tests](https://github.com/seniorlink-vela/go-param-store-config/actions/workflows/test-run.yml/badge.svg)](https://github.com/seniorlink-vela/go-param-store-config/actions/workflows/test-run.yml)

PSConfig is a utility library, built to load values from
[AWS SSM Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html)
into a custom configuration `struct` for your application.  It is partially inspired by
[ianlopshire/go-ssm-config](https://github.com/ianlopshire/go-ssm-config), but allows for nested configuration
elements.  In our projects we will typically have a config defined like this:

```HCL
type HTTPConfig struct {
	Port          int           `hcl:"port"`
	ProfilingPort int           `hcl:"profiling_port"`
	ReadTimeout   time.Duration `hcl:"read_timeout"`
	WriteTimeout  time.Duration `hcl:"write_timeout"`
}

type DbConfig struct {
	Host               string        `hcl:"host"`
	Username           string        `hcl:"username"`
	Password           string        `hcl:"password"`
	Name               string        `hcl:"name"`
	Application        string        `hcl:"application"`
}

type ApplicationConfig struct {
    HTTP HTTPConfig `hcl:"http"`
    DB   DBConfig   `hcl:"db"`
}
```

Since we have relatively complicated configuration needs, grouping by the types of configuration
made our lives easier.

## Motivation

We wanted to move our configuration from templated HCL files, into the param store, to make it
easier to manage configuration in general.  Because of our current config structs, we needed to
have something that minimized the amount of application changes we had to make, so we developed
this library to allow us to load arbitrarily nested param store items into an equally arbitrarily
nested config struct.

## Licence

MIT
