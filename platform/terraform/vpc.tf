resource "aws_vpc" "saguru" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "saguru-vpc"
  }
}

resource "aws_subnet" "public_subnet" {
  vpc_id                  = aws_vpc.saguru.id
  cidr_block              = "10.0.64.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "saguru-public-subnet"
  }
}

resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = aws_vpc.saguru.id
  tags = {
    Name = "saguru-internet-gateway"
  }
}

resource "aws_eip" "nat_gateway_ip" {
  vpc        = true
  depends_on = [aws_internet_gateway.internet_gateway]
}

resource "aws_route_table" "public_table" {
  vpc_id = aws_vpc.saguru.id
  tags = {
    Name = "saguru-public-route-table"
  }
}

resource "aws_route" "public_routes" {
  route_table_id         = aws_route_table.public_table.id
  gateway_id             = aws_internet_gateway.internet_gateway.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table_association" "cluster_associations" {
  subnet_id      = aws_subnet.public_subnet.id
  route_table_id = aws_route_table.public_table.id
}