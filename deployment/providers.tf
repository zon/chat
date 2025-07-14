terraform {
  required_providers {
    b2 = {
      source = "Backblaze/b2"
    }
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
    helm = {
      source = "hashicorp/helm"
    }
  }
}

provider "b2" {
}

provider "cloudflare" {
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "microk8s"
}

provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "microk8s"
  }
}
