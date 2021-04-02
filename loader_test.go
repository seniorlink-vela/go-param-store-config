package psconfig_test

import (
	"flag"
	"testing"
	"time"

	//"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	psconfig "github.com/seniorlink-vela/go-param-store-config"
)

var runIntegration bool

func init() {
	flag.BoolVar(&runIntegration, "integration", false, "run integration tests")
}

type LoaderSuite struct {
	suite.Suite
	l psconfig.Loader
	c interface{}
}

type mockParamStore struct {
	ssmiface.SSMAPI
	sourceData     map[string]string
	listParameters map[string]bool
	err            error
}

func (m *mockParamStore) GetParametersByPathPages(input *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error {
	if m.err != nil {
		return m.err
	}
	out := &ssm.GetParametersByPathOutput{
		Parameters: []*ssm.Parameter{},
	}
	for k, v := range m.sourceData {
		p := &ssm.Parameter{
			Name:  aws.String(k),
			Value: aws.String(v),
			Type:  aws.String(ssm.DocumentParameterTypeString),
		}
		if _, ok := m.listParameters[k]; ok {
			p.SetType(ssm.DocumentParameterTypeStringList)
		}
		out.Parameters = append(out.Parameters, p)
	}
	fn(out, true)
	return nil
}

var pm = map[string]string{
	"/env/application/http/port":                   "8085",
	"/env/application/http/profiling_port":         "6065",
	"/env/application/http/read_timeout":           "5s",
	"/env/application/http/write_timeout":          "2m",
	"/env/application/log/log_level":               "info",
	"/env/application/log/output_paths":            "stdout,stderr",
	"/env/application/caching/base_uri":            "cache.dev:6379",
	"/env/application/caching/pool_size":           "25",
	"/env/application/service_login/username":      "user-name",
	"/env/application/service_login/password":      "P@ssword!",
	"/env/application/days_valid":                  "720h",
	"/env/application/code_timeout":                "10m",
	"/env/application/api_base_uri":                "https://example.com/api/admin/v1",
	"/env/application/level1/level2/level3/value1": "one",
	"/env/application/level1/level2/level3/value2": "two",
	"/env/application/level1/level2/value":         "foo",
	"/env/application/level1/value1":               "one",
	"/env/application/level1/value2":               "two",
}

func (s *LoaderSuite) SetupSuite() {
	s.l = psconfig.Loader{
		SSM: &mockParamStore{
			sourceData: pm,
			listParameters: map[string]bool{
				"/env/application/log/output_paths": true,
			},
		},
	}
	var config struct {
		HTTP struct {
			Port          int           `ps:"port"`
			ProfilingPort int           `ps:"profiling_port"`
			ReadTimeout   time.Duration `ps:"read_timeout"`
			WriteTimeout  time.Duration `ps:"write_timeout"`
		} `ps:"http"`
		Log struct {
			LogLevel    string   `ps:"log_level"`
			OutputPaths []string `ps:"output_paths"`
		} `ps:"log"`
		Caching struct {
			BaseURI  string `ps:"base_uri"`
			PoolSize int    `ps:"pool_size"`
		} `ps:"caching"`
		ServiceLogin struct {
			Username string `ps:"username"`
			Password string `ps:"password"`
		} `ps:"service_login"`
		DaysValid   time.Duration `ps:"days_valid"`
		CodeTimeout time.Duration `ps:"code_timeout"`
		ApiBaseURI  string        `ps:"api_base_uri"`
		Level1      struct {
			Level2 struct {
				Level3 struct {
					Value1 string `ps:"value1"`
					Value2 string `ps:"value2"`
				} `ps:"level3"`
				Value string `ps:"value"`
			} `ps:"level2"`
			Value1 string `ps:"value1"`
			Value2 string `ps:"value2"`
		} `ps:"level1"`
	}
	s.c = &config
}

func TestInit(t *testing.T) {
	suite.Run(t, new(LoaderSuite))
}

func (s *LoaderSuite) TestLoadSuccess() {
	require.NoError(s.T(), s.l.Load("/env/application/", s.c))
}

func (s *LoaderSuite) TestLoadIntegrationSuccess() {
	if !runIntegration {
		s.T().Skip("Do not run integration tests unless explicitly asked")
	}
	require.NoError(s.T(), psconfig.Load("us-east-1", "/env/application/", s.c))
}
