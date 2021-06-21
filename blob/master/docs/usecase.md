# Startnow - Using the Cisco Secure Firewall Device Manager Provider with Terraform.
Using the CiscoFDM provider you will be able to apply Cisco Secure Firewall Port Objects, Port Object Groups, Network Objects, Network Object Groups, and Access Control Polcies

## Business Case
This resource will walk you through implementing the CiscoFDM Terraform Provider either through downloading the pre-built binaries to building the plugin with the assumption you have setup your Go environment.

### Installing the pre-built binaries
* Assumptions - Terraform is installed on your laptop.

To start using the provider in your environment you first need to download the latest binary for your OS and also the version of the Cisco Secure Firewall version you are running.
The latest binary can be found [here](https://github.com/bluecough/terraform-provider-ciscofdm/releases/tag/v1.0.2).

Since this is a 3rd party plugin that isn't officially on the Hashicorp Providers list. You will need to perform the following in order to install the provider.

On OSX. This is needed in order to perform a `terraform init`.
```
Create a directory and place the compiled binary at the following location.
$HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/ciscofdm/1.0/darwin_amd64
```
On Linux
```
$HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/ciscofdm/1.0/linux_amd64
```

Now in order to use it make sure you have a main.tf file or use the example .tf file from this repository.
```
$ terraform init
```
### Compiling the provider
* Assumptions - You have your Golang development environment setup.

Once you have your Golang environment setup simply run the following commands in your cloned repository. It should compile the binary and you will need to place this binary using the above locations in your home directory based upon the OS platform you are on.
```
$ go mod init terraform-provider-ciscofdm
$ go mod tidy
$ go build
```
Here is a video example of the above.
[![asciicast](https://asciinema.org/a/VfMFEi1pVPf8nsH8XJXO1LIkf.svg)](https://asciinema.org/a/VfMFEi1pVPf8nsH8XJXO1LIkf)
