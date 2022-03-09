/**
 * @Time: 2022/3/6 23:42
 * @Author: yt.yin
 */

package menusrv

import (
	"errors"
	"strings"

	"github.com/goworkeryyt/aramis/server/menu/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/uuid"
	"gorm.io/gorm"
)

type MenuService struct{}

// CreateMenu
/**
 *  @Description: 新建菜单项
 *  @receiver menuService
 *  @param menu
 *  @return err
 */
func (menuService *MenuService) CreateMenu(menu menumod.Menu) (err error) {
	var m menumod.Menu
	// 若存在父菜单，则需对父菜单进行效验
	if strings.TrimSpace(menu.ParentId) != "" {
		// 判定父菜单是否合法
		_, err = menuService.GetMenuInfo(menu.ParentId)
		if err != nil {
			return errors.New("非法的父菜单")
		}
		// 判定菜单名称是否被其他子菜单占用
		if !errors.Is(global.DB.Where("menu_name = ?", menu.MenuName).Not("parent_id = ?", "").First(&m).Error, gorm.ErrRecordNotFound) {
			return errors.New("该菜单名称已被其他子菜单占用")
		}
	} else {
		// 判定菜单名称是否被其他父菜单占用
		if !errors.Is(global.DB.Where(map[string]interface{}{"menu_name": menu.MenuName, "parent_id": nil}).Or("menu_name = ? AND parent_id = ?", menu.MenuName, "").First(&m).Error, gorm.ErrRecordNotFound) {
			return errors.New("该菜单名称已被其他父菜单占用")
		}
	}
	// 效验url
	if !errors.Is(global.DB.Where("url = ?", menu.Url).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单路径（url）不可重复")
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建菜单
	err = tx.Create(&menu).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建菜单失败")
	}

	// 绑定按钮
	// 判定按钮是否存在重复值
	var buttonNames []string
	var ButtonIdentities []string
	for _, button := range menu.Buttons {
		if strings.TrimSpace(button.ButtonName) == "" {
			return errors.New("存在非法的按钮名称")
		} else {
			buttonNames = append(buttonNames, button.ButtonName)
		}
		if strings.TrimSpace(button.ButtonIdentity) == "" {
			return errors.New("存在非法的按钮标识")
		} else {
			ButtonIdentities = append(ButtonIdentities, button.ButtonIdentity)
		}
	}
	if array.IsExistRepeatInArray(buttonNames) {
		return errors.New("按钮名称不唯一")
	}
	if array.IsExistRepeatInArray(ButtonIdentities) {
		return errors.New("按钮标识不唯一")
	}

	// 获取当前菜单的URL（递归）
	url, err := menuService.GetMenuPath(tx, menu.ID)
	if err != nil {
		return errors.New("解析URL失败")
	}

	// 有按钮才创建
	if len(menu.Buttons) > 0 {
		// 遍历得到需要创建的按钮
		var buttonCreates []menumod.Button
		for _, button := range menu.Buttons {
			button.ID = uuid.UUID()
			button.MenuId = menu.ID
			button.PageName = menu.MenuName
			button.PagePath = url
			button.ButtonFullIdentity = strings.Replace(url, "/", ":", -1) + ":" + button.ButtonIdentity
			buttonCreates = append(buttonCreates, button)
		}

		// 将按钮批量写入数据库
		err = tx.Create(&buttonCreates).Error
		if err != nil {
			tx.Rollback()
			return errors.New("创建按钮失败")
		}
	}
	// 提交数据库事务
	return tx.Commit().Error
}

// DeleteMenu
/**
 *  @Description: 删除菜单项及子菜单项（被角色绑定的菜单不可删除）
 *  @receiver menuService
 *  @param id 需要删除的菜单项ID
 *  @return err
 */
func (menuService *MenuService) DeleteMenu(id string) (err error) {
	// 查询该该菜单的子菜单
	var menus []menumod.Menu
	global.DB.Where("parent_id = ?", id).Find(&menus)
	// 获取所有子菜单ID
	var menuIds []string
	for _, menu := range menus {
		menuIds = append(menuIds, menu.ID)
	}
	menuIds = append(menuIds, id)

	// 菜单不能被角色绑定
	var roleMenuCount int64
	err = global.DB.Model(&rolemod.RoleMenu{}).Where("menu_id in ?", menuIds).Count(&roleMenuCount).Error
	if err != nil || roleMenuCount > 0 {
		global.DB.Rollback()
		return errors.New("被角色绑定的菜单不可删除")
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 删除菜单
	err = tx.Where("id in ?", menuIds).Delete(&menumod.Menu{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除菜单失败")
	}

	// 提交数据库事务
	return tx.Commit().Error
}

// UpdateMenu
/**
 *  @Description: 更新菜单项
 *  @receiver menuService
 *  @param menu
 *  @return err
 */
func (menuService *MenuService) UpdateMenu(menu menumod.Menu) (err error) {
	// 菜单名称重复效验
	var m menumod.Menu
	// 若存在父菜单，则需对父菜单进行效验
	if strings.TrimSpace(menu.ParentId) != "" {
		// 判定父菜单是否合法
		_, err = menuService.GetMenuInfo(menu.ParentId)
		if err != nil {
			return errors.New("非法的父菜单")
		}

		// 判定菜单名称是否被其他子菜单占用
		global.DB.Where("menu_name = ?", menu.MenuName).Not("parent_id = ?", "").First(&m)
		if m.ID != "" && m.ID != menu.ID {
			return errors.New("该菜单名称已被其他子菜单占用")
		}
	} else {
		// 判定菜单名称是否被其他父菜单占用
		global.DB.Where(map[string]interface{}{"menu_name": menu.MenuName, "parent_id": nil}).Or("menu_name = ? AND parent_id = ?", menu.MenuName, "").First(&m)
		if m.ID != "" && m.ID != menu.ID {
			return errors.New("该菜单名称已被其他父菜单占用")
		}
	}

	// URL重名效验
	var m1 menumod.Menu
	global.DB.Where("url = ?", menu.Url).First(&m1)
	if m1.ID != "" && m1.ID != menu.ID {
		return errors.New("菜单路径（url）不可重复")
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 更新菜单
	err = tx.Updates(&menu).Error
	if err != nil {
		tx.Rollback()
		return errors.New("更新菜单失败")
	}

	// 效验不可删除的按钮
	var roleButtons []rolemod.RoleButton
	var canNotDeleteIds []string
	err = tx.Where("menu_id", menu.ID).Find(&roleButtons).Error
	for _, button := range roleButtons {
		canNotDeleteIds = append(canNotDeleteIds, button.ButtonId)
	}
	if len(canNotDeleteIds) > 0 {
		for _, button := range menu.Buttons {
			if strings.TrimSpace(button.ID) == "" {
				continue
			}
			ok := array.IsStrArrayExistArray(canNotDeleteIds, button.ID)
			if !ok {
				tx.Rollback()
				return errors.New("（" + button.ButtonName + "）按钮不可删除")
			}
		}
	}

	// 绑定按钮
	// 判定按钮是否存在重复值
	var buttonNames []string
	var ButtonIdentities []string
	for _, button := range menu.Buttons {
		if strings.TrimSpace(button.ButtonName) == "" {
			return errors.New("存在非法的按钮名称")
		} else {
			buttonNames = append(buttonNames, button.ButtonName)
		}
		if strings.TrimSpace(button.ButtonIdentity) == "" {
			return errors.New("存在非法的按钮标识")
		} else {
			ButtonIdentities = append(ButtonIdentities, button.ButtonIdentity)
		}
	}
	if array.IsExistRepeatInArray(buttonNames) {
		return errors.New("按钮名称不唯一")
	}
	if array.IsExistRepeatInArray(ButtonIdentities) {
		return errors.New("按钮标识不唯一")
	}

	var CreateButtons []menumod.Button
	url, err := menuService.GetMenuPath(tx, menu.ID)
	if err != nil {
		return errors.New("解析URL失败")
	}
	for _, button := range menu.Buttons {
		button.MenuId = menu.ID
		button.PageName = menu.MenuName
		button.PagePath = url
		button.ButtonFullIdentity = strings.Replace(url, "/", ":", -1) + ":" + button.ButtonIdentity
		CreateButtons = append(CreateButtons, button)
	}

	// 删除原有的按钮
	err = tx.Where("menu_id = ?", menu.ID).Delete(&menumod.Button{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除原有按钮失败")
	}

	if len(CreateButtons) > 0 {
		err = tx.Create(&CreateButtons).Error
		if err != nil {
			tx.Rollback()
			return errors.New("创建按钮失败")
		}
	}

	// 提交数据库事务
	return tx.Commit().Error
}

// GetMenuInfo
/**
 *  @Description: 获取单个菜单项详情
 *  @receiver menuService
 *  @param id
 *  @return menu
 *  @return err
 */
func (menuService *MenuService) GetMenuInfo(id string) (menu menumod.Menu, err error) {
	err = global.DB.Where("id = ?", id).First(&menu).Error
	if err != nil {
		return menu, errors.New("该菜单项不存在")
	}

	// 获取页面按钮
	var buttons []menumod.Button
	err = global.DB.Where("menu_id = ?", menu.ID).Find(&buttons).Error
	if err != nil {
		return menu, errors.New("获取页面按钮失败")
	}
	menu.Buttons = buttons
	menu.Children, err = menuService.GetMenuChildren(id)
	if err != nil {
		return menu, errors.New("查询子节点出错")
	}

	// 获取禁止操作的按钮
	var buttonIds []string
	for _, button := range buttons {
		buttonIds = append(buttonIds, button.ID)
	}

	var roleButtons []rolemod.RoleButton
	err = global.DB.Distinct("button_id").Where("button_id in ?", buttonIds).Find(&roleButtons).Error
	if err != nil {
		return menumod.Menu{}, errors.New("获取不能操作的按钮失败")
	}

	var disableButtons []menumod.Button
	if roleButtons != nil {
		buttonDisableMap := make(map[string]int)
		for _, button := range roleButtons {
			buttonDisableMap[button.ButtonId] = 1
		}
		for _, button := range buttons {
			_, ok := buttonDisableMap[button.ID]
			if ok {
				button.Disable = true
			}
			disableButtons = append(disableButtons, button)
		}
	}
	menu.Buttons = disableButtons
	return
}

// GetMenuChildren
/**
 *  @Description: 获取当前节点的子节点
 *  @receiver menuService
 *  @param id
 *  @return menuChild
 *  @return err
 */
func (menuService *MenuService) GetMenuChildren(id string) (menuChild []menumod.Menu, err error) {
	err = global.DB.Where("parent_id = ?", id).Order("sort").Find(&menuChild).Error
	return
}

// GetAllMenus
/**
 *  @Description: 查询所有菜单项（不分页）
 *  @receiver menuService
 *  @return menus
 *  @return err
 */
func (menuService *MenuService) GetAllMenus() (menus []menumod.Menu, err error) {
	err = global.DB.Find(&menus).Error
	return
}

// GetMenuPage
/**
 *  @Description: 菜单项分页查询
 *  @receiver menuService
 *  @param pageInfo
 *  @return err
 *  @return pageBean
 */
func (menuService *MenuService) GetMenuPage(pageInfo *page.PageInfo) (err error, pageBean *page.PageBean) {
	pageBean = &page.PageBean{Page: pageInfo.Current, PageSize: pageInfo.RowCount}
	rows := make([]*menumod.Menu, 0)
	err, pageBean = db.FindPage(&menumod.Menu{}, &rows, pageInfo)
	return
}

// GetMenuTree
/**
 *  @Description: 获取菜单树
 *  @receiver menuService
 *  @param menuIds 需要查询的菜单ID切片
 *  @param getAll 是否获取全部数据（切片为空时使用）
 *  @return menus
 *  @return err
 */
func (menuService *MenuService) GetMenuTree(menuIds []string, getAll bool) (menus []menumod.Menu, err error) {
	// menuIds为空则不查询
	if len(menuIds) == 0 && getAll != true {
		return nil, err
	}
	// 有则查询子菜单
	treeMap, err := menuService.GetMenuTreeMap(menuIds)
	if err != nil {
		return nil, errors.New("查询菜单树失败")
	}
	var menuBase []menumod.Menu
	// 查询一级目录
	menuBase = treeMap[""]
	for _, menu := range menuBase {
		menu.Children = treeMap[menu.ID]
		var menus1 []menumod.Menu
		for _, menu1 := range menu.Children {
			menu1.Children = treeMap[menu1.ID]
			var menus2 []menumod.Menu
			for _, menu2 := range menu1.Children {
				menu2.Children = treeMap[menu2.ID]
				menus2 = append(menus2, menu2)
			}
			menu.Children = menus1
			menus1 = append(menus1, menu1)
		}
		menu.Children = menus1
		menus = append(menus, menu)
	}
	return menus, err
}

// GetMenuTreeMap
/**
 *  @Description: 查询菜单树Map（最后一级为按钮）
 *  @receiver menuService
 *  @param menuIds
 *  @return treeMap
 *  @return err
 */
func (menuService *MenuService) GetMenuTreeMap(menuIds []string) (treeMap map[string][]menumod.Menu, err error) {
	var allMenus []menumod.Menu
	treeMap = make(map[string][]menumod.Menu)
	if len(menuIds) > 0 {
		err = global.DB.Where("id in ?", menuIds).Order("sort").Find(&allMenus).Error
	} else {
		err = global.DB.Order("sort").Find(&allMenus).Error
	}

	var buttons []menumod.Button
	if len(menuIds) > 0 {
		err = global.DB.Where("menu_id in ?", menuIds).Find(&buttons).Error
	} else {
		err = global.DB.Find(&buttons).Error
	}
	if len(buttons) > 0 {
		for _, button := range buttons {
			var buttonToMenu menumod.Menu
			// 用menu结构体给前端返回button相关的，所以用'button'区分开
			buttonToMenu.ID = "button-" + uuid.UUID()
			buttonToMenu.ParentId = button.MenuId
			buttonToMenu.ButtonId = button.ID
			buttonToMenu.ButtonIdentity = button.ButtonIdentity
			buttonToMenu.ButtonName = button.ButtonName
			buttonToMenu.IsShow = menumod.HIDE
			allMenus = append(allMenus, buttonToMenu)
		}
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return
}

// GetMenusByIds
/**
 *  @Description: 根据MenuIds获取menu对象
 *  @receiver menuService
 *  @param menuIds
 *  @return menus
 *  @return err
 */
func (menuService *MenuService) GetMenusByIds(menuIds []string) (menus []menumod.Menu, err error) {
	err = global.DB.Where("id in ?", menuIds).Find(&menus).Error
	return
}

// GetMenuPath
/**
 *  @Description: 获取当前的url
 *  @receiver menuService
 *  @param tx
 *  @param menuId
 *  @return path
 *  @return err
 */
func (menuService *MenuService) GetMenuPath(tx *gorm.DB, menuId string) (path string, err error) {
	var menu menumod.Menu
	err = tx.Where("id = ?", menuId).First(&menu).Error
	if err != nil {
		return "", errors.New("查询菜单失败")
	}
	path = menu.Url
	if strings.TrimSpace(menu.ParentId) != "" {
		menuPath, err := menuService.GetMenuPath(tx, menu.ParentId)
		if err != nil {
			return "", errors.New("查询菜单失败")
		}
		path = menuPath + "/" + menu.Url
	} else {
		return path, nil
	}
	return path, nil
}
