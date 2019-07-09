package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestMVC(t *testing.T) {
	e := httptest.New(t, newApp())

	var wg sync.WaitGroup
	e.Request("GET", "/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前参与用户： 0\n")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			e.POST("/import").WithFormField("users",
				fmt.Sprintf("test_u%d", i)).Expect().
				Status(httptest.StatusOK)
		}(i)
	}

	wg.Wait()

	e.Request("GET", "/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前参与用户： 100\n")

	e.Request("GET", "/lucky").Expect().Status(httptest.StatusOK)

	e.Request("GET", "/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前参与用户： 99\n")
}
