package task

import (
	context "context"
	"fmt"
)

type Server struct{}

func (ts *Server) CreateTask(_ context.Context, _ *CreateTaskRequest) (*CreateTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) GetTask(_ context.Context, tr *GetTaskRequest) (*GetTaskResponse, error) {
	fmt.Println(tr.Uuid)
	panic("not implemented") // TODO: Implement
}

func (ts *Server) ListTasks(_ context.Context, _ *ListTasksRequest) (*ListTasksResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) UpdateTask(_ context.Context, _ *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) DeleteTask(_ context.Context, _ *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) UpvoteTask(_ context.Context, _ *UpvoteTaskRequest) (*UpvoteTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) MergeTask(_ context.Context, _ *MergeTaskRequest) (*MergeTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (ts *Server) GetTasks(_ *GetTasksRequest, _ Task_GetTasksServer) error {
	panic("not implemented") // TODO: Implement
}
