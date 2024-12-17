import ICounterpartiesTable from './CounterpartiesTable.types';
import './CounterpartiesTable.scss';
import classes from './CounterpartiesTable.cn';
import { Button } from 'src/shared/ui';

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

		// eslint-disable-next-line no-console
		console.log(process.env.REACT_APP_API_URL);

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
	};

	if (counterparties) {
		return (
			<div className={styles.block}>
				<table className={styles.elementTable}>
					<tbody>
						{counterparties.map((item, pos) => (
							<tr className={styles.elementItem} key={pos}>
								<td>{item.code_ou}</td>
								<td>{item.inn}</td>
								<td>{item.institution_short_name}</td>
								<td>{item.address}</td>
								<td>{item.city}</td>
								<td>{item.responsible_person_job_title}</td>
								<td>{item.responsible_person_full_name}</td>
							</tr>
						))}
					</tbody>
				</table>
				<Button
					data={{ text: 'Сформировать договора' }}
					cn={{ margin: 'mt-24', width: 'w-100' }}
					handler={{ onClick: handleClick }}
				/>
			</div>
		);
	}
}
