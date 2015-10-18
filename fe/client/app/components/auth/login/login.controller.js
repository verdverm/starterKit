import {loginUser} from '../../../actions/auth'
import routerActions from '../../../actions/router'

class LoginController {
    constructor($ngRedux, $scope, $auth) {
        this.name = 'login';

        this.authenticate = $auth.authenticate;

        this.creds = {
            username: '',
            password: ''
        }

        let localActions = Object.assign({}, routerActions, {loginUser:loginUser})
        // console.log(routerActions);
        // console.log(localActions);

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            localActions

        )(this);

        $scope.$on('$destroy', unsubscribe);

        // console.log(this);

    }

    mapStateToThis(state) {
        return {
            authState: state.auth
        };
    }

    tryLogin() {
        // clientSide validation...

        this.loginUser(this.creds.username, this.creds.password);
    }

    tryAuthenticate(provider) {
        console.log("OAuth: ", provider);
        this.authenticate(provider)
        .then(function() {
            var res_str = 'You have successfully signed in with ' + provider;
            Materialize.toast(res_str, 4000)
          // $location.path('/');
        })
        .catch(function(response) {
            var err_str = "Error: " + response.data.message;
            Materialize.toast(err_str, 4000)
          // toastr.error(response.data.message);
        });
    } 


}

export default LoginController;
