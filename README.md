## Card-Service
#### Build container
- `make docker-build`  
#### Start container
- `make docker-up`  
***
#### To run local development  
- `make run`  
#### To run test  
- `make test`  
***
#### The service will serve as `http://localhost:1500` which contain path below  
- *`GET /card` for getting account data*  
- *`GET /health` for service health checking*  
***
##### *See Makefile for other command*