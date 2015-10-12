import template from './signup.html';
import controller from './signup.controller';
import './signup.styl';

let signupComponent = function () {
  return {
    restrict: 'E',
    scope: {},
    template,
    controller,
    controllerAs: 'vm',
    bindToController: true
  };
};

export default signupComponent;
