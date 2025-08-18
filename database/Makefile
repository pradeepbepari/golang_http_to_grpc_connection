migrations:
	migrate create -ext sql -dir ./migrations -seq $(filter-out $@, $(MAKECMDGOALS))
%:
	@:


up:
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down

.PHONY: migrations
