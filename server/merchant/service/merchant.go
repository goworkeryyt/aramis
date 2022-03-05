/**
 * @Time: 2022/3/5 15:39
 * @Author: yt.yin
 */

package mchtsrv

import (
	"errors"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/aramis/server/merchant/model/request"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MerchantService struct{}

// CreateMerchant
/**
 *  @Description: 创建商户
 *  @receiver merchantService
 *  @param merchant
 *  @param parentMerchantNo 父商户编号
 *  @return err
 */
func (merchantService *MerchantService) CreateMerchant(merchant mchtmod.Merchant, parentMerchantNo string) (err error) {
	var m mchtmod.Merchant
	var parentMerchant mchtmod.Merchant
	// 判断是否存在相同的商户编号
	if !errors.Is(global.DB.Where("merchant_no = ?", merchant.MerchantNo).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("商户编号被占用")
	}
	// 判断是否存在相同的商户名称
	if !errors.Is(global.DB.Where("merchant_name = ?", merchant.MerchantName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("商户名称被占用")
	}
	// 若存在父商户
	if strings.TrimSpace(parentMerchantNo) != "" && merchant.Level == 2 {
		err = global.DB.Where("merchant_no = ? AND level = ?", parentMerchantNo, 1).First(&parentMerchant).Error
		if err != nil {
			return errors.New("父商户不存在")
		}
		merchant.ParentId = parentMerchant.ID
	}
	// 商户状态赋值为 初始状态启用
	merchant.Status = mchtmod.ENABLE
	merchant.ID = uuid.UUID()
	merchant.CreateTime = carbon.Now().ToDateTimeString()
	// 创建商户
	err = global.DB.Create(&merchant).Error
	if err != nil {
		return errors.New("创建父商户失败")
	}
	return nil
}

// UpdateMerchant
/**
 *  @Description: 更新商户
 *  @receiver merchantService
 *  @param merchant
 *  @return err
 */
func (merchantService *MerchantService) UpdateMerchant(merchantInput mchtreq.UpdateMerchantReq) (err error) {
	// 判断该商户是否存在
	var merchantUpdate mchtmod.Merchant
	err = global.DB.Where("id = ?", merchantInput.ID).First(&merchantUpdate).Error
	if err != nil {
		return errors.New("该商户不存在")
	}
	var m mchtmod.Merchant
	// 判断是否存在相同的商户名称
	global.DB.Where("merchant_name = ?", merchantInput.MerchantName).First(&m)
	if m.ID != "" && m.ID != merchantInput.ID {
		return errors.New("商户名称被占用")
	}
	// 判断商户状态是否合法
	if merchantInput.Status != mchtmod.ENABLE && merchantInput.Status != mchtmod.DISABLE {
		return errors.New("非法的商户状态")
	}
	err = copier.Copy(merchantUpdate, merchantInput)
	if err != nil{
		return err
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 更新商户
	err = tx.Updates(&merchantUpdate).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建父商户失败")
	}
	// 商户停用需联动停用所有绑定该商户的用户,以及该商户的子商户
	if merchantInput.Status == mchtmod.DISABLE {
		err = tx.Model(&usermod.User{}).Where("merchant_no = ?", merchantUpdate.MerchantNo).Update("status", usermod.DISABLE).Error
		if err != nil {
			tx.Rollback()
			return errors.New("联动更新该商户对应的用户失败")
		}
		// 联动更新子商户
		err = tx.Model(&mchtmod.Merchant{}).Where("parent_id = ?", merchantInput.ID).Update("status", mchtmod.DISABLE).Error
		if err != nil {
			tx.Rollback()
			return errors.New("联动更新子商户状态失败")
		}
	}
	// 提交数据库事务
	return tx.Commit().Error
}

// DeleteMerchant
/**
 *  @Description: 删除商户（有角色，有登录过的用户不能删除）
 *  @receiver merchantService
 *  @param id
 *  @return err
 */
func (merchantService *MerchantService) DeleteMerchant(id string) (err error) {
	// 已经不存在的商户无需删除
	var merchant mchtmod.Merchant
	err = global.DB.Where("id = ?", id).First(&merchant).Error
	if err != nil {
		return nil
	}
	// 存在子商户的商户无法删除
	var subMerchantCount int64
	err = global.DB.Model(&mchtmod.Merchant{}).Where("parent_id = ?", merchant.ID).Count(&subMerchantCount).Error
	if err != nil || subMerchantCount > 0 {
		return errors.New("存在子商户的商户无法删除")
	}

	// 被角色绑定的商户不可删除
	var roleCount int64
	err = global.DB.Model(&rolemod.Role{}).Where("merchant_no = ?", merchant.MerchantNo).Count(&roleCount).Error
	if err != nil || roleCount > 0 {
		return errors.New("被角色绑定的商户不可删除")
	}

	// 该商户存在登录过的用户则不可删除
	var userCount int64
	err = global.DB.Model(&usermod.User{}).Where("merchant_no = ?", merchant.MerchantNo).Where("user_type = ?", usermod.MERT).Where("login_count > ?", 0).Count(&userCount).Error
	if err != nil || userCount > 0 {
		return errors.New("该商户已被用户使用，无法删除")
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 删除商户账号
	err = tx.Where("merchant_no = ?", merchant.MerchantNo).Where("user_type = ?", usermod.MERT).Delete(&usermod.User{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除商户账号失败")
	}
	// 删除商户
	err = tx.Where("id = ?", id).Delete(&merchant).Error
	if err != nil {
		tx.Rollback()
		return errors.New("删除商户失败")
	}
	// 提交数据库事务
	return tx.Commit().Error
}

// GetMerchantPage
/**
 *  @Description: 商户分页查询
 *  @receiver merchantService
 *  @param pageInfo
 *  @return err
 *  @return pageBean
 */
func (merchantService *MerchantService) GetMerchantPage(pageInfo *page.PageInfo) (err error, pageBean *page.PageBean) {
	pageBean = &page.PageBean{Page: pageInfo.Current, PageSize: pageInfo.RowCount}
	rows := make([]*mchtmod.Merchant, 0)
	err, pageBean = db.FindPage(&mchtmod.Merchant{}, &rows, pageInfo)
	// 关联父商户信息,得出父商户名称
	for i, _ := range rows {
		merchant := *rows[i]
		var parentMerchant mchtmod.Merchant
		// global.DB.Select("merchantName").Where("id = ?", merchant.ParentId).First(&parentMerchant)
		if merchant.Level == 2 {
			err = global.DB.Where("id = ?", merchant.ParentId).First(&parentMerchant).Error
			if err != nil {
				return errors.New("查询父商户异常," + merchant.ParentName), pageBean
			}
			merchant.ParentName = parentMerchant.MerchantName
		} else {
			merchant.ParentName = ""
		}
		*rows[i] = merchant
	}

	return
}

// FindMerchantInfo
/**
 *  @Description: 获取商户详情
 *  @receiver merchantService
 *  @param id
 *  @return merchant
 *  @return err
 */
func (merchantService *MerchantService) FindMerchantInfo(id string) (merchant mchtmod.Merchant, err error) {
	err = global.DB.Where("id = ?", id).First(&merchant).Error
	if err != nil {
		return mchtmod.Merchant{}, errors.New("该商户不存在")
	}
	children, err := merchantService.FindMerchantChildren(id)
	if err != nil {
		return mchtmod.Merchant{}, errors.New("查询子商户异常")
	}
	merchant.ChildMerchants = children
	return
}

// FindMerchantChildren
/**
 *  @Description: 查询所有子商户
 *  @receiver merchantService
 *  @param id
 *  @return merchants
 *  @return err
 */
func (merchantService *MerchantService) FindMerchantChildren(id string) (merchants []mchtmod.Merchant, err error) {
	err = global.DB.Where("parent_id = ?", id).Find(&merchants).Error
	return
}

// GetAllMerchants
/**
 *  @Description: 查询所有商户
 *  @receiver merchantService
 *  @param parentId
 *  @return merchants
 *  @return err
 */
func (merchantService *MerchantService) GetAllMerchants(parentId string) (merchants []mchtmod.Merchant, err error) {
	if strings.TrimSpace(parentId) != "" {
		err = global.DB.Where("parent_id = ?", parentId).Find(&merchants).Error
	} else {
		err = global.DB.Find(&merchants).Error
	}
	return
}

// FindMerchantByMerchantNo
/**
 *  @Description: 通过商户编号查询商户
 *  @receiver merchantService
 *  @param merchantNo
 *  @return merchant
 *  @return err
 */
func (merchantService *MerchantService) FindMerchantByMerchantNo(merchantNo string) (merchant mchtmod.Merchant, err error) {
	err = global.DB.Where("merchant_no = ?", merchantNo).First(&merchant).Error
	return
}

