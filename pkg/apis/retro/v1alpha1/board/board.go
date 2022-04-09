package board

import context "context"

type Server struct{}

func (bs *Server) CreateBoard(ctx context.Context, req *CreateBoardRequest) (*CreateBoardResponse, error) {
	return nil, nil
}

func (bs *Server) GetBoard(ctx context.Context, req *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, nil
}

func (bs *Server) ListBoards(_ context.Context, _ *ListBoardRequest) (*ListBoardResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (bs *Server) UpdateBoard(_ context.Context, _ *UpdateBoardRequest) (*UpdateBoardResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (bs *Server) DeleteBoard(_ context.Context, _ *DeleteBoardRequest) (*DeleteBoardResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (bs *Server) GetBoardTasks(_ *GetBoardTasksRequest, _ Board_GetBoardTasksServer) error {
	panic("not implemented") // TODO: Implement
}
