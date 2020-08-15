{ lib, ... }:
let inherit (lib) literalExample mkOption nameValuePair types;
in {
  host = mkOption {
    default = "localhost";
    type = types.str;
    description = ''
      Database host
    '';
  };

  port = mkOption {
    type = types.int;
    description = "Database port.";
    default = 5432;
    defaultText = "5432";
  };

  name = mkOption {
    type = types.str;
    default = "auth";
    description = "Database name.";
  };

  user = mkOption {
    type = types.str;
    default = "auth";
    description = "Database user.";
  };

  password = mkOption {
    type = types.str;
    default = "auth";
    example = "password";
    description = ''
      Password field.
    '';
  };
}
