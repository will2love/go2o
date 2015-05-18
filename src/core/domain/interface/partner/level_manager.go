/**
 * Copyright 2015 @ S1N1 Team.
 * name : member_manager.go
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package partner
import "go2o/src/core/domain/interface/valueobject"

type ILevelManager interface {
    // 获取等级设置
    GetLevelSet()[]*valueobject.MemberLevel

    // 获取等级
    GetLevelById(id int)*valueobject.MemberLevel

    // 删除等级
    DeleteLevel(id int)error

    // 保存等级
    SaveLevel(*valueobject.MemberLevel)(int,error)
}