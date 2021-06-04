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
- Firepower v6.4.x in standalone mode

Installation
------------
Download and build the latest [release](https://github.com/bluecough/terraform-provider-ciscofdm/releases) and copy to your terraform plugin directory (typically ~/.terraform.d/plugins/)

Alternatively you can build it.
```
$ go mod init terraform-provider-ciscofdm
$ go mod tidy
$ go build
```
## Update for Terraform 0.15
On OSX. This is needed in order to perform a `terraform init`.
```
Create a directory and place the compiled plugin there
$HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/ciscofdm/1.0/darwin_amd64
```
Also create a .terraformrc in your home directory
```
providers {
  customplugin = {
   versions = ["1.0"]
   source = "registry.terraform.io/hashicorp/ciscofdm"
  }
}
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

To place NetworkObjects into groups. Please note adding objects that are not already in the system, into the group will make your terraform state become out of sync.

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
| `name`        | Name of the NetworkObjectGroup you would like to create                         | Required|           |
| `objects`     | Value that can be repeated so that your group can have one or more NetworkObjects. It can only be called 'object'. | Required | object |
| `netname`     | Name of the NetworkObject you would like to add to the group                    | Required|   |
| `type`        | This is the type under the 'objects' key:value. It should always be 'networkobject' | Required | networkobject |
| `type`        | This should always be networkobjectgroup                                        | Required | networkobjectgroup |

### `ciscofdm_portobject`

This creates portobjects for use in access rules. Note that if you create an access port in reverse that it will throw an error. ie/ 5000-4000 rather than 4000-5000.

#### Example
```
resource "ciscofdm_portobject" "myportobject" {
  name = "My Application Port 4000-5000 Object"
  layer4 = "TCP"
  port = "4000-5000"
}
```
| Property    | Description                                                               | Default  |  Valid Values  |
| ----------- | ------------------------------------------------------------------------- | -------  | -------------- |
| `name`        | Name of the Port Object you want to create.                               | Required |                |
| `layer4`      | Layer4 option of either TCP or UDP                                        | Required | TCP or UDP     |
| `port`        | This can either be a single port, a range of ports. And not comma delimited ports. | Required | 1 or 1-2 and NOT 1-2,5 | 

### `ciscofdm_portobjectgroup`

This is the Port Object group resource, that is similar to the network object group.

#### Example
```
resource "ciscofdm_portobjectgroup" "myportobjectgroup" {
  name = "GSPORTGROUP"
  description = "My Port Group"
  type = "portobjectgroup"
  objects {
    name = "FTP"
    type = "tcpportobject"
  }
  objects {
    name = "GTP_PORTS-2123"
    type = "tcpportobject"
  }
}
```

| Property    | Description                                                               | Default  |  Valid Values  |
| ----------- | ------------------------------------------------------------------------- | -------  | -------------- |
| `name`        | Name of the Port Object Group you want to create.                         | Required |                |
| `description` | Description for the Port Object Group                                     | Optional |                |
| `type`        | The tag is static but I kept it in the config.                            | Required | portobjectgroup| 
| `objects`     | This is the TypeSet heirarchy                                             | Required |                |
| `name`        | This is the name of any existing port object that you want to add         | Required |                |
| `type`        | Only two options but no error check in place so anything else will error  | Required | tcpportobject or udpportobject | 

### `ciscofdm_accessrule`

This is the Access Rule resource. Applying this configuration you can implement access rules to your FDM. Please note there isn't any error checking to see if objects you're calling are actually there(Although not hard to do, it's tedious right now). As an example lets say you enter an arbitrary Intrusion or File Policy into your rule. If it doesnt it exist it will error out.

#### Example(s)
```
resource "ciscofdm_accessrule" "myaccessrules" {
  name = "GSTERRAFORMRULE001"
  ruleaction = "PERMIT"  
  intrusionpolicy = {
    name = "Connectivity Over Security"
    type = "intrusionpolicy"
  }
}

```
| Property      | Description                                                               | Default  |  Valid Values  |
| ------------- | ------------------------------------------------------------------------- | -------  | -------------- |
| `name`        | Name of the Rule you want                                                 | Required | string         |
| `ruleid`      | RuleID so if you want to place rules before or another                    | Optional | int            |
| `sourcezones` | This is a TypeSet so it is declared similarly like "objects" above        |          |                |
| `name`        | Name of an existing zone. Required if sourcezones defined                 | Required | string         |
| `destinationzones` | This is a TypeSet so it is declared similarly like "objects" above | | |
| `name`        | Name of an existing zone. Required if destinationzones defined            | Required | string         |
