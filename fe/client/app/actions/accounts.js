import {actionCreator, optionsActionCreator} from 'redux-action-utils';

import angular from 'angular';
import authProvider from 'satellizer';

import routerActions from "./router"
import PDB from '../pdb'


export const LINKING_ACCT_STARTED = 'LINKING_ACCT_STARTED';
export const LINKING_ACCT_FAILURE = 'LINKING_ACCT_FAILURE';
export const LINKING_ACCT_SUCCESS = 'LINKING_ACCT_SUCCESS';
export const linkAccountStarted = actionCreator(LINKING_ACCT_STARTED, 'provider');
export const linkAccountFailure = actionCreator(LINKING_ACCT_FAILURE, 'error');
export const linkAccountSuccess = actionCreator(LINKING_ACCT_SUCCESS, 'provider');

export function linkToProvider(provider, success_url = "profile", failure_url = "") {
    return function(dispatch) {

        dispatch( linkAccountStarted(provider) );

        // var injector = angular.element(document.querySelector('[ng-app]')).injector(['satellizer']);
        // var auth = injector.get("auth");

		// angular.injector(['ng', 'satellizer']).get("auth")
        // authProvider

		var $injector = angular.injector(['ng','satellizer']);

		$injector.invoke(function($auth){

	        $auth.authenticate(provider)
	        .then(function(response) {

	            dispatch( linkAccountSuccess(provider) );

	            var res_str = 'You have successfully linked with ' + provider;
	            Materialize.toast(res_str, 4000);

	            console.log(response);

	        })
	        .catch(function(response) {

	            dispatch( linkAccountFailure(response.data.message) );
	            console.log(response);
	            var err_str = "Error: " + response.data.message;
	            Materialize.toast(err_str, 4000)
	        });

		});
    };
}


export const UNLINKING_ACCT_STARTED = 'UNLINKING_ACCT_STARTED';
export const UNLINKING_ACCT_FAILURE = 'UNLINKING_ACCT_FAILURE';
export const UNLINKING_ACCT_SUCCESS = 'UNLINKING_ACCT_SUCCESS';
export const unlinkAccountStarted = actionCreator(UNLINKING_ACCT_STARTED, 'provider');
export const unlinkAccountFailure = actionCreator(UNLINKING_ACCT_FAILURE, 'error');
export const unlinkAccountSuccess = actionCreator(UNLINKING_ACCT_SUCCESS, 'provider');

export function unlinkFromProvider(provider, success_url = "profile", failure_url = "") {
    return function(dispatch) {

        dispatch( unlinkAccountStarted(provider) );

		var $injector = angular.injector(['ng','satellizer']);
		$injector.invoke(function($auth){

	        $auth.unlink(provider)
	        .then(function(response) {

	            dispatch( unlinkAccountSuccess(provider) );
	            console.log(response);

	            var res_str = 'You have successfully unlinked from ' + provider;
	            Materialize.toast(res_str, 4000);
	        })
	        .catch(function(response) {

	            dispatch( unlinkAccountFailure(response.data.message) );
	            console.log(response);

	            var err_str = "Error: " + response.data.message;
	            Materialize.toast(err_str, 4000)
	        });

		});
    };
}
