DB_USER?=""
DB_PASS?=""
GL_BIN := retrogl-dbloader
GL_MAIN := cmd/retrogl-dbloader/main.go
SCHED_BIN := retrosched-dbloader
SCHED_MAIN := cmd/retrosched-dbloader/main.go
RELEASE_DIR := release/

all: vet  build_all

gl_build: $(GL_MAIN) 
	go build -o bin/$(GL_BIN) $(GL_MAIN)

sched_build: $(SCHED_MAIN) 
	go build -o bin/$(SCHED_BIN) $(SCHED_MAIN)

build_all: sched_build gl_build

vet: 
	go vet -all ./internal/... ./cmd/databank-dbloader/... ./cmd/retrosched-dbloader/... ./cmd/retrogl-dbloader/...

test: 
	go test ./...

release_dir:
	-mkdir $(RELEASE_DIR)
