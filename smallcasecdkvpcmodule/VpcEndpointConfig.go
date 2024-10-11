package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type VpcEndpointConfig struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Service interface{} `field:"required" json:"service" yaml:"service"`
	SubnetGroupNames *[]*string `field:"required" json:"subnetGroupNames" yaml:"subnetGroupNames"`
	AdditionalTags *map[string]*string `field:"optional" json:"additionalTags" yaml:"additionalTags"`
	ExternalSubnets *[]IExternalVPEndpointSubnets `field:"optional" json:"externalSubnets" yaml:"externalSubnets"`
	IamPolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"iamPolicyStatements" yaml:"iamPolicyStatements"`
	SecurityGroupRules *[]*SecurityGroupRule `field:"optional" json:"securityGroupRules" yaml:"securityGroupRules"`
}

