# Receipt terraform example provider

## What?

This is a terraform provider to serve as example. This will give an API to generate receipts based on a terraform 
provider.

## How?

First step in this journey was to install go language. For that you can just follow something like [this](https://go.dev/doc/install).

Next we need to init this as a [go module](https://go.dev/doc/tutorial/create-module) this is basically to help us to manage go dependencies.

For this first step we ended with

```sh
go mod init balhau.net/receita-provider 
```

the output of this command was the creation of the `go.mod` file. 

Since the goal here is to build a terraform provider we must add the terraform libraries needed for that. This can be achieved with 

```sh
go get github.com/hashicorp/terraform-plugin-framework
```

After this command we ended with the following line 
```sh
require github.com/hashicorp/terraform-plugin-framework v1.4.2 // indirect
```

added to our `go.mod` file. A `go.sum`, containing the checksum of the download dependencies, is also created
