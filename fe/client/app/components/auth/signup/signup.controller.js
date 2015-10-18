import { signupUser } from '../../../actions/auth'

class SignupController {
    constructor($ngRedux, $scope, $auth) {
        this.name = 'signup';

        this.refresh = function() {
        	$scope.$apply();
        }

        this.authenticate = $auth.authenticate;

        this.creds = {
            email: {
            	name: 'email',
            	value: '',
            	error: ''
            },
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
            confirm: {
            	name: 'confirm',
            	value: '',
            	error: ''
            }
        }

        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            {signupUser}

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
    	this.creds.email.error = '';
    	this.creds.password.error = '';
    	this.creds.confirm.error = '';
    }

    signup() {
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
    	passed = checkEmpty(this.creds.email) && passed;
    	passed = checkEmpty(this.creds.password) && passed;
    	passed = checkEmpty(this.creds.confirm) && passed;
    	passed = checkMatch(this.creds.password, this.creds.confirm) && passed;

        if (passed) {
	        var username = this.creds.username.value;
	        var email    = this.creds.email.value;
	        var password = this.creds.password.value;
	        var confirm  = this.creds.confirm.value;

	        this.signupUser(username, email, password);
        }

    }

    tryAuthenticate(provider) {
        console.log("OAuth: ", provider);
        this.authenticate(provider)
        .then(function() {
            var res_str = 'You have successfully signed up with ' + provider;
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

export default SignupController;
