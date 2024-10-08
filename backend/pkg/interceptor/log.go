package interceptor

import (
	lang2 "backend/pkg/lang"
	"backend/pkg/xerror"
	"backend/pkg/xvalidator"
	"context"
	"github.com/dtm-labs/dtmcli"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ServerLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	var lang string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		lang = "en"
	} else {
		langArray := md.Get("lang")
		if len(langArray) == 0 {
			lang = "en"
		} else {
			lang = langArray[0]
		}
	}
	if err != nil {
		if errors.Is(err, status.Error(codes.Aborted, dtmcli.ResultFailure)) {
			// 分布式事务错误不转化
			return resp, err
		} else {
			causeErr := errors.Cause(err)
			// log error
			switch e := causeErr.(type) {
			case xerror.XError:
				//switch e {
				//case xerror.ServerError:
				//	code := codes.Internal
				//	message := localizeError(lang, e)
				//	newe :=  e.SetTranslateMessage(message)
				//	logx.WithContext(ctx).Errorf("[rpc] %+v", err)
				//	err = status.Error(codes.Code(uint32(code)), newe.Error())
				//	return resp, err
				//}
				code := e.Code
				message := localizeError(lang, e)
				logx.WithContext(ctx).Errorf("[rpc] %+v", err)
				err = status.Error(codes.Code(uint32(code)), message)
				return resp, err
			default:
				logx.WithContext(ctx).Errorf("[rpc] %+v", err)
				return resp, err
			}

		}

	}
	return resp, nil
}

func localizeError(lang string, e xerror.XError) string {
	localizer := i18n.NewLocalizer(lang2.GetBundle(), lang)
	errDetail, nerr := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &e.Message,
		TemplateData:   e.TemplateData,
	})
	if nerr != nil {
		return xvalidator.Translate(nerr)
	}
	return errDetail
}
