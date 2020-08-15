{ config, lib, pkgs, ... }:
let secrets = import ../../../secrets/online-payments/secrets.nix;
in with secrets; {
  imports = [ ./service.nix ];

  services.onlinePayments = {
    enable = true;
    https = true;
    inherit domain email environment;
  };
}
