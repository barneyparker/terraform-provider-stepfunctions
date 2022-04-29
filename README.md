# Terraform Provider Sfn

Terraform Data Provider for Step Functions

Run the following command to build the provider

```shell
make init
make vendor
make build
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```