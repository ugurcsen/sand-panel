package docker

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestComposeYamlBuilder(t *testing.T) {
	tf, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tf.Name())

	if err = composeYamlBuilder(testCollection, tf); err != nil {
		t.Fatal(err)
	}
	tf.Seek(0, 0)
	a, err := io.ReadAll(tf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println()
	for _, b := range a {
		fmt.Printf("0x%x ,", b)
	}
	fmt.Println()

	cmpValue := []byte{0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x22, 0x33, 0x2e, 0x37, 0x22, 0xa, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x74, 0x65, 0x73, 0x74, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x20, 0x74, 0x65, 0x73, 0x74, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x20, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x20, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x3a, 0x20, 0x77, 0x65, 0x62, 0x2c, 0x20, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x3a, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x28, 0x60, 0x74, 0x65, 0x73, 0x74, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x60, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x28, 0x60, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x60, 0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x3a, 0x20, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x20, 0x22, 0x38, 0x30, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x20, 0x72, 0x75, 0x6e, 0x73, 0x63, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x2e, 0x2f, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3a, 0x2f, 0x75, 0x73, 0x72, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0xa, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0xa, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x74, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x6e, 0x65, 0x74, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0xa}
	if bytes.Compare(a, cmpValue) != 0 {
		t.Error("yaml not created correctly")
	}
	tf.Close()
}
