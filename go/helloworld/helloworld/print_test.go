package helloworld_test

import (
	"bytes"
	"testing"

	"github.com/KoichiWada/playbox/helloworld/helloworld"
	"github.com/stretchr/testify/assert"
)

func TestPrintln(t *testing.T) {
	assert := assert.New(t)

	want := "こんにちは、世界！\n"

	buf := new(bytes.Buffer)
	helloworld.Println(buf)
	assert.Equal(want, buf.String())
}
