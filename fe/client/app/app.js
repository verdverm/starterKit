import angular from 'angular';
import ngTouch from 'angular-touch';
import Permission from 'angular-permission'

import ngRedux from 'ng-redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';
import { devTools } from 'redux-devtools';

import uiRouter from 'angular-ui-router';
import ngReduxRouter from 'redux-ui-router';

import PDB from './pdb'

import rootReducer from './reducers';

import Common from './common/common';
import Components from './components/components';
import AppComponent from './app.component';

import { loadedUserAuth, loadedUserAuthBegin, loadedUserAuthFail } from './actions/auth';
import routerActions from "./actions/router"


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
    console.log("Config'n App");

    const logger = createLogger({
        level: 'info',
        collapsed: true
    });

    $ngReduxProvider.createStoreWith(
        rootReducer, 
        ['ngUiRouterMiddleware', thunk, logger], 
        [devTools()]
    );
})

.run(function(Permission, $ngRedux, $state) {

    // grab local copy of redux and routerState for closures
    let localRedux = $ngRedux;
    let localState = $state;

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

    console.log("Running");


    // dispatch that we are starting to load auth
    localRedux.dispatch(loadedUserAuthBegin());

    // start the auth load
    PDB.get('auth_data').then(function (doc) {
        console.log("PouchDB loaded:", doc);
        
        localRedux.dispatch(loadedUserAuth(
            doc.username,
            doc.uid,
            doc.token,
            doc._rev
        ));

        console.log("PouchDB dispatched");

        // determine where we should redirect after successful auth load
        var self_name = localState.$current.self.name;
        if (self_name === "login" || self_name === "register") {
            localRedux.dispatch(routerActions.stateGo('profile'));
        }

    }, function (error) {
        // determine where we should redirect after successful auth load
        // var self_name = localState.$current.self.name;
        // if (self_name === "login" || self_name === "register") {
        //     localRedux.dispatch(routerActions.stateGo('profile'));
        // }
        // THIS SHOULD REDIRECT PROPERLY
        localRedux.dispatch(loadedUserAuthFail(error));
        alert("Error loading auth: " + error);
    });

    console.log("Leaving RUN");
    
})

.directive('app', AppComponent);
