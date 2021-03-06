type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// Create a group and a context to use with that group
func WithContext(ctx context.Context) (*Group, context.Context) { 
	ctx, cancel := context.WithCancel(ctx) // HL
	return &Group{cancel: cancel}, ctx
}
