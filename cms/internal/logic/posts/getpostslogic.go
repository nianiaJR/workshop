package posts

import (
	"context"

	"cms/internal/svc"
	"cms/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsLogic) GetPosts(req *types.GetPostsReq) (resp *types.GetPostsResp, err error) {
	return &types.GetPostsResp{
		List:  []types.Post{},
		Total: 0,
	}, nil
}
