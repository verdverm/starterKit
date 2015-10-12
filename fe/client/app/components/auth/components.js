import angular from 'angular';

import Signup from './signup/signup';
import Login from './login/login';

let componentModule = angular.module('auth.components', [
  Signup.name,
  Login.name
]);

export default componentModule;
