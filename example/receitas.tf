terraform {
  required_providers {
    receita = {
      source = "terraform.local/balhau/receita"
    }
  }
}

provider "receita" {
  endpoint = "localhost:9999"
}

resource "receita_receita" "receita_one" {
  name   = "Bola de carne"
  author = "Maria Bacalhau"
}
