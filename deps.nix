{ pkgs ? import <nixpkgs> {} }:

[
  pkgs.go
  pkgs.nodejs_21
  pkgs.rustup
  pkgs.wasm-pack
]
