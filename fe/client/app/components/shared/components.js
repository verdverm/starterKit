import angular from 'angular';

import Spinner from './spinner/spinner';

let componentModule = angular.module('shared.components', [
  Spinner.name
]);

export default componentModule;
