package main

import "github.com/golang/oauth2"

func oauth() {
	c := &http.Client{Transport: &mockTransport{}}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c)
	conf := &oauth2.Config{ /* ... */ }
	conf.Exchange(ctx, "code")
}
