package smallcasecdkvpcmodule


type VpcEndpontServiceConfig struct {
	Alb *LoadBalancerConfig `field:"required" json:"alb" yaml:"alb"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Nlb *NetworkLoadBalancerConfig `field:"required" json:"nlb" yaml:"nlb"`
	AcceptanceRequired *bool `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	AdditionalTags *map[string]*string `field:"optional" json:"additionalTags" yaml:"additionalTags"`
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
}

