{
  description = "CUHWC Signup Tool";

  inputs.utils.url = "github:numtide/flake-utils";
  inputs.devshell = {
    url = "github:numtide/devshell";
    inputs.utils.follows = "utils";
  };

  outputs = {
    self,
    nixpkgs,
    utils,
    devshell,
  }:
    utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [devshell.overlays.default];
      };
    in {
      devShells.default =
        pkgs.devshell.mkShell {};
      formatter = pkgs.alejandra;
    });
}
