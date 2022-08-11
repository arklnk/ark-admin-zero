package utils

import (
	"context"
	"crypto/md5"
	"fmt"
	"strconv"

	"ark-zero-admin/pkg/sysconstant"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func UserId(ctx context.Context) int64 {
	id := ctx.Value(sysconstant.JwtUserId)
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", id))
	return int64(userId)
}
