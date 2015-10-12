import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngTouch from 'angular-touch';

import ngRedux from 'ng-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';
import {
    devTools
}
from 'redux-devtools';

import ngReduxRouter from 'redux-ui-router';

import Permission from 'angular-permission'


import rootReducer from './reducers';

import Common from './common/common';
import Components from './components/components';
import AppComponent from './app.component';


angular.module('app', [
    ngTouch,
    ngRedux,

    uiRouter,
    ngReduxRouter,

    'permission',

    Common.name,
    Components.name
])

.config(($ngReduxProvider) => {
    console.log("Config'n App")

    const logger = createLogger({
        level: 'info',
        collapsed: true
    });

    $ngReduxProvider.createStoreWith(
        rootReducer, ['ngUiRouterMiddleware', thunk, logger], [devTools()]
    );
})

.run(function(Permission, $ngRedux) {
    console.log("Permission:", Permission)
    console.log($ngRedux)
    // Configure permissions on the UI

    let localRedux = $ngRedux

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
})

.directive('app', AppComponent);
