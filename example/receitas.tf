terraform {
  required_providers {
    receita = {
      source = "terraform.local/balhau/receita"
    }
  }
}

provider "receita" {
  endpoint = "http://localhost:9999"
}

resource "receita_receita" "receita_one" {
  name = "Bola de carne"
  #name   = "Batata frita"
  author = "Maria Bacalhau"
}

resource "receita_receita" "receita_two" {
  name   = "Bacalhau com Todos"
  author = "Antonio Mariscada"
}
