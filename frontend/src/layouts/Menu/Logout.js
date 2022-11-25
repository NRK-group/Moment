import config from '../../../config';

export default function Logout(redir, setauth) {
    fetch(config.api + '/logout', {
        credentials: 'include',
    })
        .then(async (response) => {
            return await response.text();
        })
        .then((resp) => {
            if (resp === 'Logged Out') {
                setauth(false);
                redir('/');
                return true;
            }
            return false;
        });
}
