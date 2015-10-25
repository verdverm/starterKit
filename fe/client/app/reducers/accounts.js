import PDB from '../pdb'

import * as ACCOUNTS from '../actions/accounts';


var initial_state = {
	linking: false,
	unlinking: false,
	retrieving: false,
	loading: false,
	saving: false,
	deleting: false,

	current: '',
	error: null,

	providers: {
		facebook: false,
		google: false,
		yahoo: false,
		github: false,
		twitter: false,
		soundcloud: false,
		spotify: false,
		dropbox: false,
	},
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
			obj.providers[action.provider] = true;
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
			obj.providers[action.provider] = false;
			return obj



		case ACCOUNTS.LOAD_SERVER_ACCOUNTS_STARTED:
			return Object.assign({}, state, {
				retrieving: true,
			});

		case ACCOUNTS.LOAD_SERVER_ACCOUNTS_FAILURE:
			return Object.assign({}, state, {
				retrieving: false,
				error: action.error
			});

		case ACCOUNTS.LOAD_SERVER_ACCOUNTS_SUCCESS:
			var obj = Object.assign({}, state, {
				retrieving: false,
			});
			for (var provider in action.providers) {
				console.log("Linked to", provider, action.providers[provider])
				obj.providers[provider] = action.providers[provider];
			};



		case ACCOUNTS.LOAD_ACCOUNT_STARTED:
			return Object.assign({}, state, {
				loading: true,
			});

		case ACCOUNTS.LOAD_ACCOUNT_FAILURE:
			return Object.assign({}, state, {
				loading: false,
				error: action.error,
			});

		case ACCOUNTS.LOAD_ACCOUNT_SUCCESS:
			return Object.assign({}, state, {
				loading: false,
				providers: action.providers,
			});



		case ACCOUNTS.SAVE_ACCOUNT_STARTED:
			return Object.assign({}, state, {
				saving: true,
			});

		case ACCOUNTS.SAVE_ACCOUNT_FAILURE:
			return Object.assign({}, state, {
				saving: false,
				error: action.error
			});

		case ACCOUNTS.SAVE_ACCOUNT_SUCCESS:
			return Object.assign({}, state, {
				saving: false,
			});



		case ACCOUNTS.DELETE_ACCOUNT_STARTED:
			return Object.assign({}, state, {
				deleting: true,
			});

		case ACCOUNTS.DELETE_ACCOUNT_FAILURE:
			return Object.assign({}, state, {
				deleting: false,
				error: action.error
			});

		case ACCOUNTS.DELETE_ACCOUNT_SUCCESS:
			return Object.assign({}, state, {
				deleting: false,
			});




		default:
			return state;
	}
}

export default accounts;
