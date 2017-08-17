package main

func DoTwoRequestsAtOnce(ctx context.Context) error {
	eg, egCtx := errgroup.WithContext(ctx) // HL
	var resp1, resp2 *http.Response
	f := func(loc string, respIn **http.Response) func() error {
		return func() error {
			reqCtx, cancel := context.WithTimeout(egCtx, time.Second) // HL
			defer cancel()
			req, _ := http.NewRequest("GET", loc, nil) // TODO: Check this error
			var err error
			*resp, err = http.DefaultClient.Do(req.WithContext(reqCtx)) // HL
			if err == nil && (*respIn).StatusCode >= 500 {
				return errors.New("unexpected!")
			}
			return err
		}
	}

	eg.Go(f("http://localhost:8080/fast_request", &resp1)) // <--- Run two requests
	eg.Go(f("http://localhost:8080/slow_request", &resp2))

	return eg.Wait() // TODO: Actually do something with resp1 and resp2
}
