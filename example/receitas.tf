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

resource "receita_receita" "bola_carne" {
  name = "Bola de carne"
  #name   = "Batata frita"
  author = "Maria Bacalhau"
}

resource "receita_receita" "bacalhau_todos" {
  name   = "Bacalhau com todos"
  author = "Antonio Mariscada"
}

resource "receita_receita" "pato_bravo" {
  name   = "Pato bravo"
  author = "Jose Pato"
}
