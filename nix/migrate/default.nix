{ pkgs ? import <nixpkgs> { } }:

let
  stdenv = pkgs.stdenv;
  dpkg = pkgs.dpkg;

  appName = "migrate-${version}";
  version = "4.11.0";
  description = "Database migrations written in Go.";
in stdenv.mkDerivation rec {
  name = appName;
  builder = ./builder.sh;

  src = pkgs.fetchurl {
    url =
      "https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz";
    sha256 = "0z1q6286ymg1n9xhj3rf2lqqrnf9i6596pvhy39xdlad1vza0fcv";
  };

  meta = with stdenv.lib; {
    inherit description;
    homepage = "https://github.com/golang-migrate/migrate";
    license = licenses.agpl3;
    platforms = platforms.linux;
    maintainers = with maintainers; [ shitikovkirill ];
  };
}
