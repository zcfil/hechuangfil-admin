package export

import (
	"hechuangfil-admin/config"
)

/**
 *@Project     hechuangfil-admin
 *@Author      king
 *@CreateTime  2020/4/12 2:17 下午
 *@ClassName   excel
 *@Description excle 导出，导入设置
 */

const EXT = ".xlsx"

// GetExcelFullUrl get the full access path of the Excel file
func GetExcelFullUrl(name string) string {
	return config.ApplicationConfig.Host + config.ApplicationConfig.Port + GetExcelPath() + name
	//setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return config.ApplicationConfig.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return config.ApplicationConfig.RuntimeRootPath + GetExcelPath()
}
