resource "aws_lambda_function" "web" {
  function_name = "MyLambdaWeb"

  package_type = "Image"
  image_uri    = "${aws_ecr_repository.web.repository_url}:2"

  role = aws_iam_role.lambda_ecr_role.arn
}
