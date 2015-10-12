import angular from 'angular';

import Profile from './profile/profile';

let componentModule = angular.module('user.components', [
  Profile.name
]);

export default componentModule;
