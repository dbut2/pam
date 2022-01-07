cli:
	rm -f pam
	go build -o pam cmd/cli/main.go

server:
	rm -f server
	go build -o server cmd/server/main.go

.PHONY: build
build:
	docker build -t gcr.io/dbut-0/pam .

.PHONY: update
update: build
	docker push gcr.io/dbut-0/pam

.PHONY: deploy
deploy: update
	gcloud beta run deploy pam \
	--image=gcr.io/dbut-0/pam \
	--allow-unauthenticated \
	--max-instances=4 \
	--set-secrets=/secret-config/config.yaml=pam-config:latest,/secret-creds/creds.json=pam-creds:latest \
	--region=us-west1 \
	--project=dbut-0