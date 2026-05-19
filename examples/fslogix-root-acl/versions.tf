terraform {
  required_version = ">= 1.6.0"

  required_providers {
    azurefilesacl = {
      source  = "compunet/azurefilesacl"
      version = "1.0.0"
    }
  }
}
