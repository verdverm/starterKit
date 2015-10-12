import { combineReducers } from 'redux';
import {router} from 'redux-ui-router';

import error from './error';
import auth from './auth';


const rootReducer = combineReducers({
	router,
	error,
	auth
});

export default rootReducer;
