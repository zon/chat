terraform {
  backend "s3" {
    bucket                      = "haralovich-lab-terraform-state"
    key                         = "wurbs.json"
    skip_credentials_validation = true
    skip_region_validation      = true
    skip_requesting_account_id  = true
    skip_s3_checksum            = true
    endpoints = {
      sts = "https://s3.us-east-005.backblazeb2.com"
      s3  = "https://s3.us-east-005.backblazeb2.com"
    }
    region  = "us-east-005"
  }
}