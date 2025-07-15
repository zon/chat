locals {
  bucket_path = "/file/${b2_bucket.client.bucket_name}"
}

resource "b2_bucket" "client" {
  bucket_name = var.b2_bucket_name
  bucket_type = "allPublic"
}

resource "cloudflare_zone" "main" {
  account = {
    id = var.cloudflare_account_id
  }
  name = var.domain
}

resource "cloudflare_dns_record" "client" {
  zone_id = cloudflare_zone.main.id
  name    = var.domain
  type    = "CNAME"
  content = var.b2_domain
  proxied = true
  ttl     = 1
}

resource "cloudflare_ruleset" "main" {
  zone_id = cloudflare_zone.main.id
  name    = "default"
  kind    = "zone"
  phase   = "http_request_transform"

  rules = [{
    ref         = "assets"
    description = "Assets"
    expression  = "(starts_with(http.request.uri.path, \"/assets\") and http.host eq \"${var.domain}\")"
    action      = "rewrite"
    action_parameters = {
      uri = {
        path = {
          expression = "concat(\"${local.bucket_path}\", http.request.uri.path)"
        }
      }
    }
  }, {
    ref         = "favicon"
    description = "Favicon"
    expression  = "(http.request.uri.path eq \"/favicon.ico\" and http.host eq \"${var.domain}\")"
    action      = "rewrite"
    action_parameters = {
      uri = {
        path = {
          value = "/${local.bucket_path}/favicon.ico"
        }
      }
    }
  }, {
    ref         = "catch-all"
    description = "Catch all"
    expression  = "(http.host eq \"${var.domain}\" and not starts_with(http.request.uri.path, \"/assets\") and http.request.uri.path ne \"/favicon.ico\")"
    action      = "rewrite"
    action_parameters = {
      uri = {
        path = {
          value = "${local.bucket_path}/index.html"
        }
      }
    }
  }]
}
