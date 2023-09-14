package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-vpc/command/vpccmd"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func GetVpcByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*ec2.DescribeVpcsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetVpcByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetVpcByUserCreds(region string, accesskey string, secretKey string, crossAccountRoleArn string, externalId string) (*ec2.DescribeVpcsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accesskey, secretKey, crossAccountRoleArn, externalId)
	return GetVpcByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetVpcByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*ec2.DescribeVpcsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := GetVpc(clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetVpc(clientAuth *client.Auth) (*ec2.DescribeVpcsOutput, error) {
	response, err := vpccmd.GetVpcList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetVpcIds(clientAuth *client.Auth) ([]string, error) {
	response, err := vpccmd.GetVpcIdList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
