module server

go 1.17

require github.com/fetch v0.0.0

require github.com/topo_client v0.0.0

require github.com/proto v0.0.0 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/fetch => ../fetch

replace github.com/topo_client => ../../topology_client/topo_client

replace github.com/proto => ../../topology_client/proto
