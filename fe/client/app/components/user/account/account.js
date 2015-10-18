import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import accountComponent from './account.component';

let accountModule = angular.module('account', [
  uiRouter,
  ngRedux
])

.config(($stateProvider) => {
  $stateProvider
    .state('account', {
      url: '/account',
      template: '<account></account>',
      data: {
        permissions: {
          except: ['anonymous'],
          redirectTo: 'login'
        }
      }
    });
})

.directive('account', accountComponent);

export default accountModule;
