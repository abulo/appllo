package role

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/role"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// SystemRoleMenuCreate 创建数据
func SystemRoleMenuCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleMenuCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSystemRoleMenuServiceClient(grpcClient)
	request := &role.SystemRoleMenuCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemRoleMenuCustom
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	systemRoleId := cast.ToInt64(newCtx.Param("systemRoleId"))
	reqInfo.Creator = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	reqInfo.Updater = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.SystemRoleId = proto.Int64(systemRoleId)
	reqInfo.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId"))
	request.Data = role.SystemRoleMenuCustomProto(reqInfo)

	// 执行服务
	res, err := client.SystemRoleMenuCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleMenuCreate")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

func SystemRoleMenuList(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统角色:system_role:SystemRoleMenuList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	systemRoleId := cast.ToInt64(newCtx.Param("systemRoleId"))
	//链接服务
	client := role.NewSystemRoleMenuServiceClient(grpcClient)
	request := &role.SystemRoleMenuListRequest{}
	request.SystemTenantId = proto.Int64(newCtx.GetInt64("systemTenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                        // 删除状态
	request.SystemRoleId = proto.Int64(cast.ToInt64(systemRoleId))
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("systemMenuId"); ok {
		request.SystemMenuId = proto.Int64(cast.ToInt64(val))
	}
	// 执行服务
	res, err := client.SystemRoleMenuList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统角色:system_role:SystemRoleMenuList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []int64
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, *item.SystemMenuId)
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"systemRoleId":  systemRoleId,
			"systemMenuIds": list,
		},
	})
}
