let

  region = "us-east-2";
  accessKeyId =
    "magento"; # symbolic name looked up in ~/.ec2-keys or a ~/.aws/credentials profile name

  ec2 = { resources, ... }: {
    deployment.targetEnv = "ec2";
    deployment.ec2.accessKeyId = accessKeyId;
    deployment.ec2.region = region;
    deployment.ec2.ebsInitialRootDiskSize = 10;
    deployment.ec2.instanceType = "t2.small";
    deployment.ec2.keyPair = resources.ec2KeyPairs.magento;
    deployment.ec2.spotInstancePrice = 3;
  };

in {
  online-payments = ec2;

  # Provision an EC2 key pair.
  resources.ec2KeyPairs.magento = { inherit region accessKeyId; };
}
