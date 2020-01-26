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
Download and build the latest [release](https://github.com/bluecough/terraform-provider-ciscofdm/releases) and copy to your terraform plugin directory (typically ~/.terraform.d/plugins/)

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

A resource for managing FDM NetworkObjects. There are a couple of fields that can have certain values but no checking of input has been done.

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

| Property             | Description                                                                     | Default |  Valid Values      |
| -------------------- | ------------------------------------------------------------------------------- | ------- | ------- |
| `name`               | Name you wish to call the network object                                        |  Required|       |
| `subtype`            | String type that can only be HOST or NETWORK.                                   |  Required|  HOST, NETWORK  |
| `value`              | If its a host simply put the IP address. If its a NETWORK X.X.X.X/YY            |  Required |  X.X.X.X , X.X.X.X/YY   |


### `ciscofdm_networkobjectgroup`

To place NetworkObjects into groups. Please note adding objects that are not already in the system into the group will make your terraform state become out of sync.

#### Example

```
resource "ciscofdm_networkobjectgroup" "myobjectgroup" {
  name = "myNetworkGroup"
  objects {
     netname = "any-ipv6"
     type = "networkobject"
  }
  objects {
     netname = "SomeOtherNet"
     type = "networkobject"
  }
  type = "networkobjectgroup"
}
```

| Property    | Description                                                                     | Default |  Valid Values |
| ----------- | ------------------------------------------------------------------------------- | ------- | --------- |
| name        | Name of the NetworkObjectGroup you would like to create                         | Required|           |
| objects     | Value that can be repeated so that your group can have one or more NetworkObjects. It can only be called 'object'. | Required | object |
| netname     | Name of the NetworkObject you would like to add to the group                    | Required|   |
| type        | This is the type under the 'objects' key:value. It should always be 'networkobject' | Required | networkobject |
| type        | This should always be networkobjectgroup                                        | Required | networkobjectgroup |



