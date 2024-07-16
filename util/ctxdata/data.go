package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	if u, ok := ctx.Value("payload").(string); ok {
		return u
	}
	return ""
}
