const getters = {
  sidebar: state => state.app.sidebar,
  language: state => state.app.language,
  size: state => state.app.size,
  device: state => state.app.device,
  userMenuList: state => state.router.userMenuList,
  perms: state => state.user.perms
}
export default getters
