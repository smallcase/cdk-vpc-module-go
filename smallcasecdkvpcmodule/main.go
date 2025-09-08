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
			_jsii_.MemberProperty{JsiiProperty: "useNestedStacks", GoGetter: "UseNestedStacks"},
			_jsii_.MemberProperty{JsiiProperty: "useSubnetForNAT", GoGetter: "UseSubnetForNAT"},
		},
		func() interface{} {
			return &jsiiProxy_ISubnetsProps{}
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.LoadBalancerConfig",
		reflect.TypeOf((*LoadBalancerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-vpc-module.Network",
		reflect.TypeOf((*Network)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createSubnet", GoMethod: "CreateSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "endpointOutputs", GoGetter: "EndpointOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "mergeSubnetsByGroupNames", GoMethod: "MergeSubnetsByGroupNames"},
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
		"@smallcase/cdk-vpc-module.NetworkLoadBalancerConfig",
		reflect.TypeOf((*NetworkLoadBalancerConfig)(nil)).Elem(),
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
		"@smallcase/cdk-vpc-module.TargetGroupConfig",
		reflect.TypeOf((*TargetGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VPCProps",
		reflect.TypeOf((*VPCProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VpcEndpointConfig",
		reflect.TypeOf((*VpcEndpointConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-vpc-module.VpcEndpointServiceNestedStack",
		reflect.TypeOf((*VpcEndpointServiceNestedStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportStringListValue", GoMethod: "ExportStringListValue"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "setParameter", GoMethod: "SetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYamlString", GoMethod: "ToYamlString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_VpcEndpointServiceNestedStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkNestedStack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VpcEndpointServiceNestedStackProps",
		reflect.TypeOf((*VpcEndpointServiceNestedStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-vpc-module.VpcEndpontServiceConfig",
		reflect.TypeOf((*VpcEndpontServiceConfig)(nil)).Elem(),
	)
}
