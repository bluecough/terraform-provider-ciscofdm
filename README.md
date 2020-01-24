Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Installation
------------
Download and build the latest release and copy to your terraform plugin directory (typically ~/.terraform.d/plugins/)

Alternatively you can build it.
```
$ go mod init terraform-provider-ciscofdm 
$ go build
```
## Provider Configuration

### Example

Example provider:
```hcl
provider "ciscofdm" {
  api_url = "192.168.128.30"
  username = "admin"
  password = "Admin123"
  ssl_no_verify = true
}
```
| Property            | Description                                                                             | Default    |
| ------------------- | --------------------------------------------------------------------------------------- | ---------- |
| `api_url`           | The IP address of the FDM                                                               | `Required` |
| `username`          | The Username to login to the FDM                                                        | `Required` |
| `password`          | The Password to login to the FDM                                                        | `Required` |
| `ssl_no_verify`     | Boolean to ignore self signed certs                                                     | `Required` |


## Resources
### `ciscofdm_networkobject`

A resource for managing FDM network objects.

#### Example

```hcl
provider "ciscofdm" {
  api_url = "192.168.128.30"
  username = "admin"
  password = "Admin123"
  ssl_no_verify = true
}

resource "ciscofdm_networkobject" "myobject" {
  name = "Terraform Network Object"
  subtype = "HOST"
  value = "2.2.2.3"
}
```
#### Properties

| Property             | Description                                                                     | Default |
| -------------------- | ------------------------------------------------------------------------------- | ------- |
| `name`               | Name you wish to call the network object                                        |  Required       |
| `subtype`            | String type that can only be HOST or NETWORK. Note there is not string checking |  Required       |
| `value`              | If its a host simply put the IP address. If its a NETWORK X.X.X.X/YY . Again there is no checking.           |  Required   |
