package entity

type Counterparty struct {
	Code_ou                               string `json:"code_ou"`
	Inn                                   string `json:"inn"`
	Institution_short_name                string `json:"institution_short_name"`
	Institution_full_name                 string `json:"institution_full_name"`
	Address                               string `json:"address"`
	City                                  string `json:"city"`
	Bank_details                          string `json:"bank_details"`
	Responsible_person_job_title          string `json:"responsible_person_job_title"`
	Responsible_person_short_name         string `json:"responsible_person_short_name"`
	Responsible_person_full_name          string `json:"responsible_person_full_name"`
	Responsible_person_full_name_genitive string `json:"responsible_person_full_name_genitive"`
	Acting_on                             string `json:"acting_on"`
	Ikz_2025                              string `json:"ikz_2025"`
	Source_funding                        string `json:"source_funding"`
	Email                                 string `json:"email"`
	Phone_number                          string `json:"phone_number"`
	Contract_form                         string `json:"contract_form"`
	Contract_type                         string `json:"contract_type"`
	Contract_number                       string `json:"contract_number"`
	Contract_formation_data               string `json:"contract_formation_data"`
	Responsible_person_job_title_genetive string `json:"responsible_person_job_title_genetive"`
	Category                              string `json:"category"`
}

type Counterparties []*Counterparty

type KeyCounterparties struct{}
