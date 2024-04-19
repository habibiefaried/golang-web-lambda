# API Gateway
resource "aws_api_gateway_rest_api" "my_api" {
  name        = "MyAPI"
  description = "API Gateway to trigger Lambda function"
}

resource "aws_api_gateway_resource" "api_proxy_resource" {
  rest_api_id = aws_api_gateway_rest_api.my_api.id
  parent_id   = aws_api_gateway_rest_api.my_api.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy_method" {
  rest_api_id   = aws_api_gateway_rest_api.my_api.id
  resource_id   = aws_api_gateway_resource.api_proxy_resource.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_proxy_integration" {
  rest_api_id = aws_api_gateway_rest_api.my_api.id
  resource_id = aws_api_gateway_resource.api_proxy_resource.id
  http_method = aws_api_gateway_method.proxy_method.http_method
  type        = "AWS_PROXY"
  integration_http_method = "POST"
  uri         = aws_lambda_function.web.invoke_arn
}

resource "aws_api_gateway_deployment" "api_deployment" {
  depends_on = [aws_api_gateway_integration.lambda_proxy_integration]

  rest_api_id = aws_api_gateway_rest_api.my_api.id
  stage_name  = "prod"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_lambda_permission" "api_gateway_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.web.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.my_api.execution_arn}/*/*/*"
}

output "invoke_url" {
  value = "${aws_api_gateway_rest_api.my_api.execution_arn}/${aws_api_gateway_deployment.api_deployment.stage_name}"
}

output "api_gateway_invoke_url" {
  description = "The URL to invoke the API Gateway"
  value       = "https://${aws_api_gateway_rest_api.my_api.id}.execute-api.ap-northeast-1.amazonaws.com/${aws_api_gateway_deployment.api_deployment.stage_name}"
}