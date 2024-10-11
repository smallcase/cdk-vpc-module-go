package smallcasecdkvpcmodule


type PeeringConfig struct {
	PeeringVpcId *string `field:"required" json:"peeringVpcId" yaml:"peeringVpcId"`
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
	PeerAssumeRoleArn *string `field:"optional" json:"peerAssumeRoleArn" yaml:"peerAssumeRoleArn"`
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
}

