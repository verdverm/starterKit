import PDB from '../pdb'

import { LOGOUT_USER, 
		 LOGIN_USER_REQUEST, 
		 LOGIN_USER_SUCCESS, 
		 LOGIN_USER_FAILURE,
		 LOADED_USER_AUTH,  
		 LOADED_USER_AUTH_BEGIN, 
		 LOADED_USER_AUTH_FAIL 
} from '../actions/auth';

/*
 *  Pure functions on the state  (state,action) => state
 */

var initial_state = {
			attempting: false,
			loading: false,
			saving: false,
			
			authed: false,
			token: '',
			uid: '',
			username: '',
		}

function auth(state = initial_state, action) {
	switch (action.type) {

		case LOGIN_USER_REQUEST:
			console.log("Login Started")
			return Object.assign({}, state, {attempting: true});

		case LOGIN_USER_FAILURE:
			console.log("Login ERROR");
			return Object.assign({}, state, {attempting: false, error: action.error});


		case LOGIN_USER_SUCCESS:
			console.log("Login Success")

			var auth_item = {
			    _id: "auth_data",
			    uid: action.uid,
				username: action.username,
				token: action.token,
			};

			// this needs to be handled differently
			PDB.put(auth_item, function callback(err, result) {
			    if (!err) {
			        console.log('Successfully posted auth data!');
			        console.log(result);

			        // would really like to emit an event and handle here too
			  //       return Object.assign({}, state, {
					// 	attempting: false,
					// 	authed: true,
					// 	uid: action.uid,
					// 	username: action.username,
					// 	token: action.token,
					// 	rev: result.rev,
					// });
			    } else {
			    	console.log("error posting auth to PouchDB", err);
			    	alert("error posting auth to PouchDB");
					// return Object.assign({}, state, {attempting: false, error: err});
			    }
			});
			return Object.assign({}, state, {
				attempting: false,
				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
			});

		case LOADED_USER_AUTH_BEGIN:
			console.log("Loading Auth Started")
			return Object.assign({}, state, {loading: true});

		case LOADED_USER_AUTH_FAIL:
			console.log("Loading Auth ERROR");
			return Object.assign({}, state, {loading: false, error: action.error});


		case LOADED_USER_AUTH:
			console.log("loading user auth")
	        return Object.assign({}, state, {
				loading: false,
				authed: true,
				uid: action.uid,
				username: action.username,
				token: action.token,
				rev: action.rev,
			});


		case LOGOUT_USER:
			console.log("Logout user", state)

			var del_auth_item = {
			    _id: "auth_data",
			    _rev: state.rev
			}
			PDB.remove(del_auth_item, function callback(err, result) {
			    if (!err) {
			        console.log('Successfully removed auth data!');
			        console.log(result);
			    } else {
			    	console.log("error removing auth to PouchDB", err);
			    	alert("error removing auth to PouchDB");
			    }
			});

			return Object.assign({}, state, {
				authed: false,
				uid: '',
				username: '',
				token: '',
				rev: '',
			});

		default:
			return state;
	}
}

export default auth;
