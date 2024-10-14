hospital:
	go run server/hospital.go

alice:
	go run client/patient.go -name=Alice -address=localhost:5001 -input=1

bob:
	go run client/patient.go -name=Bob -address=localhost:5002 -input=2

charlie:
	go run client/patient.go -name=Charlie -address=localhost:5003 -input=3

cert:
	cd cert; chmod +x gen.sh; ./gen.sh; cd ..

.PHONY: gen clean server client test cert