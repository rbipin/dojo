locals {
  name = "terraform-${var.instance_type}"
}

resource "aws_instance" "my_server" {
  tags = {
    name = "Bipin-${local.name}"
  }
  instance_type = var.instance_type
  ami           = "ami-02bf8ce06a8ed6092"
}
