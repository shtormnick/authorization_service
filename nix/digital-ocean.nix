{
  online-payments = { config, pkgs, ... }: {
    deployment.targetEnv = "digitalOcean";
    deployment.digitalOcean.enableIpv6 = true;
    deployment.digitalOcean.region = "ams3";
    deployment.digitalOcean.size = "s-1vcpu-1gb";
  };
}
