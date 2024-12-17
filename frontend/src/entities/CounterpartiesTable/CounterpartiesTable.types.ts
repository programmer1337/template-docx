export default interface CounterpartiesTable {
	data: {
		counterparties: Counterparties | null;
	};
	cn?: {
		padding?: string;
	};
}

export interface ICounterparty {
	data: {
		counterparty: Counterparty;
	};
	cn?: {
		margin?: string;
	};
}
