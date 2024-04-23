module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "~> 6.0"

  name = "my-alb"

  load_balancer_type = "application"
  vpc_id             = data.aws_vpc.default.id                  # Replace with your VPC ID
  subnets            = data.aws_subnets.default_vpc_subnets.ids # Replace with your subnet IDs

  security_groups = [aws_security_group.alb_sg.id] # Replace with your security group IDs

  http_tcp_listeners = [
    {
      port               = 80
      protocol           = "HTTP"
      target_group_index = 0
    }
  ]

  target_groups = [
    {
      name_prefix      = "web"
      backend_protocol = "HTTP"
      backend_port     = 80
      target_type      = "lambda"
    }
  ]

  tags = {
    Environment = "production"
  }
}

resource "aws_lambda_permission" "allow_alb" {
  statement_id  = "AllowALBInvoke"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda_function_container_image.lambda_function_name
  principal     = "elasticloadbalancing.amazonaws.com"
  source_arn    = module.alb.lb_arn # Using the ARN of the entire ALB
}

resource "aws_lambda_permission" "allow_tg" {
  statement_id  = "AllowTGInvoke"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda_function_container_image.lambda_function_name
  principal     = "elasticloadbalancing.amazonaws.com"
  source_arn    = module.alb.target_group_arns[0]
}


resource "aws_lb_target_group_attachment" "lambda_attachment" {
  target_group_arn = module.alb.target_group_arns[0]
  target_id        = module.lambda_function_container_image.lambda_function_arn
}
