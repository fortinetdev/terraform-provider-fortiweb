# Configure the FortiWEB Provider
# Option 1: Token-based authentication
# provider "fortiweb" {
# 	hostname = "10.159.32.112"
# 	token = "60nHnbQkGm7r1fc1snQzNQwj5tzfdm"
# }

# Option 2: Username/Password-based authentication
provider "fortiweb" {
	hostname = "10.159.32.112"
	username = "your_username"
	password = "your_password"
}

