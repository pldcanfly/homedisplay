{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  packages = with pkgs; [ 
  cairo
  gtk3
  pkg-config
  ];
}


