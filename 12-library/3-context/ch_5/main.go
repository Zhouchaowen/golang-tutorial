package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	Ip   string
	Port int
)

func init() {
	flag.StringVar(&Ip, "i", "0.0.0.0", "server ip")
	flag.IntVar(&Port, "p", 8080, "server port")
}

func doSomeA() {
	// do something A....
	time.Sleep(3 * time.Second)
	fmt.Println("taskA do something A")

	// do something B....
	time.Sleep(3 * time.Second)
	fmt.Println("taskA do something B")

	// do something C....
	time.Sleep(3 * time.Second)
	fmt.Println("taskA do something C")
}

func doSomeB(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("exit task, reason:", ctx.Err())
		return
	default:
	}

	// do something A....
	time.Sleep(3 * time.Second)
	fmt.Println("taskB do something A")

	select {
	case <-ctx.Done():
		fmt.Println("exit task, reason:", ctx.Err())
		return
	default:
	}

	// do something B....
	time.Sleep(3 * time.Second)
	fmt.Println("taskB do something B")

	select {
	case <-ctx.Done():
		fmt.Println("exit task, reason:", ctx.Err())
		return
	default:
	}

	// do something C....
	time.Sleep(3 * time.Second)
	fmt.Println("taskB do something C")
}

func server() {
	http.HandleFunc("/noTimeoutControl", func(w http.ResponseWriter, r *http.Request) {
		// do something....
		go doSomeA()

		fmt.Fprintf(w, "Golang Tutorial!")
	})

	http.HandleFunc("/timeoutControl", func(w http.ResponseWriter, r *http.Request) {
		cxt, _ := context.WithTimeout(context.TODO(), 2*time.Second)

		// do something....
		doSomeB(cxt)

		fmt.Fprintf(w, "Golang Tutorial!")
	})

	http.HandleFunc("/httpTimeoutControl", func(w http.ResponseWriter, r *http.Request) {
		// do something....
		doSomeB(r.Context())

		fmt.Fprintf(w, "Golang Tutorial!")
	})

	// 启动服务端, 监听的地址 IP:Port
	http.ListenAndServe(fmt.Sprintf("%s:%d", Ip, Port), nil)
}

func Steps1() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, _ := http.NewRequest("GET", "http://0.0.0.0:8080/noTimeoutControl", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("response:", string(body))
}

func Steps2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, _ := http.NewRequest("GET", "http://0.0.0.0:8080/timeoutControl", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("response:", string(body))
}

func Steps3() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, _ := http.NewRequest("GET", "http://0.0.0.0:8080/timeoutControl", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("response:", string(body))
}

func main() {
	flag.Parse()

	go server()

	// 等待服务启动
	<-time.After(5 * time.Second)

	//Steps1()
	//Steps2()
	Steps3()

	<-time.After(10 * time.Second)
}
