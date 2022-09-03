package user

import "context"

type Cache interface {
	SetJson(ctx context.Context)
	GetJson(ctx context.Context)
}
