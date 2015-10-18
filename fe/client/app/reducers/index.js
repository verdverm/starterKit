import { combineReducers } from 'redux';
import {router} from 'redux-ui-router';

import auth from './auth';


const rootReducer = combineReducers({
	router,
	auth
});

export default rootReducer;
