package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type AddRouteOptions struct {
	// What type of router to route this traffic to.
	RouterType awsec2.RouterType `field:"required" json:"routerType" yaml:"routerType"`
	// IPv4 range this route applies to.
	// Default: '0.0.0.0/0'
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// IPv6 range this route applies to.
	// Default: - Uses IPv6.
	//
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Whether this route will enable internet connectivity.
	//
	// If true, this route will be added before any AWS resources that depend
	// on internet connectivity in the VPC will be created.
	// Default: false.
	//
	EnablesInternetConnectivity *bool `field:"optional" json:"enablesInternetConnectivity" yaml:"enablesInternetConnectivity"`
	ExistingVpcPeeringRouteKey *string `field:"optional" json:"existingVpcPeeringRouteKey" yaml:"existingVpcPeeringRouteKey"`
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	RouterId *string `field:"optional" json:"routerId" yaml:"routerId"`
}

