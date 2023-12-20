# Receita terraform example provider

## What?

This is a terraform provider to serve as example. This will give an API to generate receipts based on a terraform 
provider.

## Building?

First step in this journey was to install go language. For that you can just follow something like [this](https://go.dev/doc/install).

Next we need to init this as a [go module](https://go.dev/doc/tutorial/create-module) this is basically to help us to manage go dependencies.

For this first step we ended with

```sh
go mod init balhau.net/receita-provider 
```

the output of this command was the creation of the `go.mod` file. 

Since the goal here is to build a terraform provider we must add the terraform libraries needed for that. This can be achieved with 

```sh
go get
```

To be able to use this terraform plugin you'll need to build this go project with

```sh
go build
```

And then move the binary into your terraform plugin directory, something like this

```
~/.terraform.d/plugins/terraform.local/balhau/receita/1.0.0/darwin_amd64/terraform-provider-receita
```

## Running?

### Setting the http mock server

After previous steps were taken of we are ready to run the receita terraform provider. The provider will do some http calls to the endpoint defined in the provider definition. This means that we should put a service running that matches the description in the endpoint definition

```
provider "receita" {
  endpoint = "http://localhost:9999"
}
```

### Running this

Managing terraform providers involves a bit of typing in the command line. To avoid typing to much we have here a `Makefile` with most of the command line instructions we need.

To figure out how to use this makefile just type

```sh
make help
```
