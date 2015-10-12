class <%= upCaseName %>Controller {
    constructor($ngRedux, $scope) {
        this.name = '<%= name %>';

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis

        )(this);

        $scope.$on('$destroy', unsubscribe);

    }

    mapStateToThis(state) {
        return {
            // someState: state.something
        };
    }

}

export default <%= upCaseName %>Controller;
