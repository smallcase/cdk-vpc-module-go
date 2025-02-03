package smallcasecdkvpcmodule


type NetworkLoadBalancerConfig struct {
	SecurityGroupRules *[]*SecurityGroupRule `field:"required" json:"securityGroupRules" yaml:"securityGroupRules"`
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	Certificates *[]*string `field:"optional" json:"certificates" yaml:"certificates"`
	ExistingSecurityGroupId *string `field:"optional" json:"existingSecurityGroupId" yaml:"existingSecurityGroupId"`
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
}

