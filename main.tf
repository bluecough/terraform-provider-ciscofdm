provider "ciscofdm" {
  api_url = "192.168.128.30"
  username = "admin"
  password = "Admin123"
  ssl_no_verify = true 
}

#resource "ciscofdm_dummy" "my-server" {
#  address = "1.2.3.4"
#}

resource "ciscofdm_networkobject" "myobject" {
  name = "GS-Terraform"
  subtype = "HOST"
  value = "2.2.2.3"
}

resource "ciscofdm_networkobjectgroup" "mygroup" {
  name = "NetworkGroupGS"
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
