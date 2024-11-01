
resource "aws_s3_bucket" "fe" {
  bucket = "${local.name}-fe"
}

resource "aws_s3_bucket_server_side_encryption_configuration" "fe" {
  bucket = aws_s3_bucket.fe.bucket

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_acl" "fe" {
  bucket = aws_s3_bucket.fe.id
  acl    = "private"
  depends_on = [
    aws_s3_bucket_ownership_controls.fe
  ]
}

resource "aws_s3_bucket_versioning" "fe" {
  bucket = aws_s3_bucket.fe.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_policy" "fe" {
  bucket = aws_s3_bucket.fe.id
  policy = data.aws_iam_policy_document.s3_bucket_policy.json
}

resource "aws_s3_bucket_public_access_block" "fe" {
  bucket = aws_s3_bucket.fe.id

  block_public_acls       = true
  block_public_policy     = true
  restrict_public_buckets = true
  ignore_public_acls      = true
}

data "aws_iam_policy_document" "s3_bucket_policy" {
  statement {
    sid = "1"

    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.fe.arn}/*",
    ]

    principals {
      type = "AWS"

      identifiers = [
        aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn,
      ]
    }
  }
}

resource "aws_s3_bucket_ownership_controls" "fe" {
  bucket = aws_s3_bucket.fe.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

// Deploy the FE code
module "fe_dir" {
  source   = "hashicorp/dir/template"
  version  = "1.0.2"
  base_dir = "${path.module}/../fe/build/"
}

resource "aws_s3_object" "fe" {
  for_each     = module.fe_dir.files
  bucket       = aws_s3_bucket.fe.id
  key          = each.key
  source       = each.value.source_path
  content      = each.value.content
  etag         = each.value.digests.md5
  content_type = each.value.content_type
}
