// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package file

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *NewMultiUploadReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewMultiUploadReq[number], err)
}

func (x *NewMultiUploadReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileHash, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ChunkTotalNumber, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FileSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.FileName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.UserID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.FileType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *NewMultiUploadResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *GetMultiUploadUriReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMultiUploadUriReq[number], err)
}

func (x *GetMultiUploadUriReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileHash, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMultiUploadUriReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMultiUploadUriReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ChunkID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMultiUploadUriReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ChunkSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMultiUploadUriResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMultiUploadUriResp[number], err)
}

func (x *GetMultiUploadUriResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uri, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CompleteMultipartReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CompleteMultipartReq[number], err)
}

func (x *CompleteMultipartReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileHash, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CompleteMultipartReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CompleteMultipartResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *GetSuccessChunksReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetSuccessChunksReq[number], err)
}

func (x *GetSuccessChunksReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FileHash, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetSuccessChunksReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetSuccessChunksResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetSuccessChunksResp[number], err)
}

func (x *GetSuccessChunksResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsUpload, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetSuccessChunksResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Chunks, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewMultiUploadReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *NewMultiUploadReq) fastWriteField1(buf []byte) (offset int) {
	if x.FileHash == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileHash())
	return offset
}

func (x *NewMultiUploadReq) fastWriteField2(buf []byte) (offset int) {
	if x.ChunkTotalNumber == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetChunkTotalNumber())
	return offset
}

func (x *NewMultiUploadReq) fastWriteField3(buf []byte) (offset int) {
	if x.FileSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFileSize())
	return offset
}

func (x *NewMultiUploadReq) fastWriteField4(buf []byte) (offset int) {
	if x.FileName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetFileName())
	return offset
}

func (x *NewMultiUploadReq) fastWriteField5(buf []byte) (offset int) {
	if x.UserID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetUserID())
	return offset
}

func (x *NewMultiUploadReq) fastWriteField6(buf []byte) (offset int) {
	if x.FileType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetFileType())
	return offset
}

func (x *NewMultiUploadResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetMultiUploadUriReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetMultiUploadUriReq) fastWriteField1(buf []byte) (offset int) {
	if x.FileHash == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileHash())
	return offset
}

func (x *GetMultiUploadUriReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserID())
	return offset
}

func (x *GetMultiUploadUriReq) fastWriteField3(buf []byte) (offset int) {
	if x.ChunkID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetChunkID())
	return offset
}

func (x *GetMultiUploadUriReq) fastWriteField4(buf []byte) (offset int) {
	if x.ChunkSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetChunkSize())
	return offset
}

func (x *GetMultiUploadUriResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetMultiUploadUriResp) fastWriteField1(buf []byte) (offset int) {
	if x.Uri == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUri())
	return offset
}

func (x *CompleteMultipartReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CompleteMultipartReq) fastWriteField1(buf []byte) (offset int) {
	if x.FileHash == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileHash())
	return offset
}

func (x *CompleteMultipartReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserID())
	return offset
}

func (x *CompleteMultipartResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetSuccessChunksReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetSuccessChunksReq) fastWriteField1(buf []byte) (offset int) {
	if x.FileHash == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFileHash())
	return offset
}

func (x *GetSuccessChunksReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserID())
	return offset
}

func (x *GetSuccessChunksResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetSuccessChunksResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsUpload {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsUpload())
	return offset
}

func (x *GetSuccessChunksResp) fastWriteField2(buf []byte) (offset int) {
	if x.Chunks == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetChunks())
	return offset
}

func (x *NewMultiUploadReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *NewMultiUploadReq) sizeField1() (n int) {
	if x.FileHash == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileHash())
	return n
}

func (x *NewMultiUploadReq) sizeField2() (n int) {
	if x.ChunkTotalNumber == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetChunkTotalNumber())
	return n
}

func (x *NewMultiUploadReq) sizeField3() (n int) {
	if x.FileSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFileSize())
	return n
}

func (x *NewMultiUploadReq) sizeField4() (n int) {
	if x.FileName == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetFileName())
	return n
}

func (x *NewMultiUploadReq) sizeField5() (n int) {
	if x.UserID == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetUserID())
	return n
}

func (x *NewMultiUploadReq) sizeField6() (n int) {
	if x.FileType == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetFileType())
	return n
}

func (x *NewMultiUploadResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetMultiUploadUriReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetMultiUploadUriReq) sizeField1() (n int) {
	if x.FileHash == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileHash())
	return n
}

func (x *GetMultiUploadUriReq) sizeField2() (n int) {
	if x.UserID == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserID())
	return n
}

func (x *GetMultiUploadUriReq) sizeField3() (n int) {
	if x.ChunkID == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetChunkID())
	return n
}

func (x *GetMultiUploadUriReq) sizeField4() (n int) {
	if x.ChunkSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetChunkSize())
	return n
}

func (x *GetMultiUploadUriResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetMultiUploadUriResp) sizeField1() (n int) {
	if x.Uri == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUri())
	return n
}

func (x *CompleteMultipartReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CompleteMultipartReq) sizeField1() (n int) {
	if x.FileHash == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileHash())
	return n
}

func (x *CompleteMultipartReq) sizeField2() (n int) {
	if x.UserID == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserID())
	return n
}

func (x *CompleteMultipartResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetSuccessChunksReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetSuccessChunksReq) sizeField1() (n int) {
	if x.FileHash == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFileHash())
	return n
}

func (x *GetSuccessChunksReq) sizeField2() (n int) {
	if x.UserID == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserID())
	return n
}

func (x *GetSuccessChunksResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetSuccessChunksResp) sizeField1() (n int) {
	if !x.IsUpload {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsUpload())
	return n
}

func (x *GetSuccessChunksResp) sizeField2() (n int) {
	if x.Chunks == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetChunks())
	return n
}

var fieldIDToName_NewMultiUploadReq = map[int32]string{
	1: "FileHash",
	2: "ChunkTotalNumber",
	3: "FileSize",
	4: "FileName",
	5: "UserID",
	6: "FileType",
}

var fieldIDToName_NewMultiUploadResp = map[int32]string{}

var fieldIDToName_GetMultiUploadUriReq = map[int32]string{
	1: "FileHash",
	2: "UserID",
	3: "ChunkID",
	4: "ChunkSize",
}

var fieldIDToName_GetMultiUploadUriResp = map[int32]string{
	1: "Uri",
}

var fieldIDToName_CompleteMultipartReq = map[int32]string{
	1: "FileHash",
	2: "UserID",
}

var fieldIDToName_CompleteMultipartResp = map[int32]string{}

var fieldIDToName_GetSuccessChunksReq = map[int32]string{
	1: "FileHash",
	2: "UserID",
}

var fieldIDToName_GetSuccessChunksResp = map[int32]string{
	1: "IsUpload",
	2: "Chunks",
}