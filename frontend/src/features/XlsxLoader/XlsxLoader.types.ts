interface IXlsxLoader {
	cn?: {
		padding?: string;
	};
	executor?: {
		setCounterparties: (counterParties: any) => void;
	};
}

export default IXlsxLoader;
