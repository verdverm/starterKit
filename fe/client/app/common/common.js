import angular from 'angular';
import User from './user/user';

let commonModule = angular.module('app.common', [
  User.name
]);

export default commonModule;
