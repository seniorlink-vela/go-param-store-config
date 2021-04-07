# PSConfig [![PSConfig Tests](https://github.com/seniorlink-vela/go-param-store-config/actions/workflows/test-run.yml/badge.svg)](https://github.com/seniorlink-vela/go-param-store-config/actions/workflows/test-run.yml)

PSConfig is a utility library, built to load values from
[AWS SSM Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html)
into a custom configuration `struct` for your application.  It is partially inspired by
[ianlopshire/go-ssm-config](https://github.com/ianlopshire/go-ssm-config), but allows for nested configuration
elements.  In our projects we will typically have a config defined like this:

```golang
type HTTPConfig struct {
	Port          int           `hcl:"port"`
	ReadTimeout   time.Duration `hcl:"read_timeout"`
	WriteTimeout  time.Duration `hcl:"write_timeout"`
}

type DbConfig struct {
	Host        string `hcl:"host"`
	Username    string `hcl:"username"`
	Password    string `hcl:"password"`
	Name        string `hcl:"name"`
	Application string `hcl:"application"`
}

type ApplicationConfig struct {
	HTTP HTTPConfig `hcl:"http"`
	DB   DBConfig   `hcl:"db"`
}
```

This would typically be in an HCL file like:

```HCL
http {
  port           = 8080
  read_timeout   = "${duration("5s")}"
  write_timeout  = "${duration("2m")}"
}

db {
  host                  = "database-server:5432"
  username              = "username"
  password              = "passw0rd!"
  application           = "fancy-application-name"
  name                  = "database_name"
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

## Usage

Assume you have already set some parameters in
[AWS SSM Param Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html):

| Name | Value | Type |
|------|-------|------|
| /env/application/http/port | 8080 | String |
| /env/application/http/read_timeout | 5s | String |
| /env/application/http/write_timeout | 8080 | String |
| /env/application/db/host | database-server:5432 | String |
| /env/application/db/username | username | String |
| /env/application/db/password | passw0rd | SecureString |
| /env/application/db/application | fancy-application-name | String |
| /env/application/db/name | database_name | String |

You can then write code like the following:

```golang
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	psconfig "github.com/seniorlink-vela/go-param-store-config"
)

type HTTPConfig struct {
	Port         int           `ps:"port"`
	ReadTimeout  time.Duration `ps:"read_timeout"`
	WriteTimeout time.Duration `ps:"write_timeout"`
}

type DbConfig struct {
	Host        string `ps:"host"`
	Username    string `ps:"username"`
	Password    string `ps:"password"`
	Name        string `ps:"name"`
	Application string `ps:"application"`
}

type ApplicationConfig struct {
	HTTP HTTPConfig `ps:"http"`
	DB   DbConfig   `ps:"db"`
}

func main() {
	c := &ApplicationConfig{}
	err := psconfig.Load("us-east-1", "/env/application/", c)
	if err != nil {
		log.Fatal(err.Error())
	}
	j, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", j)
}
```

The output of this is:

```json
{
  "HTTP": {
    "Port": 8085,
    "ReadTimeout": 5000000000,
    "WriteTimeout": 120000000000
  },
  "DB": {
    "Host": "database-server:5432",
    "Username": "username",
    "Password": "passw0rd",
    "Name": "database_name",
    "Application": "fancy-application-name"
  }
}
```

### Tags

You'll notice in the example above the struct fields are tagged with `ps`.  This is what we use to
identify which parameter store values should be placed in these fields.

### Supported Struct Field Types

psconfig supports these struct field types:

* string
* int, int8, int16, int32, int64
* bool
* float32, float64
* []string
* struct

## Licence

MIT
