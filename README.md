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


For that we can use http module from python3 and run

```sh
python3 -m http.server 9999

```

This will open a http server in the port `9000`. This will not implement the endpoints defined in the code but will be enough to log the calls and this will be enough for us to validate the provider lifecycle.

### Running the provider

To run this provider we should jump into `example` directory and initialize `terraform` in the folder. For this it is enough to run 


```sh
terraform init
```

If all steps run properly unti this moment we should end up with a success command. Next step is the `plan`

For this we just 

```sh
terraform plan -out tf_plan.state
```

The end result should be a `json` file containing the terraform state to be executed.
To finally apply the terraform changes we need to run

```sh
terraform apply tf_plan.state
```

If all went good you should have something like this

```
receita_receita.receita_one: Creating...
receita_receita.receita_one: Creation complete after 0s [id=500c3203-7210-45e0-8599-6a4048b78179]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

as output.

At the same time the mock server that we just started should log a call to an endpoint

```
code 404, message File not found
"GET /create HTTP/1.1" 404 -
```

You should have also a `terraform.tfstate` file. Which is a human readable `json` representation of the current terraform state.
Should look something like

```json
{
  "version": 4,
  "terraform_version": "1.4.6",
  "serial": 1,
  "lineage": "bef05e13-6bcf-44f0-a62e-131ade285463",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "receita_receita",
      "name": "receita_one",
      "provider": "provider[\"terraform.local/balhau/receita\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "author": "Maria Bacalhau",
            "id": "500c3203-7210-45e0-8599-6a4048b78179",
            "name": "Bola de carne"
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}
```
