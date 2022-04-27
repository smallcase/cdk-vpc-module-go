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
			_jsii_.MemberProperty{JsiiProperty: "natProvider", GoGetter: "NatProvider"},
			_jsii_.MemberProperty{JsiiProperty: "natSubnets", GoGetter: "NatSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pbSubnets", GoGetter: "PbSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "pvSubnets", GoGetter: "PvSubnets"},
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
		"@smallcase/cdk-vpc-module.VPCProps",
		reflect.TypeOf((*VPCProps)(nil)).Elem(),
	)
}
