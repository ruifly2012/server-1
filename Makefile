v ?= latest

.PHONY: build
build:
	make -C apps/game build
	make -C apps/gate build
	make -C apps/combat build
	make -C apps/chat build
	make -C apps/client build

.PHONY: build_win
build_win:
	make -C apps/game build_win
	make -C apps/gate build_win
	make -C apps/combat build_win
	make -C apps/chat build_win
	make -C apps/client build_win

.PHONY: proto
proto:
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/att.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/player.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/hero.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/item.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/token.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/talent.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/rune.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/scene.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/game.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/gate/gate.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/combat/combat.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/account/account.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/pubsub/pubsub.proto

.PHONY: docker
docker:
	make -C apps/game docker
	make -C apps/gate docker
	make -C apps/combat docker
	make -C apps/chat docker

.PHONY: test
test:
	go test -v ./... -cover -coverprofile=test.out

.PHONY: test_html
test_html: test
	go tool cover -html=test.out

.PHONY: benchmark
benchmark:
	go test -bench=. ./...

.PHONY: run
run:
	docker-compose up -d

.PHONY: push
push:
	make -C apps/game push
	make -C apps/gate push
	make -C apps/combat push
	#make -C apps/chat push

.PHONY: push_coding
push_coding:
	make -C apps/game push_coding
	make -C apps/gate push_coding
	make -C apps/combat push_coding

.PHONY: push_github
push_github:
	make -C apps/game push_github
	make -C apps/gate push_github
	make -C apps/combat push_github

.PHONY: clean
clean:
	docker rm -f $(shell docker ps -a -q)
	docker rmi -f $(shell docker images -a -q)

.PHONY: stop
stop:
	docker-compose down
