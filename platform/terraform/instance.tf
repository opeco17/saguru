resource "aws_instance" "backend" {
  ami                    = "ami-06ad9296e6cf1e3cf"
  instance_type          = "t3.medium"
  key_name               = aws_key_pair.saguru.id
  vpc_security_group_ids = [module.global_ssh_sg.id, module.global_http_sg.id, module.global_https_sg.id]
  subnet_id              = aws_subnet.public_subnet.id
  # user_data = <<EOF
  # #!/bin/bash

  # ### install Docker
  # sudo yum install -y docker
  # sudo systemctl start docker
  # sudo usermod -a -G docker ec2-user
  # sudo systemctl enable docker

  # ### install docker compose
  # sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/bin/docker-compose
  # sudo chmod +x /usr/bin/docker-compose

  # ### install git
  # sudo yum install -y git
  # EOF
  tags = {
    Name = "saguru-backend"
  }
}

resource "aws_key_pair" "saguru" {
  key_name   = "saguru"
  public_key = file(pathexpand(var.public_key_path))
}

resource "aws_eip" "saguru_backend" {
  instance   = aws_instance.backend.id
  vpc = true
}
