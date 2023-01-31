// @smallcase/cdk-vpc-module
package smallcasecdkvpcmodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type ISubnetsProps interface {
	AvailabilityZones() *[]*string
	CidrBlock() *[]*string
	EgressNetworkACL() *[]*NetworkACL
	IngressNetworkACL() *[]*NetworkACL
	Routes() *[]*AddRouteOptions
	SubnetGroupName() *string
	SubnetType() awsec2.SubnetType
	Tags() *map[string]*string
}

// The jsii proxy for ISubnetsProps
type jsiiProxy_ISubnetsProps struct {
	_ byte // padding
}

func (j *jsiiProxy_ISubnetsProps) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) CidrBlock() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) EgressNetworkACL() *[]*NetworkACL {
	var returns *[]*NetworkACL
	_jsii_.Get(
		j,
		"egressNetworkACL",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) IngressNetworkACL() *[]*NetworkACL {
	var returns *[]*NetworkACL
	_jsii_.Get(
		j,
		"ingressNetworkACL",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) Routes() *[]*AddRouteOptions {
	var returns *[]*AddRouteOptions
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) SubnetType() awsec2.SubnetType {
	var returns awsec2.SubnetType
	_jsii_.Get(
		j,
		"subnetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetsProps) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

