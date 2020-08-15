{ pkgs ? import <nixpkgs> { } }:
with pkgs;

buildGoModule {
  pname = "authorization-service";
  version = "0.0.1";
  src = ./.;
  goPackagePath = "github.com/shitikovkirill/auth-service";
  modSha256 = "11m9ywcmk9pfmnckz847bjdxwlmbfyj3z81zgibsbxiwj7k5s9sb";
}
