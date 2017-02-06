package responses

import (
	"errors"
	"github.com/mesos/mesos-go"
	"github.com/mesos/mesos-go/master"
)

var ErrInvalidResponseType = errors.New("Invalid response type")

// UnmarshalMaster attempts to unmarshal mesos.Response to master.Response
// It will call Close() on the mesos response before returning
func UnmarshalMaster(r mesos.Response, ct master.Response_Type) (*master.Response, error) {
	defer r.Close()
	resp := &master.Response{}
	err := r.Decoder().Invoke(resp)
	if err != nil {
		return nil, err
	}
	if *resp.Type != ct {
		return nil, ErrInvalidResponseType
	}
	return resp, nil
}

// GetAgents is a convenience function for returning mesos.AgentInfos from a mesos.Response
func GetAgents(r mesos.Response, err error) ([]*mesos.AgentInfo, error) {
	if err != nil {
		return nil, err
	}
	resp, err := UnmarshalMaster(r, master.Response_GET_AGENTS)
	if err != nil {
		return nil, err
	}
	infos := []*mesos.AgentInfo{}
	for _, agent := range resp.GetAgents.Agents {
		infos = append(infos, agent.AgentInfo)
	}
	return infos, nil
}

//  GetTasks is a convenience function for returning mesos.Tasks from a mesos.Response
func GetTasks(r mesos.Response, err error) ([]*mesos.Task, error) {
	if err != nil {
		return nil, err
	}
	resp, err := UnmarshalMaster(r, master.Response_GET_TASKS)
	if err != nil {
		return nil, err
	}
	tasks := []*mesos.Task{}
	for _, task := range resp.GetTasks.Tasks {
		tasks = append(tasks, task)
	}
	for _, task := range resp.GetTasks.OrphanTasks {
		tasks = append(tasks, task)
	}
	for _, task := range resp.GetTasks.PendingTasks {
		tasks = append(tasks, task)
	}
	for _, task := range resp.GetTasks.CompletedTasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
