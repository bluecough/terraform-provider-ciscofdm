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
| Property            | Description                                                                                                           | Default    |
| ------------------- | --------------------------------------------------------------------------------------------------------------------- | ---------- |
| `bootstrap_servers` | A list of host:port addresses that will be used to discover the full set of alive brokers                             | `Required` |
| `ca_cert`           | The CA certificate or path to a CA certificate file to validate the server's certificate.                             | `""`       |
| `client_cert`       | The client certificate or path to a file containing the client certificate -- Use for Client authentication to Kafka. | `""`       |

