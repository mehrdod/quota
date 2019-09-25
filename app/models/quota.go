package models

import (
	"errors"
	"sort"
	"time"
)

// Quota - just a simple struct
type Quota struct {
	Id       int    `json:"id"`
	Quota    string `json:"quota"`
	Author   string `json:"author"`
	Category string `json:"category"`

	CreatedAt time.Time `json:"-"`
}

func CreateQuota(quota *Quota) (id int) {
	quotaIdSeq++
	quota.Id = quotaIdSeq
	quotaArr = append(quotaArr, *quota)

	return quotaIdSeq
}

func UpdateQuota(quota *Quota) (err error) {
	// We can do binary search because slice is sorted.
	// binary search is much more faster than range
	i := sort.Search(len(quotaArr), func(i int) bool { return quotaArr[i].Id >= quota.Id })

	if i != -1 && quotaArr[i].Id == quota.Id {
		quotaArr[i].Author = quota.Author
		quotaArr[i].Quota = quota.Quota
		quotaArr[i].Category = quota.Category

	} else {
		err = errors.New(`quota not found`)
	}
	return

}

func RemoveQuota(id int) (err error) {
	// sorted slice, that's why binary search
	i := sort.Search(len(quotaArr), func(i int) bool { return quotaArr[i].Id >= id })

	if i != -1 && quotaArr[i].Id == id {
		quotaArr = append(quotaArr[:i], quotaArr[i+1:]...)
	} else {
		err = errors.New(`quota not found`)
	}
	return

}
func GetAllQuotas() []Quota {
	return quotaArr
}

func GetQuotaByIndx(i int) Quota {
	return quotaArr[i]
}

func GetQuotaLen() int {
	return len(quotaArr)
}

func RemoveOldQuota(time time.Time) {
	// find the first one by binary search which is right after time and remove all of them that was before it.
	// binary search is much more faster than range
	i := sort.Search(len(quotaArr), func(i int) bool { return quotaArr[i].CreatedAt.After(time) })
	if i >= len(quotaArr) {
		quotaArr = quotaArr[:0]
	} else {
		if i >= 0 && i < len(quotaArr) {
			quotaArr = quotaArr[i:]
		}
	}

}
