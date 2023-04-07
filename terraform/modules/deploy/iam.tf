data "aws_iam_policy_document" "assume_role" {
  statement {
    sid    = ""
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["codedeploy.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "main" {
  count = length(var.groups) > 0 ? 1 : 0

  name               = "${local.app_name}-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "main" {
  count = length(var.groups) > 0 ? 1 : 0

  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRoleForLambda"
  role       = aws_iam_role.main[count.index].name
}
