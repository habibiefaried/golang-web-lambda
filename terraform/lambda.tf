module "lambda_function_container_image" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "7.2.6"

  function_name = "web"
  description   = "My example web"

  create_package             = false
  create_lambda_function_url = true

  image_uri    = "${aws_ecr_repository.web.repository_url}:${var.image_version}"
  package_type = "Image"

  environment_variables = {
    RULEGROUPNAME  = "test"
    COUNTERSSMPATH = "/app/service/counter"
    HOME_NET       = "10.0.0.0/24"
  }
}
