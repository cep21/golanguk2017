package http

func (req *Request) write(w io.Writer, usingProxy bool, extraHeaders Header, waitForContinue func() bool) (err error) {
	// ...
	trace := httptrace.ContextClientTrace(req.Context())
	// ...
	if trace != nil && trace.WroteHeaders != nil {
		trace.WroteHeaders()
	}
}
