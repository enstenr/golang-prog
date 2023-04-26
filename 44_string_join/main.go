package main

import (
	"fmt"
	"strings"
)

func main() {
	tags := []string{"TreePublishChangeLogId", "tree_publish_job_id", "status tree_id"}

	jsonStr := (fmt.Sprintf(`"%s"`,
		strings.Join(tags, `", "`)))
	fmt.Print(tags)
	fmt.Print(jsonStr)

}