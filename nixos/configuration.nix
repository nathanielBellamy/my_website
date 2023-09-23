# This config originates as the result of following the Linode NixOS provisioning guide here:
# https://www.linode.com/docs/guides/install-nixos-on-linode/
# 
# Further alterations have been made - eg. nginx + certbot

{ config, pkgs, ... }:

{
  imports =
    [ # Include the results of the hardware scan.
      ./hardware-configuration.nix
    ];

  # Use the GRUB 2 boot loader.
  boot.loader.grub.enable = true;
  # Define on which hard drive you want to install Grub.
  boot.loader.grub.device = "nodev";
  boot.loader.timeout = 10;

  # nginx
  services.nginx = {
    enable = true;
    recommendedGzipSettings = true;
    recommendedOptimisation = true;
    recommendedProxySettings = true;
    recommendedTlsSettings = true;
    proxyTimeout = "1d";

    virtualHosts."mydomain.dev" = {
      forceSSL = true;
      enableACME = true;
      serverName = "mydomain.dev";
      #root = "/path/to/static/assets/to/serve";
      locations."/"= {
        proxyPass = "http://localhost:8080";
        proxyWebsockets = true;
      };
    };

    # HTTPS redirect secondary domain
    virtualHosts."mydomain.com" = {
      forceSSL = true;
      enableACME = true;
      serverName = "mydomain.com";
      globalRedirect = "mydomain.dev";
    };
  };

  # Certbot
  security.acme.acceptTerms = true;
  security.acme.email = "example@email.com";
  security.acme.defaults.email = "example@email.com";
  security.acme.certs = {
     "mydomain.dev".email = "example@email.com";
     "mydomain.com".email = "example@email.com";
  };

  # Define a user account. Don't forget to set a password with ‘passwd’.
  users.users.nate = {
    isNormalUser = true;
    extraGroups = [ "wheel" "networkmanager" ]; # Enable ‘sudo’ for the user.
    openssh.authorizedKeys.keys = ["ssh-ed25519  <Rest of SSH KEY>"];
   };

  # List packages installed in system profile. To search, run:
  # $ nix search wget
  environment.systemPackages = with pkgs; [
    certbot
    git
    go
    inetutils
    lsof
    mtr
    nginx
    sysstat
    vim # The Nano editor is also installed by default.
    wget
  ];

  # List services that you want to enable:

  # Enable the OpenSSH daemon.
  services.openssh.enable = true;
  services.openssh.settings.PermitRootLogin = "no";

  # Disable predictable interface names
  networking.usePredictableInterfaceNames = false;
  networking.useDHCP = false;
  networking.interfaces.eth0.useDHCP = true;

  # Open ports in the firewall.
  networking.firewall.allowedTCPPorts = [ 22 80 443 ]; # 80 - certs, 22 - http, 443 - https, nginx forces ssl
  networking.firewall.allowedUDPPorts = [ 8080 ]; # Allow Go to serve on 8080
  # Or disable the firewall altogether.
  # networking.firewall.enable = false;

  # Copy the NixOS configuration file and link it from the resulting system
  # (/run/current-system/configuration.nix). This is useful in case you
  # accidentally delete configuration.nix.
  # system.copySystemConfiguration = true;

  # This value determines the NixOS release from which the default
  # settings for stateful data, like file locations and database versions
  # on your system were taken. It's perfectly fine and recommended to leave
  # this value at the release version of the first install of this system.
  # Before changing this value read the documentation for this option
  # (e.g. man configuration.nix or on https://nixos.org/nixos/options.html).
  system.stateVersion = "23.05"; # Did you read the comment?
}
