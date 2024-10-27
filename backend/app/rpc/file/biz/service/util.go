package service

import "strings"

func ObjectName(basePath, uid, uuid, ext string) string {
	return strings.TrimPrefix(basePath+"/"+uid+"/"+uuid+"."+ext, "/")
}
