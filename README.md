# Card-Service
### Build container
- `make docker-build`  
### Start container
- `make docker-up`  
***
### Local Development
- *`make run` To run local development*  
#### The service will serve as `http://localhost:1500` which contain path below  
- *`GET /card` for getting account data*  
- *`GET /health` for service health checking*  
***
### Unit test  
#### To generate mock file with mockery  
- *First, you need to install mockery by using command
`make mock-install`*  
- *Then use this command to generate repository mock 
`mockery --name=<repo_name> --dir=internal/core/port  --output=internal/core/port/mocks --outpkg=mocks`*  
- *`make test-service` to run service unit test*  
***
#### *See Makefile for other command*