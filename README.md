# Azure Infrastructure as Code Demo

Infrastructure as Code Demonstration for Azure.

## Tools Overview

* Terraform
* Azure Container Instance
* Azure Storage Accounts
* Azure CosmosDB

## Prerequsite

* Azure Subscription

## Setup

| Variables | Description |
|-----------|-------------|
| `ARM_CLIENT_ID` | Azure Service Principal ID |
| `ARM_CLIENT_SECRET` | Azure Service Principal Secret |
| `ARM_SUBSCRIPTION_ID` | Azure Subscription ID |
| `ARM_TENANT_ID` | Azure Tenant ID |
| `TF_VAR_resource_group` | Azure Resource Group Name |

Install and set up the [Azure CLI][Azure CLI].

[Azure CLI]: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli

```
az login
az account set --subscription="$ARM_SUBSCRIPTION_ID"
```

Create a new resource group

```
az group create --name azure-iac-demo --location westeurope
```

Create a new Service Principal

```
az ad sp create-for-rbac \
  --name http://$TF_VAR_resource_group \
  --role contributor \
  --scopes /subscriptions/$ARM_SUBSCRIPTION_ID/resourceGroups/$TF_VAR_resource_group
```

## Reference

* https://docs.microsoft.com/en-us/azure/developer/terraform/authenticate-to-azure
* https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/service_principal_client_secret
* https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_group
* https://docs.microsoft.com/en-us/azure/container-instances/container-instances-quickstart-portal
