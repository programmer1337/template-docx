import ISvgSprite from './SvgSprite.types';
import SvgSpriteIcon from './SvgSprite.svg?url';

export function SvgSprite(props: ISvgSprite) {
	const { data, className } = props;

	return (
		<svg className={className} width={data.width} height={data.height}>
			<use href={`${SvgSpriteIcon}#${data.name}`} />
		</svg>
	);
}
