import template from './account.html';
import controller from './account.controller';
import './account.styl';

let accountComponent = function () {
  return {
    restrict: 'E',
    scope: {},
    template,
    controller,
    controllerAs: 'vm',
    bindToController: true
  };
};

export default accountComponent;
