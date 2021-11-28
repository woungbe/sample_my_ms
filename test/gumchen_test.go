package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"sample_my_ms/model/pos"
)

func TestFoo(t *testing.T) {

	expected := 1
    actual := 0

	rows, err := pos.GetProductList("");
	fmt.Println(rows, err);


    assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")

}