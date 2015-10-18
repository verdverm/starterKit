import {logoutUser} from './actions/auth'

class AppController {
    constructor($ngRedux, $scope) {
        this.name = 'app';

        this.globalState = {};

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            {logoutUser}
        )(this);
        $scope.$on('$destroy', unsubscribe);

    }

    logout() {
        this.logoutUser(this.auth.rev);
    }

    mapStateToThis(state) {
        return {
            globalState: state,
            auth: state.auth   // for the menu and titles
        };
    }
}

export default AppController;
