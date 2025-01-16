
variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  default = "0.0.2"
}

group "default" {
  targets = ["chronicles-of-aethelgard"]
}

target "chronicles-of-aethelgard" {
  context = "."
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/chronicles-of-aethelgard:${TAG}"]
}

# docker login -u username -p token
# docker buildx bake --push --file docker-bake.hcl