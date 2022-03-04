import store from '@/store'
/**
 * 判断用户是否拥有操作权限
 * 根据传入的权限标识，查看是否存在用户权限标识集合
 * @param perms
 */
export function hasPermission(perms) {
  let hasPermission = false
  let permissionName = ''
  const permissions = store.state.user.perms
  for (var key in permissions) {
    	if (key === perms) {
    		hasPermission = true
    		permissionName = permissions[key]
      break
    	}
  }
  return {
    	show: hasPermission,
    	label: permissionName
  }
}

