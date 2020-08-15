with import <nixpkgs> { };

mkShell {
  buildInputs = [ go golint golangci-lint gcc python37Packages.pre-commit ];
}
