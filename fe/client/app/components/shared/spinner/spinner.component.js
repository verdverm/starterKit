import template from './spinner.html';
// import './spinner.styl';

let spinnerComponent = function () {
  return {
    restrict: 'E',
    scope: {},
    template,
  };
};

export default spinnerComponent;
