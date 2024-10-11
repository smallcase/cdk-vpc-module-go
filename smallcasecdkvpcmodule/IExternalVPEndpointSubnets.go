package smallcasecdkvpcmodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IExternalVPEndpointSubnets interface {
	AvailabilityZone() *string
	Id() *string
	RouteTableId() *string
}

// The jsii proxy for IExternalVPEndpointSubnets
type jsiiProxy_IExternalVPEndpointSubnets struct {
	_ byte // padding
}

func (j *jsiiProxy_IExternalVPEndpointSubnets) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExternalVPEndpointSubnets) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExternalVPEndpointSubnets) RouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableId",
		&returns,
	)
	return returns
}

