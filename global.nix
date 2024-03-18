{ pkgs ? import <nixpkgs> {} }:

let
  dependencies = import ./deps.nix { inherit pkgs; };
in
pkgs.stdenv.mkDerivation {
  name = "my_website_env";
  src = builtins.filterSource (path: type: false) ./.;
  buildInputs = dependencies;
  buildCommand = ''
    touch $out
  '';

  # Adjust the phases if needed, here we keep it simple
}
