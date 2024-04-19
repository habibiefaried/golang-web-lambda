resource "aws_lambda_function" "web" {
  function_name = "MyLambdaWeb"

  package_type = "Image"
  image_uri    = "${aws_ecr_repository.web.repository_url}:6"

  role = aws_iam_role.lambda_ecr_role.arn
}

# Creating a Lambda Function URL
resource "aws_lambda_function_url" "lambda_function_url" {
  function_name      = aws_lambda_function.web.function_name
  authorization_type = "NONE" # Use "AWS_IAM" for IAM-based authorization

  cors {
    allow_credentials = false
    allow_headers     = ["*"]
    allow_methods     = ["*"]
    allow_origins     = ["*"]
    max_age           = 3600
  }
}

output "lambda_function_url" {
  value = aws_lambda_function_url.lambda_function_url.function_url
}