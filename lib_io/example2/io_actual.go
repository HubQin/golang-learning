package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func DownFileWithReadAll() {
	url := "https://file-examples.com/fe-0129qp/2017/04/file_example_MP4_640_3MG.mp4"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()
	fmt.Println("down with ReadAll...")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("start write...")

	_ = ioutil.WriteFile("./file_example_MP4_640_3MG.mp4", data, 0755)
}

func DownFileWithCopy() {
	url := "https://file-examples.com/fe-0129qp/2017/04/file_example_MP4_640_3MG.mp4"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()

	out, err := os.Create("./file_example_MP4_640_3MG-2.mp4")
	wt := bufio.NewWriter(out)

	defer out.Close()
	fmt.Println("down with Copy...")
	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		fmt.Println(err)
	}
	wt.Flush()
}

func main() {
	DownFileWithReadAll()
	DownFileWithCopy()
}
