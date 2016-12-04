package api

/************************************************
  generated by IDE. for [AutoBackupAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件リセット
func (api *AutoBackupAPI) Reset() *AutoBackupAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *AutoBackupAPI) Offset(offset int) *AutoBackupAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *AutoBackupAPI) Limit(limit int) *AutoBackupAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *AutoBackupAPI) Include(key string) *AutoBackupAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *AutoBackupAPI) Exclude(key string) *AutoBackupAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルタ
func (api *AutoBackupAPI) FilterBy(key string, value interface{}) *AutoBackupAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *AutoBackupAPI) FilterMultiBy(key string, value interface{}) *AutoBackupAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

// WithNameLike 名称条件
func (api *AutoBackupAPI) WithNameLike(name string) *AutoBackupAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *AutoBackupAPI) WithTag(tag string) *AutoBackupAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *AutoBackupAPI) WithTags(tags []string) *AutoBackupAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *AutoBackupAPI) WithSizeGib(size int) *AutoBackupAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *AutoBackupAPI) WithSharedScope() *AutoBackupAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *AutoBackupAPI) WithUserScope() *AutoBackupAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *AutoBackupAPI) SortBy(key string, reverse bool) *AutoBackupAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名前でのソート
func (api *AutoBackupAPI) SortByName(reverse bool) *AutoBackupAPI {
	api.sortByName(reverse)
	return api
}

// func (api *AutoBackupAPI) SortBySize(reverse bool) *AutoBackupAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *AutoBackupAPI) New() *sacloud.AutoBackup {
// 	return &sacloud.AutoBackup{}
// }

// func (api *AutoBackupAPI) Create(value *sacloud.AutoBackup) (*sacloud.AutoBackup, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *AutoBackupAPI) Read(id string) (*sacloud.AutoBackup, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *AutoBackupAPI) Update(id string, value *sacloud.AutoBackup) (*sacloud.AutoBackup, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *AutoBackupAPI) Delete(id string) (*sacloud.AutoBackup, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *AutoBackupAPI) setStateValue(setFunc func(*sacloud.Request)) *AutoBackupAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *AutoBackupAPI) request(f func(*sacloud.Response) error) (*sacloud.AutoBackup, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.AutoBackup, nil
//}
//
//func (api *AutoBackupAPI) createRequest(value *sacloud.AutoBackup) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.AutoBackup = value
//	return req
//}
