{
  network = {
    description = "Online payments";
    enableRollback = true;
  };

  resources.sshKeyPairs.ssh-key = { };

  online-payments = { config, lib, pkgs, ... }: {
    imports = [ ./default.nix ./system ];
  };
}
