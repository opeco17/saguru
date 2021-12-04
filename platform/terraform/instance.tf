resource "aws_instance" "backend" {
  ami                    = "ami-06ad9296e6cf1e3cf"
  instance_type          = "t3.small"
  key_name               = aws_key_pair.gitnavi.id
  vpc_security_group_ids = [module.global_ssh_sg.id, module.global_http_sg.id, module.global_https_sg.id]
  subnet_id              = aws_subnet.public_subnet.id
  tags = {
    Name = "gitnavi-backend"
  }
}

resource "aws_key_pair" "gitnavi" {
  key_name   = "gitnavi"
  public_key = file(pathexpand(var.public_key_path))
}

resource "aws_eip" "gitnavi_backend" {
  instance   = aws_instance.backend.id
  vpc = true
}
