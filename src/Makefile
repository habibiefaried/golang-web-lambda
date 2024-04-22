test:
	go clean -testcache ; \
	HOME_NET=10.0.0.0/24 RULEGROUPNAME=test COUNTERSSMPATH=/app/service/counter go test -v ./...
build:
	export VERSION=20; \
	docker build . --platform linux/amd64 -t 985417314751.dkr.ecr.ap-northeast-1.amazonaws.com/lambda-container-repo:$$VERSION && \
	docker push 985417314751.dkr.ecr.ap-northeast-1.amazonaws.com/lambda-container-repo:$$VERSION; \
	cd ../terraform/ && \
	terraform apply -var="image_version=$$VERSION"