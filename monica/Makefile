.PHONY: update
update:
	docker pull monica
	docker tag monica gcr.io/dbut-0/monica
	docker push gcr.io/dbut-0/monica

.PHONY: deploy
deploy: update
	gcloud beta run deploy monica \
    --image=gcr.io/dbut-0/monica \
    --allow-unauthenticated \
    --port=80 \
    --max-instances=4 \
    --set-env-vars=APP_ENV=production,APP_URL=https://monica.dylanbutler.net \
    --set-secrets=APP_KEY=monica-app-key:latest,DB_HOST=monica-db-host:latest,DB_PASSWORD=monica-db-password:latest,DB_USERNAME=monica-db-username:latest \
    --platform=managed \
    --region=australia-southeast1 \
    --project=dbut-0
