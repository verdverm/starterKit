import * as AccountActions from '../../../actions/accounts';
import OauthProviders from '../../shared/consts/accounts';

class AccountController {
    constructor($ngRedux, $scope, $auth) {
        this.name = 'account';

        this.authenticate = $auth.authenticate;

        this.providers = OauthProviders;


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
