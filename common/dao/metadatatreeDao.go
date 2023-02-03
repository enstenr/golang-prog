package dao

import (
	"bytes"

	"fmt"
	"github.com/enstenr/common/connection"

	"github.com/enstenr/customtypes"
	
	_ "github.com/lib/pq"
)


func FetchTree(env string) []customtypes.MetadataTree {
	itemArray := make([]customtypes.MetadataTree, 0)
	db := connection.InitConnection(env)
	defer db.Close()

	buf := bytes.NewBufferString(`select mt."metadataTreeConfigurationId",mtc.name,mt.hash,mt.modified_date,mt."rePublishStatus" 
	from metadata_tree mt, metadata_tree_configuration mtc
	where mt."metadataTreeConfigurationId"=mtc."metadataTreeConfigurationId"`)

	rows, err := db.Query(buf.String())

	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {

		metadataTree := new(customtypes.MetadataTree)

		_ = rows.Scan(&metadataTree.MetadataTreeConfigurationId,&metadataTree.Name, &metadataTree.Hash, &metadataTree.Modified_date, &metadataTree.RePublishStatus)

		itemArray = append(itemArray, *metadataTree)
	}
	return itemArray
}


