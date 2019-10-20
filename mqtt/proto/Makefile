.PHONY: proto

proto:
	protoc --go_out=. *.proto
	protoc-go-inject-tag -input=mqtt.pb.go