package models

type AwsConf struct {
	Endpoint         string
	DataFileS3Bucket string `validate:"required"`
	AccountID        string `validate:"required"`
	SQS              SQS    `validate:"required"`
}

type SQS struct {
	ABHAEventsFifo     string `json:"abha_fifo" validate:"required"`
	ABHAEventsStandard string `json:"abha_standard" validate:"required"`
}
