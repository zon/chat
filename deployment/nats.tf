resource "cloudflare_dns_record" "nats" {
  zone_id = cloudflare_zone.main.id
  name    = "nats.${var.domain}"
  type    = "A"
  content = local.ip
  proxied = false
  ttl     = 1
}
