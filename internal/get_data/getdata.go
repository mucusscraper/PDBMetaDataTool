package getdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProteinIssue struct {
	ID           string             `json:"rcsb_id"`
	AccessInfo   RcsbAccessionInfo  `json:"rcsb_accession_info"`
	ArticleInfo  ArticleAccessInfo  `json:"rcsb_primary_citation"`
	ExptlInfo    []ExptlAccessInfo  `json:"exptl"`
	EntitiesInfo EntitiesAccessInfo `json:"rcsb_entry_container_identifiers"`
}

type RcsbAccessionInfo struct {
	DepositDate string `json:"deposit_date"`
}

type ArticleAccessInfo struct {
	DOI   string `json:"pdbx_database_id_doi"`
	Title string `json:"title"`
}

type ExptlAccessInfo struct {
	Method string `json:"method"`
}

type EntitiesAccessInfo struct {
	PolymerID    []string `json:"polymer_entity_ids"`
	NonPolymerID []string `json:"non_polymer_entity_ids"`
}

type NonPolymerIssue struct {
	Entity NameEntityNonPolymerAccession `json:"pdbx_entity_nonpoly"`
	Data   DataEntityNonPolymerAccession `json:"rcsb_nonpolymer_entity"`
}

type NameEntityNonPolymerAccession struct {
	Name   string `json:"name"`
	CompID string `json:"comp_id"`
}

type DataEntityNonPolymerAccession struct {
	FormulaWeight     float32 `json:"formula_weight"`
	Description       string  `json:"pdbx_description"`
	NumberOfMolecules int     `json:"pdbx_number_of_molecules"`
}

func GetIssueDataEntry(url string) (ProteinIssue, error, []string, []string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error here 1\n")
		return ProteinIssue{}, err, nil, nil
	}
	defer res.Body.Close()
	var PDB ProteinIssue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&PDB)
	if err != nil {
		fmt.Printf("Error here 2\n")
		return ProteinIssue{}, err, nil, nil
	}
	var list_of_urls_polymers []string
	for _, polymer_id := range PDB.EntitiesInfo.PolymerID {
		list_of_urls_polymers = append(list_of_urls_polymers, fmt.Sprintf("https://data.rcsb.org/rest/v1/core/polymer_entity/7o52/%v", polymer_id))
	}
	var list_of_urls_non_polymers []string
	for _, non_polymer_id := range PDB.EntitiesInfo.NonPolymerID {
		list_of_urls_non_polymers = append(list_of_urls_non_polymers, fmt.Sprintf("https://data.rcsb.org/rest/v1/core/nonpolymer_entity/7o52/%v", non_polymer_id))
	}
	return PDB, nil, list_of_urls_polymers, list_of_urls_non_polymers
}

func GetDataForNonPolymers(url string) (NonPolymerIssue, error) {
	res, err := http.Get(url)
	if err != nil {
		return NonPolymerIssue{}, err
	}
	defer res.Body.Close()
	var NonPolymer NonPolymerIssue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&NonPolymer)
	if err != nil {
		return NonPolymerIssue{}, err
	}
	return NonPolymer, nil
}

func GetEntitiesURL(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error here 3\n")
		return nil, err
	}
	defer res.Body.Close()
	var PDB ProteinIssue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&PDB)
	if err != nil {
		fmt.Printf("Error here 3\n")
		return nil, err
	}
	var list_of_urls []string
	for _, entity := range PDB.EntitiesInfo.PolymerID {
		list_of_urls = append(list_of_urls, fmt.Sprintf("curl https://data.rcsb.org/rest/v1/core/polymer_entity/7urx/%v", entity))
	}
	return list_of_urls, nil
}
