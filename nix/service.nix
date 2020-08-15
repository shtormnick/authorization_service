{ config, lib, pkgs, ... }:

with lib;

let
  cfg = config.services.onlinePayments;
  is_local_db = (cfg.database.host == "localhost");
  dbUrl =
    "postgres://${cfg.database.user}:${cfg.database.password}@${cfg.database.host}:${
      toString (cfg.database.port)
    }/${cfg.database.name}?sslmode=disable";

  proxyPass = "127.0.0.1:8000";

  afterService = [ "postgresql.service" "redis.service" "rabbitmq.service" ];

  pacageBack = import ../default.nix { };

  migratePkg = import ./migrate { };

  migrationPath = lib.cleanSourceWith {
    filter = lib.cleanSourceFilter;
    src = ../migrations;
  };

  templatesPath = lib.cleanSourceWith {
    filter = lib.cleanSourceFilter;
    src = ../templates;
  };

  swaggerPath = import ./swagger { };

  user = cfg.database.user;
  group = cfg.database.user;
in {

  options = {
    services.onlinePayments = {
      enable = mkOption {
        type = types.bool;
        default = false;
        description = ''
          Enable online payments server.
        '';
      };

      domain = mkOption {
        type = types.str;
        example = "example.com";
        description = ''
          Domain
        '';
      };

      https = mkOption {
        default = false;
        description = ''
          Enable https.
        '';
      };

      environment = mkOption {
        type = types.attrs;
        default = { };
        description = ''
          Add environment to systemd
        '';
      };

      email = mkOption {
        type = types.str;
        example = "admin@example.com";
        description = ''
          Admin emale
        '';
      };

      database = mkOption {
        type = types.submodule ({
          options = import ./db-options.nix { inherit lib; };
        });
        description = ''
          Database settings
        '';
        default = { };
      };
    };
  };

  config = mkIf cfg.enable {

    security.acme = {
      email = cfg.email;
      acceptTerms = true;
      #server = "https://acme-staging-v02.api.letsencrypt.org/directory";
    };

    services.nginx = {
      enable = true;
      statusPage = true;
      recommendedGzipSettings = true;
      virtualHosts."${cfg.domain}" = {
        enableACME = cfg.https;
        forceSSL = cfg.https;
        locations = { "/" = { proxyPass = "http://${proxyPass}"; }; };
      };
    };

    systemd.services.onlinePaymentsGo = {
      description = "Running onlinePaymentsGo";
      wantedBy = [ "multi-user.target" ];
      after = afterService;
      environment = {
        BIND_ADDR = proxyPass;
        DB_URL = dbUrl;
        REDIS_ADDRESS = "localhost:6379";
        TEMPLATE_PATH = templatesPath;
      } // cfg.environment;

      script = ''
        ${pacageBack}/bin/apiserver
      '';

      serviceConfig = {
        User = user;
        Group = group;
      };
    };

    services.redis = { enable = true; };

    systemd.services.runMigrations = {
      description = "Run migrations";
      before = [ "onlinePaymentsGo.service" ];
      after = [ "postgresql.service" ];
      wantedBy = [ "multi-user.target" ];
      script = ''
        ${migratePkg}/bin/migrate -path=${migrationPath} -database ${dbUrl} up
      '';
      serviceConfig = { Type = "oneshot"; };
    };

    users.groups.${group} = { };
    users.users = {
      ${user} = {
        isSystemUser = true;
        extraGroups = [ "postgres" ];
        inherit group;
      };
    };

    services.postgresql = mkIf is_local_db {
      enable = true;
      ensureDatabases = [ cfg.database.name ];
      ensureUsers = [{
        name = cfg.database.user;
        ensurePermissions = {
          "DATABASE ${cfg.database.name}" = "ALL PRIVILEGES";
          "ALL TABLES IN SCHEMA public" = "ALL PRIVILEGES";
        };
      }];
    };
  };
}
