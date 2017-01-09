.DEFAULT_GOAL := help

REPO := devopsgig

IMAGE_PROD_NAME := arithsubscriber

CONTAINER_PROD_NAME := "${IMAGE_PROD_NAME}"

.PHONY: help build/prod/image run clean

help:
	@echo "------------------------------------------------------------------------"
	@echo "devopsgig arithpsubscriber"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build/prod/image: ## build prod image
	@docker build -t  "${REPO}"/"${IMAGE_PROD_NAME}" -f ./resources/prod/Dockerfile .

run: clean build/prod/image ## start service
	@docker run -d -p 8081:8081 --name "${CONTAINER_PROD_NAME}" "${REPO}"/"${IMAGE_PROD_NAME}"

clean: ## stop and remove running production container
	@./scripts/rm-container.sh "${REPO}"/"${IMAGE_PROD_NAME}" "${CONTAINER_PROD_NAME}" &> /dev/null
