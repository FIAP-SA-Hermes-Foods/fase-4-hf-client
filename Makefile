build-proto:
	rm -f ./client_proto/client.pb.go ./client_proto/client_grpc.pb.go;
	protoc \
	--go_out=client_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=client_proto \
	--go-grpc_opt=paths=source_relative \
	client.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply -auto-approve;

run-bdd:
	docker build -f ./infrastructure/docker/Dockerfile.go_app_bdd -t hf-client-bdd:latest .;
	docker run --rm --name hf-client-bdd hf-client-bdd:latest
	@docker rmi -f hf-client-bdd >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1
