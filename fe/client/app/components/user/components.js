import angular from 'angular';

import Profile from './profile/profile';
import Account from './account/account';

let componentModule = angular.module('user.components', [
  Profile.name,
  Account.name
]);

export default componentModule;
