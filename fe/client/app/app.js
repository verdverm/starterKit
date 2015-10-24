import angular from 'angular';
import ngTouch from 'angular-touch';
import Permission from 'angular-permission';
// import authProvider from 'satellizer';

import ngRedux from 'ng-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';
import { devTools } from 'redux-devtools';

import uiRouter from 'angular-ui-router';
import ngReduxRouter from 'redux-ui-router';

import rootReducer from './reducers';

import Components from './components/components';
import AppComponent from './app.component';

import { loadAuth } from './actions/auth';
import routerActions from "./actions/router"
import accountsActions from "./actions/accounts"


angular.module('app', [
    ngTouch,
    ngRedux,

    uiRouter,
    ngReduxRouter,

    'permission',

    Components.name
])

.config(['$ngReduxProvider', ($ngReduxProvider) => {
    console.log("Config'n App");

    const logger = createLogger({
        level: 'info',
        collapsed: true
    });

    $ngReduxProvider.createStoreWith(
        rootReducer,
        ['ngUiRouterMiddleware', thunk, logger] 
        ,[devTools()]
    );
}])

.config([ '$authProvider', ($authProvider) => {

    $authProvider.unlinkUrl = 'http://localhost:8000/auth/unlink/';

    $authProvider.facebook({
        url: 'http://localhost:8000/auth/facebook/',
        clientId: '855012704576907',
        scope: ['email','public_profile'],
    });

    $authProvider.google({
        url: 'http://localhost:8000/auth/google/',
        clientId: '87612612394-3uq02vaa8drdkmsoeu43c8hrfq665oin.apps.googleusercontent.com'
    });

    $authProvider.github({
        url: 'http://localhost:8000/auth/github/',
        clientId: 'GitHub Client ID'
    });

    $authProvider.linkedin({
        url: 'http://localhost:8000/auth/linkedin/',
        clientId: 'LinkedIn Client ID'
    });

    $authProvider.twitter({
        url: 'http://localhost:8000/auth/twitter/',
        clientId: 'Twitter Client ID'
    });
}])

.run([ 'Permission', '$ngRedux', '$rootScope', '$state', '$location', '$http', '$auth', 
    (Permission, $ngRedux, $rootScope, $state, $location, $http, $auth) => {
    // grab local copy of redux and routerState for closures
    let localRedux = $ngRedux;

    let url = $location.url();
    console.log("URL: ", url);

    let states = $state.get();
    for (var i = states.length - 1; i >= 0; i--) {
        if (states[i].url === url) {
            url = states[i].name;
        }
    };

    $ngRedux.subscribe(() => {
        let state = $ngRedux.getState();
        let token = state.auth.token;
        // console.log("HTTP token checker: ", token, $http.defaults.headers.common.Authorization)
        if (token !== "" && $http.defaults.headers.common.Authorization === undefined) {
            console.log("Setting HTTP token to: ", token)
            $http.defaults.headers.common.Authorization = "Bearer " + token;
        } else if (token === "" && $http.defaults.headers.common.Authorization !== undefined) {
            console.log("Deleting HTTP token")
            delete $http.defaults.headers.common.Authorization;
        }
        // if (token !== "" && $auth.getToken() !== token) {
        //     console.log("Setting AUTH token to: ", token)
        //     $auth.setToken(token);
        // } else if (token === "" && $auth.getToken()) {
        //     console.log("Deleting AUTH token")
        //     delete $auth.removeToken();
        // }
    })


    let first = true;
    $rootScope.$on('$stateChangeSuccess',
        function(event, toState, toParams, fromState, fromParams) {
            if ( first == true ) {
                first = false;

                var pass_url = url;
                var fail_url = pass_url;
                if (pass_url === "login" || pass_url === "register") {
                    pass_url = 'profile';
                }
                
                localRedux.dispatch( loadAuth(pass_url, fail_url) );        

            }
        }
    );

    // Configure permissions on the UI
    // Define anonymous role
    Permission.defineRole('anonymous', function() {
        let state = localRedux.getState();

        if (state.auth.authed === true) {
            return false;
        }
        return true;
    });

    // Define user role calling back-end
    Permission.defineRole('user', function() {
        let state = localRedux.getState();

        return state.auth.authed === true;
    });

    // A different example for admin
    Permission.defineRole('admin', function() {
        let state = localRedux.getState();

        return state.auth.admin === true;
    });


    // determine where we should redirect after successful auth load
    // if we are successful, go where ever, unless authed and trying to
    // if failure, go where ever too, will get redirected to login if we can

    console.log("Runnin' App")
    
}])

.directive('app', AppComponent);
