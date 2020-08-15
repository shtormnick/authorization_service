source $stdenv/setup
PATH=$dpkg/bin:$PATH

tar -xvzf $src

mkdir -p $out/bin/
cp -r migrate* $out/
ln -s $out/migrate.linux-amd64 $out/bin/migrate
