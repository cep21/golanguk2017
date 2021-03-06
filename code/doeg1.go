func (g *Group) Go(f func() error) {     // HL
	g.wg.Add(1)
	go func() {                      // HL
		defer g.wg.Done()
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
