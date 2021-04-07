// Package psconfig is a utility package, built to load values from
// AWS SSM Param Store into a custom configuration `struct` for your
// application.

package psconfig

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/mitchellh/mapstructure"
)

var (
	// A KindError is thrown when you try to pass an invlaid type to `Load` or `Loader.Load`
	KindError = errors.New("Incorrect config argument. Must be an address to a struct.")
	// A SessionError gets thrown if you can not create an AWS session successfully
	SessionError = errors.New("Could not start AWS session.")
)

// This is an SSM provider.  It allows us to call the AWS parameter store
// API to retrieve the values stored there.
type Loader struct {
	SSM ssmiface.SSMAPI
}

// Load calls the SSM Parameter Store API, and loads the values stored there into the passed
// config.  It will automatically decrypt encrypted values.  config must be a pointer to a struct.
//
// The values are associated with your struct using the `ps` tag.
//
// The prefixPath is the search string used to look up the associated values in the parameter
// store.  For example, passing `/env/application` will find all of the keys and associated values
// nested under that prefix.
func (l *Loader) Load(pathPrefix string, config interface{}) (err error) {
	err = validateConfig(config)
	if err != nil {
		return
	}

	in := &ssm.GetParametersByPathInput{}
	in.SetPath(pathPrefix)
	in.SetWithDecryption(true)
	in.SetRecursive(true)

	pm := make(map[string]interface{})
	err = l.SSM.GetParametersByPathPages(in, func(params *ssm.GetParametersByPathOutput, lastPage bool) bool {
		for _, p := range params.Parameters {
			pt := *p.Type
			if pt == ssm.DocumentParameterTypeStringList {
				val := strings.Split(*p.Value, ",")
				pm[strings.TrimPrefix(*p.Name, pathPrefix)] = val
			} else {
				pm[strings.TrimPrefix(*p.Name, pathPrefix)] = *p.Value
			}
		}
		return !lastPage
	})
	if err != nil {
		return
	}

	cm := map[string]interface{}{}
	for k, v := range pm {
		k = strings.TrimPrefix(k, pathPrefix)
		ks := strings.Split(k, "/")
		if len(ks) == 1 {
			cm[ks[0]] = v
			continue
		}
		if _, ok := cm[ks[0]]; !ok {
			cm[ks[0]] = map[string]interface{}{}
		}
		m := cm[ks[0]].(map[string]interface{})

		var i int
		for i = 1; i < len(ks)-1; i++ {
			if _, ok := m[ks[i]]; !ok {
				m[ks[i]] = map[string]interface{}{}
			}
			m = m[ks[i]].(map[string]interface{})
		}
		m[ks[i]] = v
	}
	decoderConfig := &mapstructure.DecoderConfig{
		DecodeHook:       decodeHooks(),
		WeaklyTypedInput: true,
		Result:           config,
		TagName:          "ps",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return
	}
	err = decoder.Decode(cm)
	return
}

// Load is a simple utility method, that will instantiate an AWS session for you,
// and call Loader.Load with a new SSM provider.  It's a shortcut, so you don't
// have to instantiate the session and ssm service yourself.
func Load(region, pathPrefix string, config interface{}) (err error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		err = SessionError
		return
	}
	svc := ssm.New(sess)
	l := Loader{SSM: svc}
	return l.Load(pathPrefix, config)
}

func decodeHooks() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		var duration time.Duration

		switch t {
		case reflect.TypeOf(duration):
			return time.ParseDuration(data.(string))
		}
		return data, nil
	}
}

func validateConfig(config interface{}) error {
	valConfig := reflect.ValueOf(config)
	if valConfig.Kind() != reflect.Ptr {
		return KindError
	}
	if valConfig.IsNil() || valConfig.Elem().Kind() != reflect.Struct {
		return KindError
	}
	return nil
}
