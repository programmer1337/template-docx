package entity

// -------------------------------------------//
// Counterparty //
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

var CounterpartyRuEngMap = map[string]string{
	"Код ОУ": "Code_ou",
	"ИНН":    "Inn",
	"Сокращенное наименование учреждения":                      "Institution_short_name",
	"Полное наименование":                                      "Institution_full_name",
	"Юридический и почтовый адрес образовательного учреждения": "Address",
	"Район": "City",
	"Банковские реквизиты":                  "Bank_details",
	"Должность подписанта (им.падеж)":       "Responsible_person_job_title",
	"ФИО подписанта (сокр.)":                "Responsible_person_short_name",
	"ФИО подписанта полностью (им. Падеж)":  "Responsible_person_full_name",
	"ФИО подписанта полностью (род. падеж)": "Responsible_person_full_name_genitive",
	"действующего на основании":             "Acting_on",
	"ИКЗ 2025": "Ikz_2025",
	"Источник финансирования": "Source_funding",
	"e-mail":         "Email",
	"Телефон":        "Phone_number",
	"Форма договора": "Contract_form",
	"Тип договора":   "Contract_type",
	"Номер договора": "Contract_number",
	"Дата формирования договора":        "Contract_formation_data",
	"Должность подписанта (род. Падеж)": "Responsible_person_job_title_genetive",
	"Категория": "Category",
}

var CounterpartyEngRuAlias = map[string]string{
	"Code_ou":                               "Код ОУ",
	"Inn":                                   "ИНН",
	"Institution_short_name":                "Сокращенное наименование учреждения",
	"Institution_full_name":                 "Полное наименование",
	"Address":                               "Юридический и почтовый адрес образовательного учреждения",
	"City":                                  "Район",
	"Bank_details":                          "Банковские реквизиты",
	"Responsible_person_job_title":          "Должность подписанта (им.падеж)",
	"Responsible_person_short_name":         "ФИО подписанта (сокр.)",
	"Responsible_person_full_name":          "ФИО подписанта полностью (им. Падеж)",
	"Responsible_person_full_name_genitive": "ФИО подписанта полностью (род. падеж)",
	"Acting_on":                             "действующего на основании",
	"Ikz_2025":                              "ИКЗ 2025",
	"Source_funding":                        "Источник финансирования",
	"Email":                                 "e-mail",
	"Phone_number":                          "Телефон",
	"Contract_form":                         "Форма договора",
	"Contract_type":                         "Тип договора",
	"Contract_number":                       "Номер договора",
	"Contract_formation_data":               "Дата формирования договора",
	"Responsible_person_job_title_genetive": "Должность подписанта (род. Падеж)",
	"Category":                              "Категория",
}

// -------------------------------------------//
// Entity //

// -------------------------------------------//
// Entity //

// -------------------------------------------//
// Entity //
