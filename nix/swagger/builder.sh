source $stdenv/setup
PATH=$dpkg/bin:$PATH

mkdir -p $out/www/
cp -r $src $out/www/swagger.yaml
