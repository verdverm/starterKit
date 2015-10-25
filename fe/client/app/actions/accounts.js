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

		var $injector = angular.injector(['ng','satellizer']);

		$injector.invoke(function($auth){
			var real_provider = provider;
			if (real_provider === "windows") {
	    		real_provider = "live"
	    	}
	        $auth.authenticate(real_provider)
	        .then(function(response) {

	            dispatch( linkAccountSuccess(provider) );
                // dispatch(saveAccount(resp))

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
                // dispatch(saveAccount(resp))
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


export const LOAD_SERVER_ACCOUNTS_STARTED = 'LOAD_SERVER_ACCOUNTS_STARTED';
export const LOAD_SERVER_ACCOUNTS_FAILURE = 'LOAD_SERVER_ACCOUNTS_FAILURE';
export const LOAD_SERVER_ACCOUNTS_SUCCESS = 'LOAD_SERVER_ACCOUNTS_SUCCESS';
export const loadServerAccountsStarted = actionCreator(LOAD_SERVER_ACCOUNTS_STARTED);
export const loadServerAccountsFailure = actionCreator(LOAD_SERVER_ACCOUNTS_FAILURE, 'error');
export const loadServerAccountsSuccess = actionCreator(LOAD_SERVER_ACCOUNTS_SUCCESS, 'providers');
export function loadServerAccounts() {
    return function(dispatch) {
        dispatch(loadServerAccountsStarted());

		var $injector = angular.injector(['ng','satellizer']);

		$injector.invoke(function($auth){

	        $.ajax({
	            url: "http://localhost:8000/auth/accounts/",
	            type: 'GET',
	            dataType: 'json',
	            headers: {
	            	'Authorization': 'Bearer ' + $auth.getToken(),
	            },
	            cache: false,
	            success: function(resp) {
	            	console.log("loadServerAccounts: ", resp);
	                dispatch(loadServerAccountsSuccess(resp));
	                dispatch(saveAccount(resp))
	            }.bind(this),
	            error: function(xhr, status, error) {
	                dispatch(loadServerAccountsFailure(error));
	            }.bind(this)
	        });
	    });
    };
}


export const LOAD_ACCOUNT_STARTED = 'LOAD_ACCOUNT_STARTED';
export const LOAD_ACCOUNT_FAILURE = 'LOAD_ACCOUNT_FAILURE';
export const LOAD_ACCOUNT_SUCCESS = 'LOAD_ACCOUNT_SUCCESS';
export const loadAccountStarted = actionCreator(LOAD_ACCOUNT_STARTED);
export const loadAccountFailure = actionCreator(LOAD_ACCOUNT_FAILURE, 'error');
export const loadAccountSuccess = actionCreator(LOAD_ACCOUNT_SUCCESS, 'providers', 'rev');
export function loadAccount(success_url = "", failure_url = "") {
    return function(dispatch) {
        dispatch(loadAccountStarted())

        PDB.get('account_data').then(function(doc) {

            dispatch(loadAccountSuccess(
			 	providers,
                doc._rev
            ));

            if (success_url !== "") {
                dispatch(routerActions.stateGo(success_url));
            }
        }, function(error) {
            dispatch(loadAccountFailure(error));
            if (failure_url !== "") {
                dispatch(routerActions.stateGo(failure_url));
            }
        });
    };
}

export const SAVE_ACCOUNT_STARTED = 'SAVE_ACCOUNT_STARTED';
export const SAVE_ACCOUNT_FAILURE = 'SAVE_ACCOUNT_FAILURE';
export const SAVE_ACCOUNT_SUCCESS = 'SAVE_ACCOUNT_SUCCESS';
export const saveAccountStarted = actionCreator(SAVE_ACCOUNT_STARTED);
export const saveAccountFailure = actionCreator(SAVE_ACCOUNT_FAILURE, 'error');
export const saveAccountSuccess = actionCreator(SAVE_ACCOUNT_SUCCESS, 'rev');
export function saveAccount(providers) {
    return function(dispatch) {

        dispatch(saveAccountStarted());

        var account_item = {
            _id: "account_data",
        	providers,
        };

        PDB.put(account_item, function callback(error, result) {
            if (!error) {
                dispatch(saveAccountSuccess(result.rev));
            } else {
                dispatch(saveAccountFailure(error));
            }
        });
    };
}


export const DELETE_ACCOUNT_STARTED = 'DELETE_ACCOUNT_STARTED';
export const DELETE_ACCOUNT_FAILURE = 'DELETE_ACCOUNT_FAILURE';
export const DELETE_ACCOUNT_SUCCESS = 'DELETE_ACCOUNT_SUCCESS';
export const deleteAccountStarted = actionCreator(DELETE_ACCOUNT_STARTED);
export const deleteAccountFailure = actionCreator(DELETE_ACCOUNT_FAILURE, 'error');
export const deleteAccountSuccess = actionCreator(DELETE_ACCOUNT_SUCCESS);
export function deleteAccount(rev, success_url = "", failure_url = "") {
    return function(dispatch) {

        dispatch(deleteAccountStarted());

        var account_item = {
            _id: "account_data",
            _rev: rev,
        }
        console.log("deleting: ", account_item);
        PDB.remove(account_item, function callback(error, result) {
            if (!error) {
                dispatch(deleteAccountSuccess());
	            if (success_url !== "") {
	                dispatch(routerActions.stateGo(success_url));
	            }
            } else {
            	console.log("deleteAccount(error):", error);
                dispatch(deleteAccountFailure(error));
	            if (failure_url !== "") {
	                dispatch(routerActions.stateGo(failure_url));
	            }
            }
        });

    };
}
