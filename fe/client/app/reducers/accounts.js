import PDB from '../pdb'

import * as ACCOUNTS from '../actions/accounts';


var initial_state = {
	facebook: false,
	google: false,

	current: '',
	linking: false,
	unlinking: false,
	error: null
}


function accounts(state = initial_state, action) {
	switch (action.type) {

		case ACCOUNTS.LINKING_ACCT_STARTED:
			return Object.assign({}, state, {
				linking: true,
				current: action.provider,
			});

		case ACCOUNTS.LINKING_ACCT_FAILURE:
			return Object.assign({}, state, {
				linking: false,
				error: action.error
			});

		case ACCOUNTS.LINKING_ACCT_SUCCESS:
			var obj = Object.assign({}, state, {
				linking: false,
				current: '',
			});
			obj[action.provider] = true;
			return obj


		case ACCOUNTS.UNLINKING_ACCT_STARTED:
			return Object.assign({}, state, {
				unlinking: true,
				current: action.provider,
			});

		case ACCOUNTS.UNLINKING_ACCT_FAILURE:
			return Object.assign({}, state, {
				unlinking: false,
				error: action.error
			});

		case ACCOUNTS.UNLINKING_ACCT_SUCCESS:
			var obj = Object.assign({}, state, {
				unlinking: false,
				current: '',
			});
			obj[action.provider] = false;
			return obj






		default:
			return state;
	}
}

export default accounts;
