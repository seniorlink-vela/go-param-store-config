aws ssm put-parameter --name "/env/application/http/port" \
	--value "8085" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/profiling_port" \
	--value "6065" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/read_timeout" \
	--value "5s" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/http/write_timeout" \
	--value "2m" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/log/log_level" \
	--value "info" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/log/output_paths" \
	--value "stdout,stdout" \
	--type "StringList" \
	--overwrite
aws ssm put-parameter --name "/env/application/caching/base_uri" \
	--value "cache.dev:6379" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/caching/pool_size" \
	--value "25" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/service_login/username" \
	--value "user-name" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/service_login/password" \
	--value "P@ssword!" \
	--type "SecureString" \
	--overwrite
aws ssm put-parameter --name "/env/application/days_valid" \
	--value "720h" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/code_timeout" \
	--value "10m" \
	--type "String" \
	--overwrite
aws ssm put-parameter --name "/env/application/api_base_uri" \
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
