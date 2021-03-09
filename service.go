package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translator MyTranslator
}

type MyTranslator struct {
	myTranslator Translator
}

var N = 10 //No of retries
// var N = os.Args[1]

func NewService() *Service {
	t := newRandomTranslator(
		100*time.Millisecond,
		500*time.Millisecond,
		0.8,
	)

	a := &MyTranslator{
		myTranslator: t,
	}

	return &Service{
		translator: *a,
	}
}

// Mocking the external translation
func (m *MyTranslator) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	var res string
	var err error
	var duration = 2
	for N > 0 {
		res, err = m.myTranslator.Translate(ctx, from, to, data)
		if err != nil {
			N--
			waitTime := time.Duration(duration) * time.Millisecond
			time.Sleep(waitTime)
			fmt.Println("Error Occured, Waiting after ", waitTime)
			duration = duration * 2
			continue
		} else {
			return res, err
		}
	}
	return res, err
}
