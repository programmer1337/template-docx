import './XlsxLoader.scss';
import classes from './XlsxLoader.cn';
import { InputFile } from 'src/shared/ui/InputFile/InputFile';
import { Button } from 'src/shared/ui';
import { useState } from 'react';
import { read, utils } from 'xlsx';
import IXlsxLoader from './XlsxLoader.types';

const readExcel = (setValue: (counterparties: Counterparties) => void, file: File) => {
	(async () => {
		const data = await file.arrayBuffer();
		const workbook = read(data, { type: 'binary' });

		const first_sheet_name = workbook.SheetNames[0];

		/* Get worksheet */
		const worksheet = workbook.Sheets[first_sheet_name];

		const headers = [
			'code_ou',
			'inn',
			'institution_short_name',
			'institution_full_name',
			'address',
			'city',
			'bank_details',
			'responsible_person_job_title',
			'responsible_person_short_name',
			'responsible_person_full_name',
			'responsible_person_full_name_genitive',
			'acting_on',
			'ikz_2025',
			'source_funding',
			'email',
			'phone_number',
			'contract_form',
			'contract_type',
			'contract_number',
			'contract_formation_data',
			'responsible_person_job_title_genetive',
			'category',
		];
		const contragents: Counterparties = utils.sheet_to_json(worksheet, {
			range: worksheet['!ref'],
			header: headers,
			defval: '',
		});
		setValue(contragents);
	})();
};

export function XlsxLoader(props: IXlsxLoader) {
	const { executor } = props;
	const styles = classes();
	const [file, setFile] = useState<File | null>(null);

	const handleClick = () => {
		if (file && executor?.setCounterparties) {
			readExcel(executor.setCounterparties, file);
		}
	};

	return (
		<div className={styles.block}>
			<InputFile data={{ file: file }} executor={{ setFile: setFile }} />
			<Button data={{ text: 'Загрузить' }} handler={{ onClick: handleClick }} />
		</div>
	);
}
