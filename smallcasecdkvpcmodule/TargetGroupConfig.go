package smallcasecdkvpcmodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

type TargetGroupConfig struct {
	ApplicationPort *float64 `field:"required" json:"applicationPort" yaml:"applicationPort"`
	Host *string `field:"required" json:"host" yaml:"host"`
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	HealthCheckPort *float64 `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	HealthCheckProtocol awselasticloadbalancingv2.Protocol `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	Protocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

