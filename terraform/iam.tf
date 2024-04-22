resource "aws_iam_role" "lambda_ecr_role" {
  name = "lambda_ecr_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Action = "sts:AssumeRole",
      },
    ],
  })
}

resource "aws_iam_policy" "lambda_ecr_policy" {
  name = "lambda_ecr_policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability",
        ],
        Resource = "${aws_ecr_repository.web.arn}",
      },
      {
        Effect = "Allow",
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:CreateLogGroup",
        ],
        Resource = "arn:aws:logs:*:*:*",
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "lambda_ecr_policy_attach" {
  role       = aws_iam_role.lambda_ecr_role.name
  policy_arn = aws_iam_policy.lambda_ecr_policy.arn
}

resource "aws_iam_policy" "lambda_svc_policy" {
  name = "lambda_svc_policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "network-firewall:UpdateRuleGroup",
          "network-firewall:DescribeRuleGroup",
          "network-firewall:ListRuleGroups",
        ],
        Resource = "*",
        Effect   = "Allow"
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "lambda_svc_policy_attach" {
  role       = module.lambda_function_container_image.lambda_role_name
  policy_arn = aws_iam_policy.lambda_svc_policy.arn
}
