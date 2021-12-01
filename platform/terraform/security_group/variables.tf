variable "name" {}
variable "vpc_id" {}
variable "from_port" {}
variable "to_port" {}
variable "protocol" {}
variable "cidr_blocks" {
  type = list(string)
}