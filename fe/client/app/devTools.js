import { persistState } from 'redux-devtools';
import { DevTools, DebugPanel, LogMonitor } from 'redux-devtools/lib/react';
import React, { Component } from 'react';
import ReactDOM from 'react-dom';

angular.module('app')
  .run(($ngRedux, $rootScope) => {
    ReactDOM.render(
      <DevApp store={ $ngRedux }/>,
      document.getElementById('devTools')
    );
    //Hack to reflect state changes when disabling/enabling actions via the monitor
    $ngRedux.subscribe(_ => {
        setTimeout($rootScope.$apply.bind($rootScope), 100);
    });
  });


class DevApp extends Component {
  render() {
    return (
      <div>
        <DebugPanel top left bottom>
          <DevTools store={ this.props.store } monitor = { LogMonitor } />
        </DebugPanel>
      </div>
    );
  }
}


