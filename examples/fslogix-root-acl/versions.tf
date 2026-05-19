terraform {
  required_version = ">= 1.6.0"

  required_providers {
    azurefilesacl = {
      source  = "compunet-cloud/azurefilesacl"
      version = "0.0.0"
    }
  }
}
