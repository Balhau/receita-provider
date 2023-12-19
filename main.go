package main

import(
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/provider"
)

func main(){
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() terraform.ResourceProvider {
      return Provider()
    }
  })
}
