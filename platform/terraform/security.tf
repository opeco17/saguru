module "global_ssh_sg" {
  source      = "./security_group"
  name        = "global_ssh_sg"
  vpc_id      = aws_vpc.gitnavi.id
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

module "global_http_sg" {
  source      = "./security_group"
  name        = "global_http_sg"
  vpc_id      = aws_vpc.gitnavi.id
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

module "global_https_sg" {
  source      = "./security_group"
  name        = "global_https_sg"
  vpc_id      = aws_vpc.gitnavi.id
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}
