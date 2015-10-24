import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import loginComponent from './login.component';

let loginModule = angular.module('login', [
  'ui.router',
  'ngRedux',
  'satellizer'
])

.config(['$stateProvider', '$urlRouterProvider', ($stateProvider, $urlRouterProvider) => {
  $urlRouterProvider.otherwise('/');

  $stateProvider
    .state('login', {
      url: '/login',
      template: '<login></login>',
      data: {
        permissions: {
          only: ['anonymous'],
          redirectTo: 'profile'
        }
      }
    });
}])

.directive('login', loginComponent);

export default loginModule;
