// Asynchronous Action Creators

import routerActions from "./router"

export function registerUser(username,email,password) {
	return function (dispatch) {

		dispatch(loginUserStarted());

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
				console.log("Logged in: ", resp)
				dispatch(loginUserSuccess(username,resp.token))
			}.bind(this),
			error: function(xhr, status, err) {
				dispatch(loginUserFailure(err));
			}.bind(this)
		});

	};
}

export const LOADED_USER_AUTH = 'LOADED_USER_AUTH';
export function loadedUserAuth(username, uid, token, rev) {
	return {
		type: LOADED_USER_AUTH,
		username,
		token,
		uid,
		rev
	}
}

export const LOADED_USER_AUTH_BEGIN = 'LOADED_USER_AUTH_BEGIN';
export function loadedUserAuthBegin() {
	return {
		type: LOADED_USER_AUTH_BEGIN,
	}
}
export const LOADED_USER_AUTH_FAIL = 'LOADED_USER_AUTH_FAIL';
export function loadedUserAuthFail(error) {
	return {
		type: LOADED_USER_AUTH_FAIL,
		error
	}
}

export function loginUser(username,password) {
	return function (dispatch) {

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
				console.log("Logged in: ", resp)

				dispatch(loginUserSuccess(username,resp.token))
				dispatch(routerActions.stateGo('profile'))
			}.bind(this),
			error: function(xhr, status, err) {
				dispatch(loginUserFailure(err));
			}.bind(this)
		});

	};
}

// Synchronous Action Creators
export const LOGIN_USER_STARTED = 'LOGIN_USER_STARTED';
export function loginUserStarted() {
	return {
		type: LOGIN_USER_STARTED,
	}
}

export const LOGIN_USER_FAILURE = 'LOGIN_USER_FAILURE';
export function loginUserFailure(error) {
	return {
		type: LOGIN_USER_FAILURE,
		error,
	}
}

export const LOGIN_USER_SUCCESS = 'LOGIN_USER_SUCCESS';
export function loginUserSuccess(username, token) {
	return {
		type: LOGIN_USER_SUCCESS,
		username,
		token,
	}
}

export const LOGOUT_USER = 'LOGOUT_USER';
export function logoutUser() {
	return {
		type: LOGOUT_USER,
	}
}
