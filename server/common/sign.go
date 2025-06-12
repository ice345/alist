package common

import (
	stdpath "path"

	"github.com/NewAlist/alist/internal/conf"
	"github.com/NewAlist/alist/internal/model"
	"github.com/NewAlist/alist/internal/setting"
	"github.com/NewAlist/alist/internal/sign"
)

func Sign(obj model.Obj, parent string, encrypt bool) string {
	if obj.IsDir() || (!encrypt && !setting.GetBool(conf.SignAll)) {
		return ""
	}
	return sign.Sign(stdpath.Join(parent, obj.GetName()))
}
