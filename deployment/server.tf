locals {
  namespace     = kubernetes_namespace.server.metadata.0.name
  server_domain = "api.${var.domain}"
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
    host = local.server_domain
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

resource "cloudflare_dns_record" "server" {
  zone_id = cloudflare_zone.main.id
  name    = local.server_domain
  type    = "A"
  content = local.ip
  proxied = false
  ttl     = 1
}

data "http" "ip" {
  url = "https://ipv4.icanhazip.com/"
}

locals {
  ip = trimspace(data.http.ip.response_body)
}

resource "kubernetes_secret" "zitadel" {
  metadata {
    name      = "zitadel"
    namespace = local.namespace
  }

  data = {
    "token.json" = file("${path.module}/../server/token.json")
  }
}
