class AppController {
    constructor($ngRedux, $scope) {
        this.name = 'app';

        this.globalState = {};

        let unsubscribe = $ngRedux.connect(this.mapStateToThis)(this);
        $scope.$on('$destroy', unsubscribe);

    }

    mapStateToThis(state) {
        return {
            globalState: state,
            auth: state.auth
        };
    }
}

export default AppController;
