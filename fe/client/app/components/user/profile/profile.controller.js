class ProfileController {
    constructor($ngRedux, $scope) {
        this.name = 'profile';

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis

        )(this);

        $scope.$on('$destroy', unsubscribe);

    }

    mapStateToThis(state) {
        return {
            auth: state.auth,
            profile: state.profile
        };
    }

}

export default ProfileController;
