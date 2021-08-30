{
  description = "OKEx API CLI client";

  inputs.devshell.url = "github:numtide/devshell";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = inputs@{ self, nixpkgs, devshell, flake-utils }:
    with flake-utils.lib;
    {
      overlay = final: prev: {
        gokex = final.callPackage ./package.nix { };
      };
    } // eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        gokex = pkgs.callPackage ./package.nix { };
      in
      {
        packages = { inherit gokex; };
        defaultPackage = gokex;
        defaultApp = mkApp { drv = gokex; };
        devShell = devshell.legacyPackages.${system}.fromTOML ./devshell.toml;
      });
}
