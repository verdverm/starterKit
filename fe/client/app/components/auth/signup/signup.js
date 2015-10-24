import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import signupComponent from './signup.component';

let signupModule = angular.module('signup', [
  'ui.router',
  'ngRedux',
])

.config(['$stateProvider', '$urlRouterProvider', ($stateProvider, $urlRouterProvider) => {
  $urlRouterProvider.otherwise('/');

  $stateProvider
    .state('signup', {
      url: '/signup',
      template: '<signup></signup>',
      data: {
        permissions: {
          only: ['anonymous'],
          redirectTo: 'profile'
        }
      }

    });
}])

.directive('signup', signupComponent);

export default signupModule;
