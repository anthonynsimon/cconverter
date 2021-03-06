package commands

import "context"

func extractAPIHost(ctx context.Context) string {
	value := ctx.Value("apiHost")
	apiHost, ok := value.(string)
	if !ok {
		return ""
	}
	return apiHost
}
