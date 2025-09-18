package smallcasecdkvpcmodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-vpc-module-go/smallcasecdkvpcmodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-vpc-module-go/smallcasecdkvpcmodule/internal"
)

type Network interface {
	constructs.Construct
	EndpointOutputs() *map[string]interface{}
	NatProvider() awsec2.NatProvider
	NatSubnets() *[]awsec2.PublicSubnet
	SetNatSubnets(val *[]awsec2.PublicSubnet)
	// The tree node.
	Node() constructs.Node
	PbSubnets() *[]awsec2.PublicSubnet
	SetPbSubnets(val *[]awsec2.PublicSubnet)
	PvSubnets() *[]awsec2.PrivateSubnet
	SetPvSubnets(val *[]awsec2.PrivateSubnet)
	SecurityGroupOutputs() *map[string]awsec2.SecurityGroup
	Subnets() *map[string]*[]awsec2.Subnet
	SetSubnets(val *map[string]*[]awsec2.Subnet)
	Vpc() awsec2.Vpc
	CreateSubnet(option ISubnetsProps, vpc awsec2.Vpc, peeringConnectionId *PeeringConnectionInternalType, useGlobalNestedStacks *bool) *[]awsec2.Subnet
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Network
type jsiiProxy_Network struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Network) EndpointOutputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"endpointOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NatProvider() awsec2.NatProvider {
	var returns awsec2.NatProvider
	_jsii_.Get(
		j,
		"natProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NatSubnets() *[]awsec2.PublicSubnet {
	var returns *[]awsec2.PublicSubnet
	_jsii_.Get(
		j,
		"natSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) PbSubnets() *[]awsec2.PublicSubnet {
	var returns *[]awsec2.PublicSubnet
	_jsii_.Get(
		j,
		"pbSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) PvSubnets() *[]awsec2.PrivateSubnet {
	var returns *[]awsec2.PrivateSubnet
	_jsii_.Get(
		j,
		"pvSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) SecurityGroupOutputs() *map[string]awsec2.SecurityGroup {
	var returns *map[string]awsec2.SecurityGroup
	_jsii_.Get(
		j,
		"securityGroupOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Subnets() *map[string]*[]awsec2.Subnet {
	var returns *map[string]*[]awsec2.Subnet
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Vpc() awsec2.Vpc {
	var returns awsec2.Vpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewNetwork(scope constructs.Construct, id *string, props *VPCProps) Network {
	_init_.Initialize()

	if err := validateNewNetworkParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Network{}

	_jsii_.Create(
		"@smallcase/cdk-vpc-module.Network",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetwork_Override(n Network, scope constructs.Construct, id *string, props *VPCProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-vpc-module.Network",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_Network)SetNatSubnets(val *[]awsec2.PublicSubnet) {
	if err := j.validateSetNatSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natSubnets",
		val,
	)
}

func (j *jsiiProxy_Network)SetPbSubnets(val *[]awsec2.PublicSubnet) {
	if err := j.validateSetPbSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pbSubnets",
		val,
	)
}

func (j *jsiiProxy_Network)SetPvSubnets(val *[]awsec2.PrivateSubnet) {
	if err := j.validateSetPvSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pvSubnets",
		val,
	)
}

func (j *jsiiProxy_Network)SetSubnets(val *map[string]*[]awsec2.Subnet) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Network_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-vpc-module.Network",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) CreateSubnet(option ISubnetsProps, vpc awsec2.Vpc, peeringConnectionId *PeeringConnectionInternalType, useGlobalNestedStacks *bool) *[]awsec2.Subnet {
	if err := n.validateCreateSubnetParameters(option, vpc, peeringConnectionId); err != nil {
		panic(err)
	}
	var returns *[]awsec2.Subnet

	_jsii_.Invoke(
		n,
		"createSubnet",
		[]interface{}{option, vpc, peeringConnectionId, useGlobalNestedStacks},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

