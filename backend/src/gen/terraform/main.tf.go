package terraform

func genMain() string {
	return `terraform {
  required_providers {
    shipa = {
      version = "0.0.13"
      source = "shipa-corp/shipa"
    }
  }
}

provider "shipa" {}
`
}
