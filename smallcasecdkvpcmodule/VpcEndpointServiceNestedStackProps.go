package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type VpcEndpointServiceNestedStackProps struct {
	// A description of the stack.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Simple Notification Service (SNS) topics to publish stack related events.
	// Default: - notifications are not sent for this stack.
	//
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding
	// to a parameter defined in the embedded template and a value representing
	// the value that you want to set for the parameter.
	//
	// The nested stack construct will automatically synthesize parameters in order
	// to bind references from the parent stack(s) into the nested stack.
	// Default: - no user-defined parameters are passed to the nested stack.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Policy to apply when the nested stack is removed.
	//
	// The default is `Destroy`, because all Removal Policies of resources inside the
	// Nested Stack should already have been set correctly. You normally should
	// not need to set this value.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The length of time that CloudFormation waits for the nested stack to reach the CREATE_COMPLETE state.
	//
	// When CloudFormation detects that the nested stack has reached the
	// CREATE_COMPLETE state, it marks the nested stack resource as
	// CREATE_COMPLETE in the parent stack and resumes creating the parent stack.
	// If the timeout period expires before the nested stack reaches
	// CREATE_COMPLETE, CloudFormation marks the nested stack as failed and rolls
	// back both the nested stack and parent stack.
	// Default: - no timeout.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	Subnets *map[string]*[]awsec2.Subnet `field:"required" json:"subnets" yaml:"subnets"`
	Vpc awsec2.Vpc `field:"required" json:"vpc" yaml:"vpc"`
	VpcEndpointServiceConfigs *[]*VpcEndpontServiceConfig `field:"required" json:"vpcEndpointServiceConfigs" yaml:"vpcEndpointServiceConfigs"`
}

