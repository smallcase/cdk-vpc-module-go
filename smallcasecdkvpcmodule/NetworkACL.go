package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type NetworkACL struct {
	Cidr awsec2.AclCidr `field:"required" json:"cidr" yaml:"cidr"`
	Traffic awsec2.AclTraffic `field:"required" json:"traffic" yaml:"traffic"`
}

