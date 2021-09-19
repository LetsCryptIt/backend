# Backend
Backend code for LetCryptIt 

## Project Setup

### Go Setup
* `git clone https://github.com/LetsCryptIt/backend.git`
* `cd backend`
* `go mod download`. 

### Firebase Setup
* Create a Firebase Project
* Authentication -> Enable Email Login
* Settings > Service Accounts
* Click **Generate New Private Key**, then confirm by clicking **Generate Key**
* Securely store the JSON file containing the key
* export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"

### Start backend
* `go run main.go`
