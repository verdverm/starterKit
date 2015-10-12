import angular from 'angular';

import Landing from './landing/components';
import Auth from './auth/components';
import User from './user/components';


let componentModule = angular.module('app.components', [
  Landing.name,
  Auth.name,
  User.name
]);

export default componentModule;
