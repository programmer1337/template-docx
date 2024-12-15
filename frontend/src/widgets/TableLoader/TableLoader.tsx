import ITableLoader from './TableLoader.types';
import './TableLoader.scss';
import classes from './TableLoader.cn';
import { XlsxLoader } from 'src/features';
import { Container } from 'src/shared/ui';
import { CounterpartiesTable } from 'src/entities/CounterpartiesTable/CounterpartiesTable';
import { useState } from 'react';

export function TableLoader(props: ITableLoader) {
	const { cn } = props;
	const styles = classes({ ...cn });

	const [counterparties, setCounterparties] = useState<Counterparties | null>(null);
	const handleCounterparties = (counterparties: any) => {
		setCounterparties(counterparties);
	};

	return (
		<section className={styles.block}>
			<Container>
				<XlsxLoader executor={{ setCounterparties: handleCounterparties }} />
				<CounterpartiesTable data={{ counterparties: counterparties }} />
			</Container>
		</section>
	);
}
