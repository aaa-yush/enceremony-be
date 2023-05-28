package models

type AwsConf struct {
	Endpoint  string
	AccountID string `validate:"required"`
}
