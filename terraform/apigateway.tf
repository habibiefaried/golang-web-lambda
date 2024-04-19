# API Gateway HTTP API to expose the Lambda function
resource "aws_apigatewayv2_api" "http_api" {
  name          = "lambda_http_api"
  protocol_type = "HTTP"
  description   = "HTTP API for Lambda Function"
}

# Integration between API Gateway and Lambda
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id              = aws_apigatewayv2_api.http_api.id
  integration_type    = "AWS_PROXY"
  integration_uri     = aws_lambda_function.web.invoke_arn
  integration_method  = "POST"
  payload_format_version = "2.0"
}

# Default route to invoke the Lambda function
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.http_api.id
  route_key = "ANY /"
  target    = format("integrations/%s", aws_apigatewayv2_integration.lambda_integration.id)
}

# Deploy the API and make it accessible
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id     = aws_apigatewayv2_api.http_api.id
  name       = "$default"
  auto_deploy = true
}

# Permission for API Gateway to invoke the Lambda function
resource "aws_lambda_permission" "api_gw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.web.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http_api.execution_arn}/*/*"
}

output "lambda_function_url" {
  value = "${aws_apigatewayv2_api.http_api.api_endpoint}/"
}