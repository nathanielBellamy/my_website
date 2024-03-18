{ pkgs ? import <nixpkgs> {} }:

let
  dependencies = import ./deps.nix { inherit pkgs; };
in
pkgs.mkShell {
  buildInputs = dependencies;
}
