package calls

import (
	pb "github.com/gogo/protobuf/proto"
	"github.com/mesos/mesos-go/agent"
)

func GetState() *agent.Call {
	return &agent.Call{
		Type: agent.Call_GET_STATE.Enum(),
	}
}

func ListFiles(path string) *agent.Call {
	return &agent.Call{
		Type: agent.Call_LIST_FILES.Enum(),
		ListFiles: &agent.Call_ListFiles{
			Path: pb.String(path),
		},
	}
}

func ReadFile(path string, length, offset uint64) *agent.Call {
	return &agent.Call{
		Type: agent.Call_READ_FILE.Enum(),
		ReadFile: &agent.Call_ReadFile{
			Path:   path,
			Length: length,
			Offset: offset,
		},
	}
}
