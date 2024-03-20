package tenant

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/service/system/tenant"

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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// system_tenant 租户
// SystemTenantCreate 创建数据
func SystemTenantCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantCreate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	request := &tenant.SystemTenantCreateRequest{}
	// 数据绑定
	var reqInfo dao.SystemTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Deleted = proto.Int32(0)
	reqInfo.Creator = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	if reqInfo.Password != nil {
		reqInfo.Password = proto.String(util.Md5(cast.ToString(reqInfo.Password)))
	}
	request.Data = tenant.SystemTenantProto(reqInfo)
	// 执行服务
	res, err := client.SystemTenantCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantCreate")
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

// SystemTenantUpdate 更新数据
func SystemTenantUpdate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantUpdate")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantUpdateRequest{}
	request.SystemTenantId = systemTenantId
	// 数据绑定
	var reqInfo dao.SystemTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Updater = null.StringFrom(newCtx.GetString("systemUserName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	request.Data = tenant.SystemTenantProto(reqInfo)
	// 执行服务
	res, err := client.SystemTenantUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantUpdate")
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

// SystemTenantDelete 删除数据
func SystemTenantDelete(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantDelete")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantDeleteRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenantDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantDelete")
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

// SystemTenant 查询单条数据
func SystemTenant(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenant")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenant(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenant")
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
		"data": tenant.SystemTenantDao(res.GetData()),
	})
}

// SystemTenantRecover 恢复数据
func SystemTenantRecover(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantRecover")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	systemTenantId := cast.ToInt64(newCtx.Param("systemTenantId"))
	request := &tenant.SystemTenantRecoverRequest{}
	request.SystemTenantId = systemTenantId
	// 执行服务
	res, err := client.SystemTenantRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantRecover")
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

// SystemTenantList 列表数据
func SystemTenantSearch(ctx context.Context, newCtx *app.RequestContext) {
	SystemTenantList(ctx, newCtx)
}

// SystemTenantList 列表数据
func SystemTenantList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:system_tenant:SystemTenantList")
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSystemTenantServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SystemTenantListRequest{}
	requestTotal := &tenant.SystemTenantListTotalRequest{}

	request.Deleted = proto.Int32(0)      // 删除状态
	requestTotal.Deleted = proto.Int32(0) // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("systemTenantPackageId"); ok {
		request.SystemTenantPackageId = proto.Int64(cast.ToInt64(val))      // 套餐编号
		requestTotal.SystemTenantPackageId = proto.Int64(cast.ToInt64(val)) // 套餐编号
	}
	if val, ok := newCtx.GetQuery("deleted"); ok {
		request.Deleted = proto.Int32(cast.ToInt32(val))      // 是否删除(0否 1是)
		requestTotal.Deleted = proto.Int32(cast.ToInt32(val)) // 是否删除(0否 1是)
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("beginExpireDate"); ok {
		request.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("finishExpireDate"); ok {
		request.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}

	// 执行服务,获取数据量
	resTotal, err := client.SystemTenantListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	request.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	request.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SystemTenantList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:system_tenant:SystemTenantList")
		fromError := status.Convert(err)
		newCtx.JSON(consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SystemTenant
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, tenant.SystemTenantDao(item))
		}
	}
	newCtx.JSON(consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  request.PageNum,
			"pageSize": request.PageSize,
		},
	})
}
