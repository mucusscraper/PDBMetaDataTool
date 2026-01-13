package getdata

import (
	"encoding/json"
	"net/http"
)

type ProteinIssue struct {
	ID          string            `json:"rcsb_id"`
	AccessInfo  RcsbAccessionInfo `json:"rcsb_accession_info"`
	ArticleInfo ArticleAccessInfo `json:"rcsb_primary_citation"`
	ExptlInfo   ExptlAccessInfo   `json:"exptl"`
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

func GetIssueData(url string) (ProteinIssue, error) {
	res, err := http.Get(url)
	if err != nil {
		return ProteinIssue{}, err
	}
	defer res.Body.Close()
	var protein ProteinIssue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&protein)
	if err != nil {
		return ProteinIssue{}, err
	}
	return protein, nil
}
