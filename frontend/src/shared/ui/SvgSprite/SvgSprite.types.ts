import React from 'react';

interface ISvgSprite {
	data: {
		name: string;
		width?: number;
		height?: number;
	};
	className?: string;
	ref?: React.Ref<HTMLImageElement>;
}

export default ISvgSprite;
