module "lambda_function_container_image" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "7.2.6"

  function_name = "web"
  description   = "My example web"

  create_package             = false
  create_lambda_function_url = true

  image_uri    = "${aws_ecr_repository.web.repository_url}:12"
  package_type = "Image"
}

output "lambda_function_url" {
  value = module.lambda_function_container_image.lambda_function_url
}
