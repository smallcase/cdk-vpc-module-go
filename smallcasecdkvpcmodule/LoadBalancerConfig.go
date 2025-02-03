package smallcasecdkvpcmodule


type LoadBalancerConfig struct {
	Certificates *[]*string `field:"optional" json:"certificates" yaml:"certificates"`
	ExistingArn *string `field:"optional" json:"existingArn" yaml:"existingArn"`
	ExistingSecurityGroupId *string `field:"optional" json:"existingSecurityGroupId" yaml:"existingSecurityGroupId"`
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	SecurityGroupRules *[]*SecurityGroupRule `field:"optional" json:"securityGroupRules" yaml:"securityGroupRules"`
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	TargetGroups *[]*TargetGroupConfig `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

