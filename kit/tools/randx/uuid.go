package randx

import (
	"github.com/google/uuid"
)

// UUID 生成小写的36位UUID字符串 935ef64c-284c-4170-ac2f-f2a070464d5a
func UUID() string {
	return uuid.NewString()
}
