import config from '../../../config';

function CheckCreds(email, password) {
    //Check the values inside the login input fields
    //Check the password isn't too short
    if (password.length < 8 || password.length > 16 || !mixedCase(password)) {
        return false;
    }
    //Check the email is valid
    if (!ValidateEmail(email)) return false;
    return true;
}
export function ValidateEmail(email) {
    var re =
        /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    return re.test(email);
}

function mixedCase(str) {
    if (str === str.toUpperCase() || str === str.toLowerCase()) return false;
    return true;
}

export default function ValidateLoginAttempt(email, password, errMsg) {
    errMsg.innerHTML = `<div class="dot-flashing"></div>`;
    if (!CheckCreds(email, password)) {
        errMsg.innerHTML = 'Incorrect email or password';
        return false; //Display the error message to client
    }
    const LOGIN_CREDS = {
        //Make the obj with the login details
        Email: email.toLowerCase(),
        Password: password,
    };
    // Send the data to the server to be validated by login handler
    let auth = fetch(config.api + '/login', {
        credentials: 'include',
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(LOGIN_CREDS),
    })
        .then((response) => {
            return response.json();
        })
        .then((resp) => {
            errMsg.innerHTML = '';

            if (resp.Message !== 'Valid Login') {
                errMsg.innerHTML = resp.Message;
                return false;
            }
            if (resp.Message === 'Valid Login') return true;
        });
    return auth;
}
