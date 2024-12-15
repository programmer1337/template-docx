import CN from '../../shared/lib/ClassBem';
import ICounterpartiesTable from './CounterpartiesTable.types';

export default function classes(cn: ICounterpartiesTable['cn']) {
	const BLOCK = CN('counterparties-table');

	return {
		block: BLOCK({}, [cn?.padding]),
		elementTable: BLOCK('table'),
		elementItem: BLOCK('item'),
	};
}
