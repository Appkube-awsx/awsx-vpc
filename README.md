What is awsx-vpc
How to write plugin subcommand
How to build / Test
what it does
command input
command output
How to run
awsx-vpc
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

alt text

How to write plugin subcommand
Please refer to the instaruction - https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug / publish and integrate into the main commmand.

How to build / Test

        - With Vault URL    
          go run .\main.go --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx

        - With Access/Secret key
          go run .\main.go --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########
        
        Another way of testing is by running go install command
        go install
        - go install command creates an exe with the name of the module (e.g. awsx-vpc) and save it in the GOPATH
        - Now we can execute this command on command prompt as below
        - With Vault URL    
        awsx-vpc --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/111111111 --vaultToken=xxxxxxxxxx
        - With Access/Secret key
        awsx-vpc --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########
what it does
This subcommand implement the following functionalities

getVpcList - It will get the list of VPCs for a given AWS account id
command input
valutURL = URL location of vault - that stores credentials to call API
acountId = The AWS account id.
zone = AWS region
vaultToken = Vault token to access vault
command output
Vpcs: [
{
  CidrBlock: "10.0.0.0/16",
  CidrBlockAssociationSet: [{
  AssociationId: "#####################",
  CidrBlock: "10.0.0.0/16",
  CidrBlockState: {
  State: "associated"
  }
  }],
  DhcpOptionsId: "###################",
  InstanceTenancy: "default",
  IsDefault: false,
  OwnerId: "#####################",
  State: "available",
  Tags: [{
  Key: "#######################",
  Value: "shared"
  },{
  Key: "Name",
  Value: "##########################"
  }],
  VpcId: "#############################"
}]

How to run
From main awsx command , it is called as follows:

awsx getVpcList --zone=us-east-1 --vaultUrl=http://<server-address>:/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx

OR

awsx getVpcList --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########

If you build it locally , you can simply run it as standalone command as

awsx-vpc --zone=us-east-1 --vaultUrl=http://<server-address>:/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx

OR

awsx-vpc --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########