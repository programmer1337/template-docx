import ICounterpartiesTable, { ICounterparty } from './CounterpartiesTable.types';
import './CounterpartiesTable.scss';
import classes, { classes2 } from './CounterpartiesTable.cn';
import { Button, SvgSprite } from 'src/shared/ui';
import { useState } from 'react';

function fetchReplace(body: Counterparty[]) {
	fetch(`http://${process.env.REACT_APP_API_URL}/api/replace`, {
		method: 'POST',
		body: JSON.stringify(body),
		headers: {
			'Content-Type': 'application/json',
		},
	})
		.then((resp) => {
			if (resp.status === 200) {
				return resp.blob();
			}
			throw new Error(`Ошибка: ${resp.status}`);
		})
		.then((blob) => {
			// saveAs(blob, 'all_files.zip');
			if (blob) {
				// Если не используем FileSaver, можно создать ссылку вручную:
				const link = document.createElement('a');
				link.href = URL.createObjectURL(blob);
				link.download = 'all_files.zip';
				link.click();
			}
		});
}

export function CounterpartiesTable(props: ICounterpartiesTable) {
	const { data, cn } = props;
	const { counterparties } = data;
	const styles = classes({ ...cn });

	const handleClick = () => {
		if (!counterparties) return;

		const body = counterparties.slice(1);
		body.forEach((item) => {
			for (const key in item) {
				item[key as CounterpartyKeys] = item[key as CounterpartyKeys].toString();
			}
		});

		fetchReplace(body);
	};

	if (counterparties) {
		return (
			<div className={styles.block}>
				<div className={styles.elementTable}>
					{counterparties.slice(1).map((item, pos) => (
						<Counterparty key={pos} data={{ counterparty: item }} />
					))}
				</div>
				<Button
					data={{ text: 'Сформировать договора' }}
					cn={{ margin: 'mt-24', width: 'w-100' }}
					handler={{ onClick: handleClick }}
				/>
			</div>
		);
	}
}

const counterpartyFieldAlias: Record<string, string> = {
	code_ou: 'Код ОУ',
	inn: 'Инн',
	institution_short_name: 'Сокращенное наименование учреждения',
	institution_full_name: 'Полное наименование',
	address: 'Юридический и почтовый адрес образовательного учреждения',
	city: 'Район',
	bank_details: 'Банковские реквизиты',
	responsible_person_job_title: 'Должность подписанта (им.падеж)',
	responsible_person_short_name: 'ФИО подписанта (сокр.)',
	responsible_person_full_name: 'ФИО подписанта полностью (им. Падеж)',
	responsible_person_full_name_genitive: 'ФИО подписанта полностью (род. падеж)',
	acting_on: 'Действующего на основании',
	ikz_2025: 'ИКЗ 2025',
	source_funding: 'Источник финансирования',
	email: 'E-mail',
	phone_number: 'Телефон',
	contract_form: 'Форма договора',
	contract_type: 'Тип договора',
	contract_number: 'Номер договора',
	contract_formation_data: 'Дата формирования договора',
	responsible_person_job_title_genetive: 'Должность подписанта (род. Падеж)',
	category: 'Категория',
};

function Counterparty(props: ICounterparty) {
	const { data, cn } = props;
	const { counterparty } = data;
	const styles = classes2(cn);

	const [isBodyShown, setBodyShow] = useState(false);
	const handleHeadClick = () => {
		setBodyShow((prev) => !prev);
	};

	const counterpartyKey = Object.keys(counterparty);

	return (
		<div className={`${styles.block} ${isBodyShown ? 'js-body-show' : ''}`.trim()}>
			<div className={styles.elementHead} onClick={handleHeadClick}>
				<div className={styles.elementText}>
					{/* <div>Код ОУ {counterparty.code_ou}</div> */}
					<div>ИНН {counterparty.inn}</div>
				</div>
				<SvgSprite data={{ name: 'down-shevron' }} className={styles.elementIcon} />
			</div>
			<div className={styles.elementBody}>
				<div className={styles.elementAnimation}>
					{counterpartyKey.map((key) => (
						<div key={key} className={styles.elementCharacteristic}>
							{counterpartyFieldAlias[key]}: {counterparty[key as keyof typeof Counterparty]}
						</div>
					))}
				</div>
			</div>
		</div>
	);
}
