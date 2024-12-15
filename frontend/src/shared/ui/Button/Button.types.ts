export default interface IButton {
	data: {
		text: string;
	};
	cn?: {
		size?: 'small' | 'medium' | 'large';
		width?: string;
		margin?: string;
	};
	handler?: {
		onClick: () => void;
	};
}
