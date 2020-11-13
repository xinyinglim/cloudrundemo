1. [Done] Create tests for cloud run
2. [Done] Deploy via cloudbuild.yaml
3. Auto run function/cloud run testing on develop branch via terraform
4. If build succeeds, you can deploy via main branch.

OR 
3. run tests. If build succe

5. run global unit and integration tests separately, 
6. running build after every certain commit?

Normal docker build
1. docker build -t gcr.io/demos-xy/helloworld:v1 .
2. docker push gcr.io/demos-xy/helloworld:v1
3. docker run -p 8080:8081 gcr.io/demos-xy/helloworld:v1

Give Cloud Build service account the Cloud Run Admin and Service Account User Roles


# To check if running
1. go run server.go
2. curl http://localhost:8081/helloworld
  - This should return helloworld

# Tunning tests
go test -v

# to build, just commit

# todo
try traffic splitting
