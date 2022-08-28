package utils

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"

	"ark-admin-zero/config"

	"github.com/zeromicro/go-zero/core/logx"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GetUserId(ctx context.Context) uint64 {
	var uid uint64
	if jsonUid, ok := ctx.Value(config.SysJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = uint64(int64Uid)
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

func ArrayContainValue(arr []uint64, search uint64) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}

	return false
}

func Intersect(slice1 []uint64, slice2 []uint64) []uint64 {
	m := make(map[uint64]uint64)
	n := make([]uint64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}

	return n
}

func Difference(slice1 []uint64, slice2 []uint64) []uint64 {
	m := make(map[uint64]int)
	n := make([]uint64, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, v := range slice1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}

	return n
}
