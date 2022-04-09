package retro

import (
	"github.com/necatican/agile-tools/pkg/apis/retro/v1alpha1/board"
	"github.com/necatican/agile-tools/pkg/apis/retro/v1alpha1/task"
	"google.golang.org/grpc"
)

func RegisterServers(grpcServer *grpc.Server) {
	board.RegisterBoardServer(grpcServer, &board.Server{})
	task.RegisterTaskServer(grpcServer, &task.Server{})
}
