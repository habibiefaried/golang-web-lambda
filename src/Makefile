test:
	go clean -testcache ; go test -v ./...
build:
	export VERSION=18; \
	docker build . --platform linux/amd64 -t 985417314751.dkr.ecr.ap-northeast-1.amazonaws.com/lambda-container-repo:$$VERSION && \
	docker push 985417314751.dkr.ecr.ap-northeast-1.amazonaws.com/lambda-container-repo:$$VERSION; \
	cd ../terraform/ && \
	terraform apply -var="image_version=$$VERSION"