import {loginUser} from '../../../actions/auth'
import routerActions from '../../../actions/router'

class LoginController {
    constructor($ngRedux, $scope, $auth) {
        this.name = 'login';

        this.authenticate = $auth.authenticate;

        this.creds = {
            username: {
                name: 'username',
                value: '',
                error: ''
            },
            password: {
                name: 'password',
                value: '',
                error: ''
            },

        }

        let localActions = Object.assign({}, routerActions, {loginUser:loginUser})

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            localActions

        )(this);

        $scope.$on('$destroy', unsubscribe);

    }

    mapStateToThis(state) {
        return {
            authState: state.auth
        };
    }

    clearErrors() {
        this.creds.username.error = '';
        this.creds.password.error = '';
    }

    tryLogin() {
        this.clearErrors();

        // clientSide validation...
        function checkEmpty(field) {
            if ( field.value === "" ) {
                field.error = "Field empty but required";
                return false;
            }
            return true;
        }
        function checkMatch(field1, field2) {
            if ( field1.value === field2.value ) {
                field2.error = "Field must match " + field1.name;
                return false;
            }
            return true;
        }

        let passed = true;
        passed = checkEmpty(this.creds.username) && passed;
        passed = checkEmpty(this.creds.password) && passed;

        if (passed) {
            var username = this.creds.username.value;
            var password = this.creds.password.value;

            this.loginUser(username, password);
        }

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
