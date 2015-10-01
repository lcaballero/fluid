test: .FORCE
	go test fluid/req
	go test fluid/calls

install: .FORCE
	go install fluid

start-es:
	elasticsearch .files/elasticsearch.yml

.FORCE:


