import PDB from '../pdb'

import * as AUTH from '../actions/auth';


var initial_state = {
	// these corrispond to state values 
	// for processing auth actions
	// (true means running)
	signup_ing: false,
	login_ing: false,
	test_ing: false,
	logout_ing: false,
	load_ing: false,
	save_ing: false,
	delete_ing: false,
	
	// these are the actual auth state
	authed: false,
	token: '',
	uid: '',
	username: '',
	error: null,
}

function auth(state = initial_state, action) {
	switch (action.type) {

		case AUTH.SIGNUP_USER_STARTED:
			return Object.assign({}, state, {
				signup_ing: true
			});

		case AUTH.SIGNUP_USER_FAILURE:
			return Object.assign({}, state, {
				signup_ing: false, 
				error: action.error
			});

		case AUTH.SIGNUP_USER_SUCCESS:

			return Object.assign({}, state, {
				signup_ing: false,

				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
			});


		case AUTH.LOGIN_USER_STARTED:
			return Object.assign({}, state, {
				login_ing: true
			});

		case AUTH.LOGIN_USER_FAILURE:
			return Object.assign({}, state, {
				login_ing: false, 
				error: action.error
			});

		case AUTH.LOGIN_USER_SUCCESS:
			return Object.assign({}, state, {
				login_ing: false,

				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
			});


		case AUTH.TOKEN_TEST_STARTED:
			return Object.assign({}, state, {
				test_ing: true
			});

		case AUTH.TOKEN_TEST_FAILURE:
			return Object.assign({}, state, {
				test_ing: false, 
				error: action.error
			});

		case AUTH.TOKEN_TEST_SUCCESS:
			return Object.assign({}, state, {
				test_ing: false,
			});


		case AUTH.LOGOUT_USER_STARTED:
			return Object.assign({}, state, {
				logout_ing: true
			});

		case AUTH.LOGOUT_USER_FAILURE:
			return Object.assign({}, state, {
				logout_ing: false, 
				error: action.error
			});

		case AUTH.LOGOUT_USER_SUCCESS:
			return Object.assign({}, state, {
				logout_ing: false,

				authed: false,
				uid: '',
				username: '',
				token: '',
			});


		case AUTH.LOAD_AUTH_STARTED:
			return Object.assign({}, state, {
				load_ing: true
			});

		case AUTH.LOAD_AUTH_FAILURE:
			if (action.error.status === 404) {
				action.error = null;
			}
			return Object.assign({}, state, {
				load_ing: false, 
				error: action.error
			});

		case AUTH.LOAD_AUTH_SUCCESS:
			return Object.assign({}, state, {
				load_ing: false,
				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
				rev: action.rev
			});


		case AUTH.SAVE_AUTH_STARTED:
			return Object.assign({}, state, {
				save_ing: true
			});

		case AUTH.SAVE_AUTH_FAILURE:
			return Object.assign({}, state, {
				save_ing: false, 
				error: action.error
			});

		case AUTH.SAVE_AUTH_SUCCESS:
			return Object.assign({}, state, {
				save_ing: false,
				rev: action.rev
			});


		case AUTH.DELETE_AUTH_STARTED:
			return Object.assign({}, state, {
				delete_ing: true
			});

		case AUTH.DELETE_AUTH_FAILURE:
			return Object.assign({}, state, {
				delete_ing: false, 
				error: action.error
			});

		case AUTH.DELETE_AUTH_SUCCESS:
			return Object.assign({}, state, {
				delete_ing: false,
			});



		default:
			return state;
	}
}

export default auth;
