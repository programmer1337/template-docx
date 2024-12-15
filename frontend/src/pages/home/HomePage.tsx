import './HomePage.scss';
import { TableLoader } from '../../widgets';

export function HomePage() {
	return (
		<>
			{/* <WidgetsExample data={{ text: 'Widgets Example' }} /> */}
			<div style={{ marginTop: '20px' }}>
				<TableLoader />
			</div>
		</>
	);
}
