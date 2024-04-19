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
