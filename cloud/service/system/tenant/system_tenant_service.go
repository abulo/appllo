package tenant

import (
	"cloud/code"
	"cloud/module/system/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// system_tenant 租户

// SrvSystemTenantServiceServer 租户
type SrvSystemTenantServiceServer struct {
	UnimplementedSystemTenantServiceServer
	Server *xgrpc.Server
}

// SystemTenantCreate 创建数据
func (srv SrvSystemTenantServiceServer) SystemTenantCreate(ctx context.Context, request *SystemTenantCreateRequest) (*SystemTenantCreateResponse, error) {
	req := SystemTenantDao(request.GetData())
	data, err := tenant.SystemTenantCreate(ctx, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantCreate")
		return &SystemTenantCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SystemTenantUpdate 更新数据
func (srv SrvSystemTenantServiceServer) SystemTenantUpdate(ctx context.Context, request *SystemTenantUpdateRequest) (*SystemTenantUpdateResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SystemTenantDao(request.GetData())
	_, err := tenant.SystemTenantUpdate(ctx, systemTenantId, *req)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantUpdate")
		return &SystemTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenantDelete 删除数据
func (srv SrvSystemTenantServiceServer) SystemTenantDelete(ctx context.Context, request *SystemTenantDeleteRequest) (*SystemTenantDeleteResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantDelete(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantDelete")
		return &SystemTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SystemTenant 查询单条数据
func (srv SrvSystemTenantServiceServer) SystemTenant(ctx context.Context, request *SystemTenantRequest) (*SystemTenantResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.SystemTenant(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenant")
		return &SystemTenantResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SystemTenantProto(res),
	}, nil
}

// SystemTenantRecover 恢复数据
func (srv SrvSystemTenantServiceServer) SystemTenantRecover(ctx context.Context, request *SystemTenantRecoverRequest) (*SystemTenantRecoverResponse, error) {
	systemTenantId := request.GetSystemTenantId()
	if systemTenantId < 1 {
		return &SystemTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SystemTenantRecover(ctx, systemTenantId)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": systemTenantId,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantRecover")
		return &SystemTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}
func (srv SrvSystemTenantServiceServer) SystemTenantList(ctx context.Context, request *SystemTenantListRequest) (*SystemTenantListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetBeginExpireDate()))
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetFinishExpireDate()))
	}
	if request.SystemTenantPackageId != nil {
		condition["systemTenantPackageId"] = request.GetSystemTenantPackageId()
	}

	// 当前页面
	pageNum := request.GetPageNum()
	// 每页多少数据
	pageSize := request.GetPageSize()
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	// 分页数据
	offset := pageSize * (pageNum - 1)
	condition["offset"] = offset
	condition["limit"] = pageSize
	// 获取数据集合
	list, err := tenant.SystemTenantList(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantList")
		return &SystemTenantListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SystemTenantObject
	for _, item := range list {
		res = append(res, SystemTenantProto(item))
	}
	return &SystemTenantListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SystemTenantListTotal 获取总数
func (srv SrvSystemTenantServiceServer) SystemTenantListTotal(ctx context.Context, request *SystemTenantListTotalRequest) (*SystemTenantTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetBeginExpireDate()))
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetFinishExpireDate()))
	}
	if request.SystemTenantPackageId != nil {
		condition["systemTenantPackageId"] = request.GetSystemTenantPackageId()
	}

	// 获取数据集合
	total, err := tenant.SystemTenantListTotal(ctx, condition)
	if sql.ResultAccept(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:system_tenant:SystemTenantListTotal")
		return &SystemTenantTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SystemTenantTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
