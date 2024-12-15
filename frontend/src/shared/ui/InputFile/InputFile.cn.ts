import CN from '../../lib/ClassBem';

export default function classes() {
	const BLOCK = CN('input-file');

	return {
		block: BLOCK({}),
		elementField: BLOCK('field'),
		elementFakeButton: BLOCK('fake-button'),
	};
}
