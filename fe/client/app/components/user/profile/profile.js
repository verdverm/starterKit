import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import profileComponent from './profile.component';

let profileModule = angular.module('profile', [
  uiRouter,
  ngRedux
])

.config(($stateProvider) => {
  $stateProvider
    .state('profile', {
      url: '/profile',
      template: '<profile></profile>',
      data: {
        permissions: {
          except: ['anonymous'],
          redirectTo: 'login'
        }
      }
    });
})

.directive('profile', profileComponent);

export default profileModule;
