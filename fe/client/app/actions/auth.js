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
