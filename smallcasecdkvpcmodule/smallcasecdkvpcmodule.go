// @smallcase/cdk-vpc-module
package smallcasecdkvpcmodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-vpc-module-go/smallcasecdkvpcmodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-vpc-module-go/smallcasecdkvpcmodule/internal"
)

type AddRouteOptions struct {
	// What type of router to route this traffic to.
	RouterType awsec2.RouterType `json:"routerType" yaml:"routerType"`
	// IPv4 range this route applies to.
	DestinationCidrBlock *string `json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// IPv6 range this route applies to.
	DestinationIpv6CidrBlock *string `json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Whether this route will enable internet connectivity.
	//
	// If true, this route will be added before any AWS resources that depend
	// on internet connectivity in the VPC will be created.
	EnablesInternetConnectivity *bool `json:"enablesInternetConnectivity" yaml:"enablesInternetConnectivity"`
	ExistingVpcPeeringRouteKey *string `json:"existingVpcPeeringRouteKey" yaml:"existingVpcPeeringRouteKey"`
	RouterId *string `json:"routerId" yaml:"routerId"`
}

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

type Network interface {
	constructs.Construct
	NatProvider() awsec2.NatProvider
	NatSubnets() *[]awsec2.PublicSubnet
	SetNatSubnets(val *[]awsec2.PublicSubnet)
	// The tree node.
	Node() constructs.Node
	PbSubnets() *[]awsec2.PublicSubnet
	SetPbSubnets(val *[]awsec2.PublicSubnet)
	PvSubnets() *[]awsec2.PrivateSubnet
	SetPvSubnets(val *[]awsec2.PrivateSubnet)
	Vpc() awsec2.Vpc
	CreateSubnet(option ISubnetsProps, vpc awsec2.Vpc, peeringConnectionId *PeeringConnectionInternalType) *[]awsec2.Subnet
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Network
type jsiiProxy_Network struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_Network) SetNatSubnets(val *[]awsec2.PublicSubnet) {
	_jsii_.Set(
		j,
		"natSubnets",
		val,
	)
}

func (j *jsiiProxy_Network) SetPbSubnets(val *[]awsec2.PublicSubnet) {
	_jsii_.Set(
		j,
		"pbSubnets",
		val,
	)
}

func (j *jsiiProxy_Network) SetPvSubnets(val *[]awsec2.PrivateSubnet) {
	_jsii_.Set(
		j,
		"pvSubnets",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Network_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-vpc-module.Network",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) CreateSubnet(option ISubnetsProps, vpc awsec2.Vpc, peeringConnectionId *PeeringConnectionInternalType) *[]awsec2.Subnet {
	var returns *[]awsec2.Subnet

	_jsii_.Invoke(
		n,
		"createSubnet",
		[]interface{}{option, vpc, peeringConnectionId},
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

type NetworkACL struct {
	Cidr awsec2.AclCidr `json:"cidr" yaml:"cidr"`
	Traffic awsec2.AclTraffic `json:"traffic" yaml:"traffic"`
}

type PeeringConfig struct {
	PeeringVpcId *string `json:"peeringVpcId" yaml:"peeringVpcId"`
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	PeerAssumeRoleArn *string `json:"peerAssumeRoleArn" yaml:"peerAssumeRoleArn"`
	PeerOwnerId *string `json:"peerOwnerId" yaml:"peerOwnerId"`
	PeerRegion *string `json:"peerRegion" yaml:"peerRegion"`
}

type PeeringConnectionInternalType struct {
}

type VPCProps struct {
	Subnets *[]ISubnetsProps `json:"subnets" yaml:"subnets"`
	Vpc *awsec2.VpcProps `json:"vpc" yaml:"vpc"`
	PeeringConfigs *map[string]*PeeringConfig `json:"peeringConfigs" yaml:"peeringConfigs"`
}

