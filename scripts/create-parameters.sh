aws ssm put-parameter --name "/env/application/http/port" \
	--value "8085" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/profiling-port" \
	--value "6065" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/read-timeout" \
	--value "5s" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/write-timeout" \
	--value "2m" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/log/log-level" \
	--value "info" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/log/output-paths" \
	--value "stdout,stderr" \
	--type "StringList" \
	--overwrite
aws ssm put-parameter --name "/env/application/caching/base-uri" \
	--value "cache.dev:6379" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/caching/pool-size" \
	--value "25" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/service-login/username" \
	--value "user-name" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/service-login/password" \
	--value "P@ssword!" \
	--type "SecureString" \
	--overwrite
aws ssm put-parameter --name "/env/application/days-valid" \
	--value "720h" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/code-timeout" \
	--value "10m" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/api-base-uri" \
	--value "example.com/api/admin/v1" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/level1/level2/level3/value1" \
	--value "one" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/level1/level2/level3/value2" \
	--value "two" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/level1/level2/value" \
	--value "foo" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/level1/value1" \
	--value "one" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/level1/value2" \
	--value "two" \
	--type "String" \
	--overwrite
