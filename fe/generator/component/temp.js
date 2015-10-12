import angular from 'angular';
import uiRouter from 'angular-ui-router';
import ngRedux from 'ng-redux';

import <%= name %>Component from './<%= name %>.component';

let <%= name %>Module = angular.module('<%= name %>', [
  uiRouter,
  ngRedux
])

.config(($stateProvider) => {
  $stateProvider
    .state('<%= name %>', {
      url: '/<%= name %>',
      template: '<<%= name %>></<%= name %>>',
      data: {
        permissions: {
          except: ['anonymous'],
          redirectTo: 'login'
        }
      }
    });
})

.directive('<%= name %>', <%= name %>Component);

export default <%= name %>Module;
