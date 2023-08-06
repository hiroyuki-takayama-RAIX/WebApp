package main

func main() {
	exampleWithCancel()
	exampleDone()
	processUntilCancelled()
	sampleString()
	sampleString2()
}

/*
func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("Not Cancelled")
}

func exampleWithCancel() {
	// context.WithCancel()
	// context.Background() generates top level context as root, which is not canceled, has no value and no deadline.
	// ctx is context.Context interface.
	// cancel() is context.CancelFunc(), which give ctx cocanceled value.
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}

func exampleDone() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("another gorutine!")
	}()
	fmt.Println("stop")
	<-ctx.Done()
	fmt.Println("time is running!")
}

func processUntilCancelled() {

}

func sampleString() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 30; j++ {
			fmt.Print(i)
		}
	}
}

func sampleString2() {
	s := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	for i := 0; i < 10; i++ {
		for j := 0; j < 30; j++ {
			fmt.Print(s[i])
		}
	}
}
*/
