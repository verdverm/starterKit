import { combineReducers } from 'redux';
import {router} from 'redux-ui-router';
// import { persistentReducer } from 'redux-pouchdb';

import error from './error';
import auth from './auth';

// let persistantAuth = persistentReducer(auth);
// console.log(auth);
// console.log(persistantAuth);


const rootReducer = combineReducers({
	router,
	error,
	auth
});

export default rootReducer;
