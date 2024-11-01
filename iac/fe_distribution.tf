locals {
  fe_origin_id  = "fe"
  api_origin_id = "api"
}

resource "aws_cloudfront_distribution" "fe" {

  origin {
    domain_name = aws_s3_bucket.fe.bucket_domain_name
    origin_id   = local.fe_origin_id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.origin_access_identity.cloudfront_access_identity_path
    }
  }

  origin {
    domain_name = trimprefix(aws_apigatewayv2_api.this.api_endpoint, "https://")
    origin_id   = local.api_origin_id

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  default_root_object = "index.html"
  enabled             = true
  is_ipv6_enabled     = true

  #   aliases = []

  ordered_cache_behavior {
    path_pattern = "/api/*"

    allowed_methods = [
      "GET",
      "HEAD",
      "OPTIONS",
      "PUT",
      "PATCH",
      "POST",
      "DELETE",
    ]

    cached_methods = [
      "GET",
      "HEAD",
      "OPTIONS",
    ]

    viewer_protocol_policy   = "redirect-to-https"
    origin_request_policy_id = data.aws_cloudfront_origin_request_policy.all_viewer_except_host.id
    cache_policy_id          = data.aws_cloudfront_cache_policy.caching_disabled.id

    target_origin_id = local.api_origin_id
  }

  default_cache_behavior {
    allowed_methods = [
      "GET",
      "HEAD",
    ]

    cached_methods = [
      "GET",
      "HEAD",
    ]

    # function_association {
    #   event_type   = "viewer-request"
    #   function_arn = module.fe_function.function.arn
    # }

    origin_request_policy_id = data.aws_cloudfront_origin_request_policy.all_viewer_except_host.id
    cache_policy_id          = data.aws_cloudfront_cache_policy.caching_disabled.id

    target_origin_id = local.fe_origin_id

    viewer_protocol_policy = "redirect-to-https"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
    #   acm_certificate_arn      = aws_acm_certificate.fe.arn
    #   ssl_support_method       = "sni-only"
    #   minimum_protocol_version = "TLSv1"
  }
}

resource "aws_cloudfront_origin_access_identity" "origin_access_identity" {
  comment = "${local.name}-fe"
}
