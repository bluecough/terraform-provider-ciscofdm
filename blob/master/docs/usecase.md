# Startnow - Using the Cisco Secure Firewall Device Manager Provider with Terraform.
Using the CiscoFDM provider you will be able to apply Cisco Secure Firewall Port Objects, Port Object Groups, Network Objects, Network Object Groups, and Access Control Polcies

## Business Case
This resource will walk you through implementing the CiscoFDM Terraform Provider either through downloading the pre-built binaries to building the plugin with the assumption you have setup your Go environment.

### Installing the pre-built binaries
* Assumptions - Terraform is installed on your laptop.

To start using the provider in your environment you first need to download the latest binary for your OS and also the version of the Cisco Secure Firewall version you are running.
The latest binary can be found [here](https://github.com/bluecough/terraform-provider-ciscofdm/releases/tag/v1.0.2).
