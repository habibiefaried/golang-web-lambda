resource "aws_ecr_repository" "web" {
  name                 = "lambda-container-repo" # Name of the repository
  image_tag_mutability = "IMMUTABLE"
}
