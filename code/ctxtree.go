package main

func tree() {
	ctx1 := context.Background()
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, _ := context.WithTimeout(ctx2, time.Second*5)
	ctx4, _ := context.WithTimeout(ctx3, time.Second*3)
	ctx5, _ := context.WithTimeout(ctx3, time.Second*6)
	ctx6 := context.WithValue(ctx5, "UserID", 12)
}
