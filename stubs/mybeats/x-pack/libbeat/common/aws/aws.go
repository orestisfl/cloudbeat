package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"time"
)

type ConfigAWS struct {
	RoleArn            string `config:"role_arn"`
	AssumeRoleDuration time.Duration `config:"assume_role_duration"`
	ExternalID         string `config:"external_id"`
	DefaultRegion      string `config:"default_region"`
}

func InitializeAWSConfig(cfg ConfigAWS) (aws.Config, error) {
	return aws.Config{}, nil
}