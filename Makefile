# Указание интерпретатора (опционально)
SHELL := /bin/bash

# Общая цель для запуска всех заданий
run_all:
	@for task in {1..26}; do \
		echo "Running L1.$$task..."; \
		go run L1.$$task/main.go; \
	done

.PHONY: L1.1 L1.2 L1.3 L1.4 L1.5 L1.6 L1.7 L1.8 L1.9 L1.10 \
        L1.11 L1.12 L1.13 L1.14 L1.15 L1.16 L1.17 L1.18 L1.19 L1.20 \
        L1.21 L1.22 L1.23 L1.24 L1.25 L1.26

# Цели для каждого задания
L1.1:
	go run L1.1/main.go

L1.2:
	go run L1.2/main.go

L1.3:
	go run L1.3/main.go

L1.4:
	go run L1.4/main.go

L1.5:
	go run L1.5/main.go

L1.6:
	go run L1.6/main.go

L1.7:
	go run L1.7/main.go

L1.8:
	go run L1.8/main.go

L1.9:
	go run L1.9/main.go

L1.10:
	go run L1.10/main.go

L1.11:
	go run L1.11/main.go

L1.12:
	go run L1.12/main.go

L1.13:
	go run L1.13/main.go

L1.14:
	go run L1.14/main.go

L1.15:
	go run L1.15/main.go

L1.16:
	go run L1.16/main.go

L1.17:
	go run L1.17/main.go

L1.18:
	go run L1.18/main.go

L1.19:
	go run L1.19/main.go

L1.20:
	go run L1.20/main.go

L1.21:
	go run L1.21/main.go

L1.22:
	go run L1.22/main.go

L1.23:
	go run L1.23/main.go

L1.24:
	go run L1.24/main.go

L1.25:
	go run L1.25/main.go

L1.26:
	go run L1.26/main.go

# Очистка кэша Go
clean:
	go clean

# Справка
help:
	@echo "Доступные команды:"
	@echo "  make run_all     - Запустить все задания"
	@echo "  make run_L1.X    - Запустить задание L1.X (где X - номер задания)"
	@echo "  make clean       - Очистить проект"
