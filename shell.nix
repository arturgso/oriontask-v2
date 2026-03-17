{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.pkg-config

    pkgs.webkitgtk_4_1
    pkgs.gtk3
    pkgs.glib
    pkgs.sqlite
  ];
}
