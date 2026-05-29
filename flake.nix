{
  description = "Reproducible development environment for helm-deploy";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            bashInteractive
            gnumake
            go_1_26
            kind
            kubectl
            kubernetes-helm
            nil
            nixfmt-rfc-style
            nodejs_24
          ];

          shellHook = ''
            echo "Entered helm-deploy dev shell"
          '';
        };

        formatter = pkgs.nixfmt-rfc-style;
      });
}
