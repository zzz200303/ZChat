package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	if u, ok := ctx.Value("uid").(string); ok {
		return u
	}
	return ""
}

func GetName(ctx context.Context) string {
	if u, ok := ctx.Value("name").(string); ok {
		return u
	}
	return ""
}
