package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	fmt.Println(time.Now())
	DemoFeed, err := Feed(1, time.Unix(1691394601894, 0))
	assert.Equal(t, nil, err)
	fmt.Printf("%#v", DemoFeed)
	time.Sleep(1 * time.Second)
}
