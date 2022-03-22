#!make
TASK_PATH=$(sprint)
TASK_TYPE=$(type)
TASK_NAME=$(name)

.PHONY: help run test
.SILENT:
.DEFAULT_GOAL := help

help: ## Помощь
	@echo "Доступные команды:"
	@grep -h -E '^[a-zA-Z_-].+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Для запуска используйте обязательные аргументы запуска:"
	@echo | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", "<sprint>", "номер спринта, например: sprint=1"}'
	@echo | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", "<type>", "тип задачи, например: type={final|task}"}'
	@echo | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", "<name>", "название задачи, например: name=A"}'
	@echo ""
	@echo "Пример запуска задачи:"
	@echo | awk 'BEGIN {printf "\033[36m%-15s\033[0m\n", "make run sprint=1 type=final name=A"}'
	@echo ""
	@echo "Пример запуска теста:"
	@echo | awk 'BEGIN {printf "\033[36m%-15s\033[0m\n", "make test sprint=1 type=task name=С"}'

run: ## Запускает задачу из спринта
	@echo $(RUN_ARGS)

test: ## Проводит тест по задаче
	@echo sprint_${TASK_PATH}/${TASK_TYPE}/${TASK_NAME}
	cd sprint_${TASK_PATH}/${TASK_TYPE}/${TASK_NAME} && go test -v