// Code generated by hertz generator. DO NOT EDIT.

package file

import (
	file "backend/app/http/biz/handler/file"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_auth := _v1.Group("/auth", _authMw()...)
			{
				_file := _auth.Group("/file", _fileMw()...)
				_file.POST("/complete_multipart", append(_completemultipartMw(), file.CompleteMultipart)...)
				_file.POST("/multi_upload", append(_newmultiuploadMw(), file.NewMultiUpload)...)
				_file.POST("/multi_upload_uri", append(_getmultiuploaduriMw(), file.GetMultiUploadUri)...)
				_file.POST("/success_chunks", append(_getsuccesschunksMw(), file.GetSuccessChunks)...)
			}
		}
	}
}