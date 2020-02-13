.PHONY: proto

proto:
	protoc --go_out=. *.proto
	protoc-go-inject-tag -input=user.pb.go