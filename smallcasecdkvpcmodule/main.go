// @smallcase/cdk-vpc-module
package smallcasecdkvpcmodule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.AddRouteOptions",
		reflect.TypeOf((*AddRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@smallcase/cdk-vpc-module.IExternalVPEndpointSubnets",
		reflect.TypeOf((*IExternalVPEndpointSubnets)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
		},
		func() interface{} {
			return &jsiiProxy_IExternalVPEndpointSubnets{}
		},
	)
	_jsii_.RegisterInterface(
		"@smallcase/cdk-vpc-module.ISubnetsProps",
		reflect.TypeOf((*ISubnetsProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "egressNetworkACL", GoGetter: "EgressNetworkACL"},
			_jsii_.MemberProperty{JsiiProperty: "ingressNetworkACL", GoGetter: "IngressNetworkACL"},
			_jsii_.MemberProperty{JsiiProperty: "routes", GoGetter: "Routes"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroupName", GoGetter: "SubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "subnetType", GoGetter: "SubnetType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "useSubnetForNAT", GoGetter: "UseSubnetForNAT"},
		},
		func() interface{} {
			return &jsiiProxy_ISubnetsProps{}
		},
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-vpc-module.Network",
		reflect.TypeOf((*Network)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createSubnet", GoMethod: "CreateSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "endpointOutputs", GoGetter: "EndpointOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "natProvider", GoGetter: "NatProvider"},
			_jsii_.MemberProperty{JsiiProperty: "natSubnets", GoGetter: "NatSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pbSubnets", GoGetter: "PbSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "pvSubnets", GoGetter: "PvSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupOutputs", GoGetter: "SecurityGroupOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Network{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.NetworkACL",
		reflect.TypeOf((*NetworkACL)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.PeeringConfig",
		reflect.TypeOf((*PeeringConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.PeeringConnectionInternalType",
		reflect.TypeOf((*PeeringConnectionInternalType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.SecurityGroupRule",
		reflect.TypeOf((*SecurityGroupRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VPCProps",
		reflect.TypeOf((*VPCProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VpcEndpointConfig",
		reflect.TypeOf((*VpcEndpointConfig)(nil)).Elem(),
	)
}
