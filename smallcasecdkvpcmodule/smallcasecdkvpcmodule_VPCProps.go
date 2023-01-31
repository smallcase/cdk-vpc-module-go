// @smallcase/cdk-vpc-module
package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type VPCProps struct {
	Subnets *[]ISubnetsProps `field:"required" json:"subnets" yaml:"subnets"`
	Vpc *awsec2.VpcProps `field:"required" json:"vpc" yaml:"vpc"`
	PeeringConfigs *map[string]*PeeringConfig `field:"optional" json:"peeringConfigs" yaml:"peeringConfigs"`
}

