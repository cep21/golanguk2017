package mytype

var privateCtxType int
var thingGetter privateCtxType

func GetThing(ctx context.Context) int {
	return ctx.Value(thingGetter).(int)
}

func StoreThing(ctx context.Context, thing int) context.Context {
	return ctx.WithValue(thing)
}
