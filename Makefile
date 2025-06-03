start:
	@echo "⏳ Starting services..."
	@docker-compose -f docker-compose.yaml up -d localstack

	@echo "⏳ Waiting for LocalStack to be ready..."
	@until curl -s http://localhost:4566/_localstack/health | jq -e '.services.s3 == "available"' > /dev/null; do \
	    echo "Waiting for LocalStack..."; \
	    sleep 2; \
	done

	@echo "✅ LocalStack is up!"
	@echo "☁️  Creating S3 bucket: university"
	@aws --endpoint-url=http://localhost:4566 --region=us-east-1 s3 mb s3://university

	@echo "🚀 Starting remaining services..."
	@docker-compose -f docker-compose.yaml up --build 

stop:
	@echo "🛑 Stopping all services..."
	@docker-compose -f docker-compose.yaml down
kill:
	sudo lsof -t -i -P -n | xargs sudo kill -9

aws-health:
	curl -s http://localhost:4566/_localstack/health | jq .
