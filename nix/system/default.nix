{
  nix.gc = {
    automatic = true;
    dates = "daily";
    options = "--delete-older-than 5d";
  };

  time.timeZone = "Europe/Kiev";

  boot.cleanTmpDir = true;

  security.sudo = {
    enable = true;
    wheelNeedsPassword = false;
  };

  services.clearLogs = { enable = true; };

  imports = [ ./aliases.nix ./programs.nix ./clear_logs/service.nix ];
}
