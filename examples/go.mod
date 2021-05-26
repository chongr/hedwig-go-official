module examples

go 1.16

replace github.com/cloudchacho/hedwig-go v0.0.0 => ../

replace github.com/cloudchacho/hedwig-go/aws v0.0.0 => ../aws

replace github.com/cloudchacho/hedwig-go/jsonschema v0.0.0 => ../jsonschema

replace github.com/cloudchacho/hedwig-go/protobuf v0.0.0 => ../protobuf

replace github.com/cloudchacho/hedwig-go/protobuf/internal v0.0.0 => ../protobuf/internal

require (
	github.com/cloudchacho/hedwig-go v0.0.0
	github.com/cloudchacho/hedwig-go/aws v0.0.0
	github.com/cloudchacho/hedwig-go/jsonschema v0.0.0
	github.com/cloudchacho/hedwig-go/protobuf v0.0.0
	google.golang.org/protobuf v1.26.0
)