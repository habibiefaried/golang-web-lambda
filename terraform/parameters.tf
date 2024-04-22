resource "aws_ssm_parameter" "counter" {
  name  = "/app/service/counter"
  type  = "String"
  value = "1"

  lifecycle {
    ignore_changes = [
      value,
    ]
  }
}


