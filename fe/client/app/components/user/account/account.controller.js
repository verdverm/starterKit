import * as AccountActions from '../../../actions/accounts'


class AccountController {
    constructor($ngRedux, $scope, $auth) {
        this.name = 'account';

        this.authenticate = $auth.authenticate;

        this.providers = [
            {
                name: "facebook",
                color: "#3b5998",
            },
            {
                name: "google",
                color: "#dd4b39",
            },
            {
                name: "yahoo",
                color: "#6e2a85",
            },
            {
                name: "github",
                color: "#444",
            },
            {
                name: "twitter",
                color: "#00aced",
            },
            {
                name: "soundcloud",
                color: "#fe3801",
            },
            {
                name: "spotify",
                color: "#7bb342",
            },
            {
                name: "dropbox",
                color: "#007ee5"
            }
        ];


        let unsubscribe = $ngRedux.connect(
            this.mapStateToThis,
            AccountActions
        )(this);

        $scope.$on('$destroy', unsubscribe);

    }

    mapStateToThis(state) {
        return {
            authState: state.auth,
            accounts: state.accounts,
        };
    }

    tryLinking(provider) {
        console.log("OAuth: ", provider);

        this.linkToProvider(provider);

    } 

    tryUnlinking(provider) {
        console.log("Unlink: ", provider);

        this.unlinkFromProvider(provider);
    } 

}

AccountController.$inject = ['$ngRedux', '$scope', '$auth'];

export default AccountController;
