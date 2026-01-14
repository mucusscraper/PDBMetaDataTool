package main

import (
	"fmt"

	getdata "github.com/mucusscraper/pdb_metadata_tool/internal/get_data"
)

const issueURL = "https://data.rcsb.org/rest/v1/core/entry/7o52"

func main() {
	fmt.Printf("Hello World test!\n")
	res, err, list_of_urls_polymer, list_of_urls_non_polymers := getdata.GetIssueDataEntry(issueURL)
	if err != nil {
		fmt.Printf("Error getting data\n")
		return
	}
	fmt.Printf("%v\n", res.ID)
	fmt.Printf("%v\n", res.AccessInfo.DepositDate)
	fmt.Printf("%v\n", res.ArticleInfo.DOI)
	fmt.Printf("%v\n", res.ArticleInfo.Title)
	for _, method := range res.ExptlInfo {
		fmt.Printf("%v\n", method)
	}
	for _, polymer_id := range res.EntitiesInfo.PolymerID {
		fmt.Printf("%v\n", polymer_id)
	}
	for _, non_polymer_id := range res.EntitiesInfo.NonPolymerID {
		fmt.Printf("%v\n", non_polymer_id)
	}
	for _, polymer_url := range list_of_urls_polymer {
		fmt.Printf("%v\n", polymer_url)
	}
	for _, non_polymer_url := range list_of_urls_non_polymers {
		fmt.Printf("%v\n", non_polymer_url)
		NonPolymer, err := getdata.GetDataForNonPolymers(non_polymer_url)
		if err != nil {
			fmt.Printf("Error getting data")
			return
		}
		fmt.Printf("NAME: %v\n", NonPolymer.Entity.Name)
		fmt.Printf("COMP ID: %v\n", NonPolymer.Entity.CompID)
		fmt.Printf("DESCRIPTION: %v\n", NonPolymer.Data.Description)
		fmt.Printf("FORMULA WEIGHT: %v\n", NonPolymer.Data.FormulaWeight)
		fmt.Printf("NUMBER OF MOLECULES: %v\n", NonPolymer.Data.NumberOfMolecules)
	}
}
