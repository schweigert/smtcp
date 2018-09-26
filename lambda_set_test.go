package smtcp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaSetSetAndGet(t *testing.T) {
	lambdaOne := func(r *Request) {}
	lambdaTwo := func(r *Request) {}

	lambdaSet := NewLambdaSet().Set("lambdaOne", lambdaOne).Set("lambdaTwo", lambdaTwo)

	assert.NotNil(t, lambdaSet.Get("lambdaOne"))
	assert.NotNil(t, lambdaSet.Get("lambdaTwo"))

	assert.Equal(t, reflect.ValueOf(lambdaOne), reflect.ValueOf(lambdaSet.Get("lambdaOne")))
	assert.Equal(t, reflect.ValueOf(lambdaTwo), reflect.ValueOf(lambdaSet.Get("lambdaTwo")))
}
