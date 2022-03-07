/**
 * @Time: 2022/3/7 19:05
 * @Author: yt.yin
 */

package datares

import  "github.com/goworkeryyt/aramis/server/data/model"

// ApiInfoTreeRes 给前端返回的api列表树结构
type ApiInfoTreeRes struct {

	/** API组名描述 */
	Description           string                   `json:"description"`

	/** API集合 */
	ApiInfos              []datamod.ApiInfo        `json:"apiInfos"`
}


