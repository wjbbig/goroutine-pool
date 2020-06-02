# goroutine-pool
A simple implementation of goroutine pool

case:

~~~go
pool := NewPool(3)

task := NewTask(func() error {
    fmt.Println("time now--->", time.Now().Format("2006-01-02 15:04:05"))
    return nil
})

go func() {
    for {
        time.Sleep(time.Second)
        pool.EntranceChan <- task
    }
}()

pool.run()
~~~

