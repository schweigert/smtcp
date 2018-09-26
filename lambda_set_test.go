package smtcp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaSetSetAndGet(t *testing.T) {
	anyFlag := false

	lambdaOne := func(r *Request) {
		anyFlag = true
	}

	lambdaTwo := func(r *Request) {}

	lambdaSet := NewLambdaSet().Set("lambdaOne", lambdaOne).Set("lambdaTwo", lambdaTwo)

	assert.NotNil(t, lambdaSet.Get("lambdaOne"))
	assert.NotNil(t, lambdaSet.Get("lambdaTwo"))
	assert.Nil(t, lambdaSet.Get("nil"))

	assert.Equal(t, reflect.ValueOf(lambdaOne), reflect.ValueOf(lambdaSet.Get("lambdaOne")))
	assert.Equal(t, reflect.ValueOf(lambdaTwo), reflect.ValueOf(lambdaSet.Get("lambdaTwo")))

	assert.False(t, anyFlag)

	r := NewRequest("lambdaOne", nil, nil)
	lambdaSet.Execute(r)

	assert.True(t, anyFlag)

	r = NewRequest("lambdaThree", nil, nil)
	lambdaSet.Execute(r)
}
