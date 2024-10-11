package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type SecurityGroupRule struct {
	Peer interface{} `field:"required" json:"peer" yaml:"peer"`
	Port awsec2.Port `field:"required" json:"port" yaml:"port"`
	Description *string `field:"optional" json:"description" yaml:"description"`
}

