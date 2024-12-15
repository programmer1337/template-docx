export default interface IInputFile {
	data: {
		file: File | null;
	};
	executor: {
		setFile: (file: File) => void;
	};
}
