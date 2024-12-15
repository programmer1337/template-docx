import IInputFile from './InputFile.types';
import classes from './InputFile.cn';
import './InputFile.scss';

export function InputFile(props: IInputFile) {
	const { data, executor } = props;
	const { file } = data;
	const styles = classes();

	const handleChange = (elem: any) => {
		executor.setFile(elem.target.files[0]);
	};

	return (
		<label className={styles.block}>
			<input type="file" className={styles.elementField} onChange={handleChange} />
			{file && <div>{file.name}</div>}
			<div className={styles.elementFakeButton}>{file ? 'Заменить файл' : 'Выберите файл'}</div>
		</label>
	);
}
