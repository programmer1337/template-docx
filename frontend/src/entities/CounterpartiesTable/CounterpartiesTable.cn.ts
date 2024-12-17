import CN from '../../shared/lib/ClassBem';
import ICounterpartiesTable, { ICounterparty } from './CounterpartiesTable.types';

export default function classes(cn: ICounterpartiesTable['cn']) {
	const BLOCK = CN('counterparties-table');

	return {
		block: BLOCK({}, [cn?.padding]),
		elementTable: BLOCK('table'),
		elementItem: BLOCK('item'),
		elementCounterparty: BLOCK('counterparty'),
		elementHead: BLOCK('head'),
		elementBody: BLOCK('body'),
	};
}

export function classes2(cn?: ICounterparty['cn']) {
	const BLOCK = CN('counterparty');

	return {
		block: BLOCK({}, [cn?.margin]),
		elementHead: BLOCK('head'),
		elementText: BLOCK('text'),
		elementIcon: BLOCK('icon'),
		elementBody: BLOCK('body'),
		elementAnimation: BLOCK('animation'),
		elementCharacteristic: BLOCK('characteristic'),
	};
}
