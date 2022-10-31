function CheckCreds(email, password) {
    //Check the values inside the login input fields
    if (password.length < 8 || password.length > 16 || !mixedCase(password)) return false //Check the password isn't too short
    //Check the email is valid
    if (!ValidateEmail(email)) return false
    return true
}  
function ValidateEmail(email) {
    var re = /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    return re.test(email);
  }

function mixedCase(str) {
    if (str === str.toUpperCase() || str === str.toLowerCase()) return false
    return true
}

export default function ValidateLoginAttempt(email, password, errMsg) {
    if (!CheckCreds(email, password)) return errMsg.innerHTML = 'Incorrect email or password' //return "Incorrect Email or Password" //Display the error message to client
    errMsg.innerHTML=''
    const LOGIN_CREDS = {//Make the obj with the login details
        Email:  email,
        Password: password
    }
    console.log(LOGIN_CREDS)
    //Send the data to the server to be validated by login handler
    // fetch('/login', {
    //     method: "POST",
    //     headers : {
    //         Accept: 'application/json',
    //             'Content-Type': 'application/json',
    //     }, 
    //     body: JSON.stringify(LOGIN_CREDS)
    // }).then(async response => {
    //     return await response.json()
    // }).then(resp => {
    //     console.log({resp})
    // })
}