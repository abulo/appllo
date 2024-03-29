import { PORT } from "@/api/config/servicePort";
import http from "@/api";
import { SystemRole } from "@/api/interface/systemRole";

// 获取角色列表
export const getSystemRoleListApi = (params?: SystemRole.ReqSystemRoleList) => {
  return http.get<SystemRole.ResSystemRoleItem[]>(PORT + `/system/role`, params);
};

// 获取单个角色
export const getSystemRoleItemApi = (id: number) => {
  return http.get<SystemRole.ResSystemRoleItem>(PORT + `/system/role/${id}/item`);
};

// 添加角色
export const addSystemRoleApi = (params: SystemRole.ResSystemRoleItem) => {
  return http.post(PORT + `/system/role`, params);
};

// 修改角色
export const updateSystemRoleApi = (id: number, params: SystemRole.ResSystemRoleItem) => {
  return http.put(PORT + `/system/role/${id}/update`, params);
};

// 删除角色
export const deleteSystemRoleApi = (id: number) => {
  return http.delete(PORT + `/system/role/${id}/delete`);
};

// 恢复角色
export const recoverSystemRoleApi = (id: number) => {
  return http.put(PORT + `/system/role/${id}/recover`);
};
// 菜单
export const getSystemRoleMenuListApi = (id: number, params?: SystemRole.ReqSystemRoleMenuItem) => {
  return http.get<SystemRole.ResSystemRoleMenuItem>(PORT + `/system/role/${id}/menu`, params);
};
// 添加菜单
export const addSystemRoleMenuApi = (id: number, params: SystemRole.ResSystemRoleMenuItem) => {
  return http.post(PORT + `/system/role/${id}/menu`, params);
};

// 数据范围
export const getSystemRoleScopeListApi = (id: number) => {
  return http.get<SystemRole.ResSystemRoleScopeItem>(PORT + `/system/role/${id}/scope`);
};
// 添加数据范围
export const addSystemRoleScopeApi = (id: number, params: SystemRole.ResSystemRoleScopeItem) => {
  return http.post(PORT + `/system/role/${id}/scope`, params);
};
// 获取角色列表
export const getSystemRoleSearchApi = (params?: SystemRole.ReqSystemRoleList) => {
  return http.get<SystemRole.ResSystemRoleItem[]>(PORT + `/system/role/search`, params);
};
