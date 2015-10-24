import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import aboutComponent from './about.component';

let aboutModule = angular.module('about', [
  'ui.router',
  'ngRedux',
])

.config(['$stateProvider', ($stateProvider) => {
  $stateProvider
    .state('about', {
      url: '/about',
      template: '<about></about>'
    });
}])

.directive('about', aboutComponent);

export default aboutModule;
