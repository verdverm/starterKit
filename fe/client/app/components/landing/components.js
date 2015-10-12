import angular from 'angular';

import Home from './home/home';
import About from './about/about';

let componentModule = angular.module('landing.components', [
  Home.name,
  About.name
]);

export default componentModule;
