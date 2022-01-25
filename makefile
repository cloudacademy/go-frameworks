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
