package vpccmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetVpcList(auth client.Auth) (*ec2.DescribeVpcsOutput, error) {
	log.Println("Getting aws vpc list")
	ec2Client := client.GetClient(auth, client.EC2_CLIENT).(*ec2.EC2)
	result, err := ec2Client.DescribeVpcs(nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(result)
	return result, nil

}

func GetVpcIdList(auth client.Auth) ([]string, error) {
	log.Println("Getting VPC Id list")
	result, err := GetVpcList(auth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var vpclist []string
	for _, vpc := range result.Vpcs {
		fmt.Println(*vpc.VpcId)
		vpclist = append(vpclist, *vpc.VpcId)
	}
	return vpclist, nil
}
