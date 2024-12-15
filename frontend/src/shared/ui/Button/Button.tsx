import IButton from './Button.types';
import classes from './Button.cn';
import './Button.scss';

export function Button(props: IButton) {
	const { data, cn, handler } = props;
	const styles = classes({ ...cn });

	return (
		<button className={styles.block} onClick={handler?.onClick}>
			{data.text}
		</button>
	);
}
