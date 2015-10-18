class AccountController {
    constructor($ngRedux, $scope) {
        this.name = 'account';

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

export default AccountController;
