// Create a group and a context to use with that group
func WithContext(ctx context.Context) (*Group, context.Context) {
	// ...
}

// Go executes f in another goroutine
func (g *Group) Go(f func() error) {
	// ...
}

// Wait for all Go functions to finish
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel() // HL
	}
	return g.err
}
