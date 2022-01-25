# ECHO

echostart: echoserver.PID

echoserver.PID:
	go run ./cmd/echo/main.go & echo $$! > $@;

echotest: echoserver.PID
	until curl -o /dev/null -s localhost:8080/cloudacademy; do sleep 1; done
	curl -v localhost:8080/cloudacademy && echo
	curl -v localhost:8080/cloudacademy/courses/100 && echo
	curl -v --request PUT localhost:8080/cloudacademy/courses/100 && echo
	curl -v --request DELETE localhost:8080/cloudacademy/courses/100 && echo
	curl -v --header "Content-Type: application/json" --data '{"title":"Golang + Echo","length":120}' localhost:8080/cloudacademy/courses && echo

echostop: echoserver.PID
	pkill -TERM -P `cat echoserver.PID`
	rm echoserver.PID

.PHONY: echostart echostop

echoall: echostart echotest echostop

# ================
# GORILLA

gorillastart: gorillaserver.PID

gorillaserver.PID:
	go run ./cmd/gorilla/main.go & echo $$! > $@;

gorillatest: gorillaserver.PID
	until curl -o /dev/null -s localhost:8080/ok; do sleep 1; done
	curl -v localhost:8080/languages && echo
	curl -v localhost:8080/languages/go && echo
	curl -v localhost:8080/languages/java && echo
	curl -v --header "Content-Type: application/json" --data '{"usecase":"blah","Rank": 3,"compiled":false}' localhost:8080/languages/python && echo
	curl -v --request DELETE localhost:8080/languages/java && echo

gorillastop: gorillaserver.PID
	pkill -TERM -P `cat gorillaserver.PID`
	rm gorillaserver.PID

.PHONY: gorillastart gorillastop

gorillaall: gorillastart gorillatest gorillastop

# ================
# COBRA

cobraclear:
	rm -rf ~/.rocketctl
	rm -rf $(CURDIR)/out/bin 

cobrabuild:
	mkdir -p $(CURDIR)/out/bin
	cd cmd/cobra && GO111MODULE=on go build -o $(CURDIR)/out/bin/rocketctl .

cobraclihelp:
	cd $(CURDIR)/out/bin && ./rocketctl create rocket --help
	echo
	cd $(CURDIR)/out/bin && ./rocketctl launch rocket --help

cobraclicreate:
	cd $(CURDIR)/out/bin && ./rocketctl create rocket r1 --type=saturnv --mission=apollo11 --fuel=5000 --maxspeed=25000

cobraclilaunch:
	cd $(CURDIR)/out/bin && ./rocketctl launch rocket r1 --countdown=10

cobraall: cobraclear cobrabuild cobraclihelp cobraclicreate cobraclilaunch

.PHONY: cobraclear cobraclihelp cobraclicreate cobraclilaunch

# ================

all: echoall gorillaall cobraall