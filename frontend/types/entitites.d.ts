// type Counterparty = {
// 	A: string;
// 	B: string;
// 	C: string;
// 	D: string;
// 	E: string;
// 	F: string;
// 	G: string;
// 	H: string;
// 	I: string;
// 	J: string;
// 	K: string;
// 	L: string;
// 	M: string;
// 	N: string;
// 	O: string;
// 	P: string;
// 	Q: string;
// 	R: string;
// 	S: string;
// 	T: string;
// 	U: string;
// 	V: string;
// };

type CounterpartyKeys = keyof Counterparty;

type Counterparty = {
	code_ou: string;
	inn: string;
	institution_short_name: string;
	institution_full_name: string;
	address: string;
	city: string;
	bank_details: string;
	responsible_person_job_title: string;
	responsible_person_short_name: string;
	responsible_person_full_name: string;
	responsible_person_full_name_genitive: string;
	acting_on: string;
	ikz_2025: string;
	source_funding: string;
	email: string;
	phone_number: string;
	contract_form: string;
	contract_type: string;
	contract_number: string;
	contract_formation_data: string;
	responsible_person_job_title_genetive: string;
	category: string;
};

type Counterparties = Counterparty[];
