package main

import (
    "fmt"
    "byte"
    "time"
    "strconv"
)

func see(s string) {
    a := time.Now()
    for i := 0; i < 1000; i++ {
        var sum int
        sum = 1
        sum = sum * i
        fmt.Println("sum:", sum)
        fmt.Println("s:", s)
        getid("see " + s)
    }
    fmt.Println(time.Since(a))
}

//func getid(s string) {
	//pid := unix.Getpid()
	//fmt.Println(s+" Getpid", pid)
	//fmt.Println(s+" Getppid", unix.Getppid())
	//pgid, _ := unix.Getpgid(pid)
	//fmt.Println(s+" Getpgid", pgid)
	//fmt.Println(s+" Gettid", unix.Gettid())
	//sid, _ := unix.Getsid(pid)
	//fmt.Println(s+" Getsid", sid)
	//fmt.Println(s+" Getegid", unix.Getegid())
	//fmt.Println(s+" Geteuid", unix.Geteuid())
	//fmt.Println(s+" Getgid", unix.Getgid())
	//fmt.Println(s+" Getuid", unix.Getuid())
//}

func getid() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}

func main() {
    go see("hello")
    go see("world")
    getid("main")
    time.Sleep(20 * time.Second)
}
