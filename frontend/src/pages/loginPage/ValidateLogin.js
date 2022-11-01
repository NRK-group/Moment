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
    console.log('CHECKING', CheckCreds(email, password));
    if (!CheckCreds(email, password)) {
        console.log('WRONG DETAILS', CheckCreds(email, password));
        errMsg.innerHTML = 'Incorrect email or password';
        return false; //Display the error message to client
    }
    const LOGIN_CREDS = {
        //Make the obj with the login details
        Email: email,
        Password: password,
    };
    // Send the data to the server to be validated by login handler
    console.log('FETCHING IN LOGIN');
    let auth = fetch('http://localhost:5070/login', {
        crossDomain: true,
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(LOGIN_CREDS),
    })
        .then((response) => {
            return response.text();
        })
        .then((resp) => {
        console.log("THIRD RESPONSE")
            if (resp !== 'Valid Login') {
                errMsg.innerHTML = resp;
                return false;
            }
            if (resp === 'Valid Login') return true;
        });
        console.log({auth})
    return auth;
}
