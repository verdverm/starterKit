import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import spinnerComponent from './spinner.component';

let spinnerModule = angular.module('spinner', [])

.directive('spinner', spinnerComponent);

export default spinnerModule;
