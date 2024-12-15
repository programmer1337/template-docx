import CN from '../../shared/lib/ClassBem';
import ITableLoader from './TableLoader.types';

export default function classes(cn: ITableLoader['cn']) {
	const BLOCK = CN('table-loader');

	return {
		block: BLOCK({}, [cn?.padding]),
	};
}
