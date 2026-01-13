package main

import (
	"fmt"

	getdata "github.com/mucusscraper/pdb_metadata_tool/internal/get_data"
)

const issueURL = "https://data.rcsb.org/rest/v1/core/entry/7urx"

func main() {
	fmt.Printf("Hello World test!\n")
	res, err := getdata.GetIssueData(issueURL)
	if err != nil {
		fmt.Printf("Error getting data\n")
		return
	}
	fmt.Printf("%v\n", res.ID)
	fmt.Printf("%v\n", res.AccessInfo.DepositDate)
	fmt.Printf("%v\n", res.ArticleInfo.DOI)
	fmt.Printf("%v\n", res.ArticleInfo.Title)
}
