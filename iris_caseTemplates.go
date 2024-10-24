package iris

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AddCaseTemplateResponse represents the response of Add a case template resonse
type CaseTemplateAPIResponse struct {
	CaseTemplate CaseTemplateResponse `json:"data"`
	ApiMeta
}

// CaseTemplateResponse contains the response of a case template api action
type CaseTemplateResponse struct {
	Registry             any    `json:"registry"`
	TypeDescription      string `json:"type_description"`
	TypeID               int    `json:"type_id"`
	TypeName             string `json:"type_name"`
	TypeTaxonomy         string `json:"type_taxonomy"`
	TypeValidationExpect string `json:"type_validation_expect"`
	TypeValidationRegex  string `json:"type_validation_regex"`
}


// AddCaseTemplate adds a case template using the /manag/case-templates/add endpoint
// It returns a AddCaseTemplateResponse 
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Returns:
// - *AddCaseTemplateResponse*: The response from the API containing the template informations in the CaseTemplate field.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) AddCaseTemplate(caseTemplate string) (*CaseTemplateAPIResponse, error) {
	jsondata, err := json.Marshal(map[string]string{"case_template_json": caseTemplate})
	if err != nil {
		return nil, err
	}

	builder := NewRequestBuilder().
		SetURL("/manage/case-templates/add").
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()


	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var templateResponse CaseTemplateAPIResponse
	if err := json.NewDecoder(req.Body).Decode(&templateResponse); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &templateResponse, nil
}


// UpdateCaseTemplate updates a case template using the /manag/case-templates/update/ endpoint
// It returns a AddCaseTemplateResponse 
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Returns:
// - *CaseTemplateAPIResponse*: The response from the API containing the template informations in the CaseTemplate field.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) UpdateCaseTemplate(templateId int, caseTemplate string) (*CaseTemplateAPIResponse, error) {
	jsondata, err := json.Marshal(map[string]string{"case_template_json": caseTemplate})
	if err != nil {
		return nil, err
	}

	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/case-templates/update/%d",templateId)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()


	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var templateResponse CaseTemplateAPIResponse
	if err := json.NewDecoder(req.Body).Decode(&templateResponse); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &templateResponse, nil
}

// DeleteCaseTemplate removes a case template using the /manage/case-templates/delete endpoint
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Returns:
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) DeleteCaseTemplate(templateId int) error {
	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/case-templates/delete/%d",templateId)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return err
	}
	defer req.Body.Close()


	return err
}

