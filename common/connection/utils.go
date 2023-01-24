package connection

import (
	"github.com/enstenr/customtypes"
)
func formTreeMap(duplicateSkuReportObjArray []customtypes.DupliateSkuReport) map[string]customtypes.DupliateSkuReport {
	treeMap := make(map[string]customtypes.DupliateSkuReport, 0)
	for _, duplicateSkuReportObj := range duplicateSkuReportObjArray {
		treeMap[duplicateSkuReportObj.Tree_name] = duplicateSkuReportObj

	}
	return treeMap
}