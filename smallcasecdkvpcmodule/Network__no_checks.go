//go:build no_runtime_type_checking

package smallcasecdkvpcmodule

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Network) validateCreateSubnetParameters(option ISubnetsProps, vpc awsec2.Vpc) error {
	return nil
}

func validateNetwork_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Network) validateSetNatSubnetsParameters(val *[]awsec2.PublicSubnet) error {
	return nil
}

func (j *jsiiProxy_Network) validateSetPbSubnetsParameters(val *[]awsec2.PublicSubnet) error {
	return nil
}

func (j *jsiiProxy_Network) validateSetPvSubnetsParameters(val *[]awsec2.PrivateSubnet) error {
	return nil
}

func (j *jsiiProxy_Network) validateSetSubnetsParameters(val *map[string]*[]awsec2.Subnet) error {
	return nil
}

func validateNewNetworkParameters(scope constructs.Construct, id *string, props *VPCProps) error {
	return nil
}

