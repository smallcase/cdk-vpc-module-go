//go:build !no_runtime_type_checking

package smallcasecdkvpcmodule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_Network) validateCreateSubnetParameters(option ISubnetsProps, vpc awsec2.Vpc) error {
	if option == nil {
		return fmt.Errorf("parameter option is required, but nil was provided")
	}

	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	return nil
}

func validateNetwork_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Network) validateSetNatSubnetsParameters(val *[]awsec2.PublicSubnet) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Network) validateSetPbSubnetsParameters(val *[]awsec2.PublicSubnet) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Network) validateSetPvSubnetsParameters(val *[]awsec2.PrivateSubnet) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Network) validateSetSubnetsParameters(val *map[string]*[]awsec2.Subnet) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkParameters(scope constructs.Construct, id *string, props *VPCProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

