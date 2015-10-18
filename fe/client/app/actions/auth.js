import {
    actionCreator, optionsActionCreator
}
from 'redux-action-utils';

import routerActions from "./router"
import PDB from '../pdb'


export const SIGNUP_USER_STARTED = 'SIGNUP_USER_STARTED';
export const SIGNUP_USER_FAILURE = 'SIGNUP_USER_FAILURE';
export const SIGNUP_USER_SUCCESS = 'SIGNUP_USER_SUCCESS';
export const signupUserStarted = actionCreator(SIGNUP_USER_STARTED);
export const signupUserFailure = actionCreator(SIGNUP_USER_FAILURE, 'error');
export const signupUserSuccess = actionCreator(SIGNUP_USER_SUCCESS, 'uid', 'username', 'token');
export function signupUser(username, email, password, success_url = "profile", failure_url = "") {
    return function(dispatch) {

        dispatch(signupUserStarted());

        var payload = {
            username,
            email,
            password,
        }

        $.ajax({
            url: "http://localhost:8000/registration/",
            type: 'POST',
            dataType: 'json',
            data: payload,
            cache: false,
            success: function(resp) {
                var uid = "n/a";
                var token = resp.token;
                dispatch(signupUserSuccess(uid, username, token));
                dispatch(saveAuth(uid, username, token));

	            if (success_url !== "") {
	                dispatch(routerActions.stateGo(success_url));
	            }
            }.bind(this),
            error: function(xhr, status, error) {
                dispatch(signupUserFailure(error));
	            if (failure_url !== "") {
	                dispatch(routerActions.stateGo(failure_url));
	            }
            }.bind(this)
        });

    };
}


export const LOGIN_USER_STARTED = 'LOGIN_USER_STARTED';
export const LOGIN_USER_FAILURE = 'LOGIN_USER_FAILURE';
export const LOGIN_USER_SUCCESS = 'LOGIN_USER_SUCCESS';
export const loginUserStarted = actionCreator(LOGIN_USER_STARTED);
export const loginUserFailure = actionCreator(LOGIN_USER_FAILURE, 'error');
export const loginUserSuccess = actionCreator(LOGIN_USER_SUCCESS, 'uid', 'username', 'token');
export function loginUser(username, password, success_url = "profile", failure_url = "") {
    return function(dispatch) {

        dispatch(loginUserStarted());

        var payload = {
            username: username,
            password: password
        }

        $.ajax({
            url: "http://localhost:8000/token-auth/",
            type: 'POST',
            dataType: 'json',
            data: payload,
            cache: false,
            success: function(resp) {
                var uid = "n/a";
                var token = resp.token;
                dispatch(loginUserSuccess(uid, username, token));
                dispatch(saveAuth(uid, username, token));
	            if (success_url !== "") {
	                dispatch(routerActions.stateGo(success_url));
	            }
            }.bind(this),
            error: function(xhr, status, error) {
                dispatch(loginUserFailure(error));
	            if (failure_url !== "") {
	                dispatch(routerActions.stateGo(failure_url));
	            }
            }.bind(this)
        });

    };
}


export const TOKEN_TEST_STARTED = 'TOKEN_TEST_STARTED';
export const TOKEN_TEST_FAILURE = 'TOKEN_TEST_FAILURE';
export const TOKEN_TEST_SUCCESS = 'TOKEN_TEST_SUCCESS';
export const tokenTestStarted = actionCreator(TOKEN_TEST_STARTED);
export const tokenTestFailure = actionCreator(TOKEN_TEST_FAILURE, 'error');
export const tokenTestSuccess = actionCreator(TOKEN_TEST_SUCCESS, 'passed');
export function tokenTest(token) {
    return function(dispatch) {
        dispatch(loginUserStarted());

        var payload = {
            token
        }

        $.ajax({
            url: "http://localhost:8000/token-verify/",
            type: 'POST',
            dataType: 'json',
            data: payload,
            cache: false,
            success: function(resp) {
            	var passed = 'passed';
                dispatch(testTokenSuccess(passed));
            }.bind(this),
            error: function(xhr, status, error) {
                dispatch(testTokenFailure(error));
            }.bind(this)
        });

    };
}


export const LOGOUT_USER_STARTED = 'LOGOUT_USER_STARTED';
export const LOGOUT_USER_FAILURE = 'LOGOUT_USER_FAILURE';
export const LOGOUT_USER_SUCCESS = 'LOGOUT_USER_SUCCESS';
export const logoutUserStarted = actionCreator(LOGOUT_USER_STARTED);
export const logoutUserFailure = actionCreator(LOGOUT_USER_FAILURE, 'error');
export const logoutUserSuccess = actionCreator(LOGOUT_USER_SUCCESS);
export function logoutUser(rev, success_url = "home", failure_url = "") {
    return function(dispatch) {
        dispatch(logoutUserStarted());
        dispatch(deleteAuth(rev,success_url,failure_url));
        dispatch(logoutUserSuccess());
    }
}


export const LOAD_AUTH_STARTED = 'LOAD_AUTH_STARTED';
export const LOAD_AUTH_FAILURE = 'LOAD_AUTH_FAILURE';
export const LOAD_AUTH_SUCCESS = 'LOAD_AUTH_SUCCESS';
export const loadAuthStarted = actionCreator(LOAD_AUTH_STARTED);
export const loadAuthFailure = actionCreator(LOAD_AUTH_FAILURE, 'error');
export const loadAuthSuccess = actionCreator(LOAD_AUTH_SUCCESS, 'username', 'uid', 'token', 'rev');
export function loadAuth(success_url = "", failure_url = "") {
    return function(dispatch) {
        dispatch(loadAuthStarted())

        PDB.get('auth_data').then(function(doc) {

            dispatch(loadAuthSuccess(
                doc.username,
                doc.uid,
                doc.token,
                doc._rev
            ));

            if (success_url !== "") {
                dispatch(routerActions.stateGo(success_url));
            }
        }, function(error) {
            dispatch(loadAuthFailure(error));
            if (failure_url !== "") {
                dispatch(routerActions.stateGo(failure_url));
            }
        });
    };
}

export const SAVE_AUTH_STARTED = 'SAVE_AUTH_STARTED';
export const SAVE_AUTH_FAILURE = 'SAVE_AUTH_FAILURE';
export const SAVE_AUTH_SUCCESS = 'SAVE_AUTH_SUCCESS';
export const saveAuthStarted = actionCreator(SAVE_AUTH_STARTED);
export const saveAuthFailure = actionCreator(SAVE_AUTH_FAILURE, 'error');
export const saveAuthSuccess = actionCreator(SAVE_AUTH_SUCCESS, 'rev');
export function saveAuth(uid, username, token) {
    return function(dispatch) {

        dispatch(saveAuthStarted());

        var auth_item = {
            _id: "auth_data",
            uid,
            username,
            token,
        };

        PDB.put(auth_item, function callback(error, result) {
            if (!error) {
                dispatch(saveAuthSuccess(result.rev));
            } else {
                dispatch(saveAuthFailure(error));
            }
        });
    };
}


export const DELETE_AUTH_STARTED = 'DELETE_AUTH_STARTED';
export const DELETE_AUTH_FAILURE = 'DELETE_AUTH_FAILURE';
export const DELETE_AUTH_SUCCESS = 'DELETE_AUTH_SUCCESS';
export const deleteAuthStarted = actionCreator(DELETE_AUTH_STARTED);
export const deleteAuthFailure = actionCreator(DELETE_AUTH_FAILURE, 'error');
export const deleteAuthSuccess = actionCreator(DELETE_AUTH_SUCCESS);
export function deleteAuth(rev, success_url = "", failure_url = "") {
    return function(dispatch) {

        dispatch(deleteAuthStarted());

        var auth_item = {
            _id: "auth_data",
            _rev: rev,
        }
        console.log("deleting: ", auth_item);
        PDB.remove(auth_item, function callback(error, result) {
            if (!error) {
                dispatch(deleteAuthSuccess());
	            if (success_url !== "") {
	                dispatch(routerActions.stateGo(success_url));
	            }
            } else {
            	console.log("deleteAuth(error):", error);
                dispatch(deleteAuthFailure(error));
	            if (failure_url !== "") {
	                dispatch(routerActions.stateGo(failure_url));
	            }
            }
        });

    };
}
