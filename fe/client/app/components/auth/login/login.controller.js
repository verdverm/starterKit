import {loginUser} from '../../../actions/auth'
import routerActions from '../../../actions/router'

class LoginController {
    constructor($ngRedux, $scope) {
        this.name = 'login';

        this.creds = {
            username: '',
            password: ''
        }

        let localActions = Object.assign({}, routerActions, {loginUser:loginUser})
        console.log(routerActions);
        console.log(localActions);

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            localActions

        )(this);

        $scope.$on('$destroy', unsubscribe);

        console.log(this);

    }

    mapStateToThis(state) {
        return {
            authState: state.auth
        };
    }

    tryLogin() {
        console.log("LoginController:", this.creds.username, this.creds.password);
        console.log("authState:", this.authState);

        // clientSide validation...

        this.loginUser(this.creds.username, this.creds.password);
    }


}

export default LoginController;
