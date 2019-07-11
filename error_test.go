// Copyright 2019 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package axios

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert := assert.New(t)
	err := CreateError(errors.New("abcd"), nil, 500)

	assert.Equal("code=500, message=abcd", err.Error())
	assert.Equal("code=500, message=abcd", fmt.Sprintf("%v", err))
	assert.Equal("code=500, message=abcd", fmt.Sprintf("%s", err))
	assert.Equal(`"code=500, message=abcd"`, fmt.Sprintf("%q", err))

	err = CreateError(err, nil, 400)
	assert.Equal("code=500, message=abcd", err.Error())
	assert.Equal("code=500, message=abcd", fmt.Sprintf("%v", err))
	assert.Equal("code=500, message=abcd", fmt.Sprintf("%s", err))
	assert.Equal(`"code=500, message=abcd"`, fmt.Sprintf("%q", err))
}
