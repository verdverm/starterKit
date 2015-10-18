import angular from 'angular';

import Shared from './shared/components';
import Landing from './landing/components';
import Auth from './auth/components';
import User from './user/components';


let componentModule = angular.module('app.components', [
  Shared.name,
  Landing.name,
  Auth.name,
  User.name
]);

export default componentModule;
