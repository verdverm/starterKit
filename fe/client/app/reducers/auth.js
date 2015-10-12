import { LOGOUT_USER, 
		 LOGIN_USER_REQUEST, 
		 LOGIN_USER_SUCCESS, 
		 LOGIN_USER_FAILURE } 
	from '../actions/auth';

/*
 *  Pure functions on the state  (state,action) => state
 */

var starting_state = {
			attempting: false,
			authed: false,
			uid: '',
			username: '',
			token: '',
		}

function auth(state = starting_state, action) {
	switch (action.type) {

		case LOGIN_USER_REQUEST:
			console.log("Login Started")
			return Object.assign({}, state, {attempting: true});

		case LOGIN_USER_FAILURE:
			console.log("Login ERROR");
			return Object.assign({}, state, {attempting: false, error: action.error});


		case LOGIN_USER_SUCCESS:
			console.log("Login Success")
			return Object.assign({}, state, {
				attempting: false,
				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
			});

		case LOGOUT_USER:
			console.log("Logout user")
			return Object.assign({}, state, {
				authed: false,
				uid: '',
				username: '',
				token: '',
			});

		default:
			return state;
	}
}

export default auth;
