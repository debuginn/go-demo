package main

import "fmt"

type session interface {
	add()
	del()
	close()
}

type rpcSession struct {
	session session
	reqNum  int64
}

func main() {

	trueSession := make([]*rpcSession, 0, 0)
	trueSession = append(trueSession, &rpcSession{})

	a := make([]*rpcSession, 0, 0)
	a = append(a, trueSession...)

	for _, sess := range a {

		fmt.Printf("====== ")

		sess.session.close()
	}

}
