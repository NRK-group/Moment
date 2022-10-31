function CheckCreds(email, password) {
    //Check the values inside the login input fields
    if (password.length < 8 || password.length > 16 || !mixedCase(password))
        return false; //Check the password isn't too short
    //Check the email is valid
    if (!ValidateEmail(email)) return false;
    return true;
}
function ValidateEmail(email) {
    var re =
        /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    return re.test(email);
}

function mixedCase(str) {
    if (str === str.toUpperCase() || str === str.toLowerCase()) return false;
    return true;
}

export default function ValidateLoginAttempt(email, password, errMsg) {
    if (!CheckCreds(email, password)) {
        errMsg.innerHTML = 'Incorrect email or password';
        return false//Display the error message to client
    }
    const LOGIN_CREDS = {
        //Make the obj with the login details
        Email: email,
        Password: password,
    };
    let auth = true;
    // Send the data to the server to be validated by login handler
    fetch('http://localhost:5070/login', {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(LOGIN_CREDS),
    })
        .then(async (response) => {
            return await response.text();
        })
        .then((resp) => {
            if (resp !== 'Valid Login') {
                errMsg.innerHTML = resp;
                console.log("INSIDE");
                auth = false;
            }
        });
    return auth;
}
