{ pkgs ? import <nixpkgs> { } }:

let
  stdenv = pkgs.stdenv;
  dpkg = pkgs.dpkg;

  appName = "swagger-${version}";
  version = "1.0.0";
  description = "Swagger file.";
in stdenv.mkDerivation rec {
  name = appName;
  builder = ./builder.sh;

  src = ../../swagger.yaml;

  meta = with stdenv.lib; {
    inherit description;
    license = licenses.agpl3;
    platforms = platforms.linux;
    maintainers = with maintainers; [ shitikovkirill ];
  };
}
