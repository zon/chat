locals {
  namespace = kubernetes_namespace.server.metadata.0.name
}

resource "kubernetes_namespace" "server" {
  metadata {
    name = "wurbs"
  }
}

resource "helm_release" "server" {
  name      = "wurbs"
  namespace = local.namespace
  chart     = "${path.module}/../chart"
  values = [yamlencode({
    image = {
      tag = file("${path.module}/../version")
    }
  })]
  wait = false
}

resource "kubernetes_secret" "postgres" {
  metadata {
    name      = "postgres"
    namespace = local.namespace
  }

  data = {
    password = data.kubernetes_secret.postgres.data["postgres-password"]
  }
}

data "kubernetes_secret" "postgres" {
  metadata {
    name      = "postgres"
    namespace = "postgres"
  }
}

