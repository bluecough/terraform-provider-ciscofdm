provider "ciscofdm" {
  api_url = "192.168.128.30"
  username = "admin"
  password = "Admin123"
  ssl_no_verify = true 
}

resource "ciscofdm_dummy" "my-server" {
  address = "1.2.3.4"
}
