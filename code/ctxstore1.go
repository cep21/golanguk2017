package reqinfo

type privateCtxType string

var (
	reqID = privateCtxType("req-id")
)

func GetRequestID(ctx context.Context) (int, bool) {
	id, exists := ctx.Value(reqID).(int)
	return id, exists
}

func WithRequestID(ctx context.Context, reqid int) context.Context {
	return context.WithValue(ctx, reqID, reqid)
}
