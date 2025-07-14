resource "b2_bucket" "client" {
  bucket_name = "haralovich-wurbs"
  bucket_type = "allPublic"
}

resource "cloudflare_zone" "main" {
  account = {
    id = var.cloudflare_account_id
  }
  name = var.domain
}

resource "cloudflare_dns_record" "main" {
  zone_id = cloudflare_zone.main.id
  name    = "@"
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
    ref         = "catch-all"
    description = "Catch all"
    expression  = "(not starts_with(http.request.uri.path, \"/assets\") and http.host eq \"${var.domain}\")"
    action      = "rewrite"
    action_parameters = {
      uri = {
        path = {
          value = "/file/${b2_bucket.client.bucket_name}/index.html"
        }
      }
    }
    }, {
    ref         = "assets"
    description = "Assets"
    expression  = "(starts_with(http.request.uri.path, \"/assets\") and http.host eq \"${var.domain}\")"
    action      = "rewrite"
    action_parameters = {
      uri = {
        path = {
          expression = "concat(\"/file/${b2_bucket.client.bucket_name}\", http.request.uri.path)"
        }
      }
    }
  }]
}
