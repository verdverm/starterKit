import { combineReducers } from 'redux';
import {router} from 'redux-ui-router';

import auth from './auth';
import accounts from './accounts';


const rootReducer = combineReducers({
	router,
	auth,
	accounts,
});

export default rootReducer;
