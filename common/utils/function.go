package utils

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"ark-zero-admin/common/globalkey"
	"github.com/zeromicro/go-zero/core/logx"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GetUserId(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(globalkey.JwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}

	return uid
}

func ArrayUniqueValue[T any](arr []T) []T {
	size := len(arr)
	result := make([]T, 0, size)
	temp := map[any]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}

	return result
}

func ArrayContainValue(arr []int64, search int64) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}

	return false
}

func StrToTime(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}
	return date
}
