const emailInput, passwordInput
function CheckCreds() {
    //Check the values inside the login input fields
    emailInput = document.querySelector('.loginEmailInput'), passwordInput = document.querySelector('.loginPasswordInput')
    if (passwordInput.value.length < 8 || passwordInput.value.length > 16 || !mixedCase(passwordInput.value)) return [false, "Incorrect Password"]//Check the password isn't too short
    //Check the email is valid
    if (!ValidateEmail(emailInput.value)) return [false, "Account not found"]
    return [true, "Valid entry", emailInput, passwordInput]
}  
function ValidateEmail(email) {
    var re = /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    return re.test(email);
  }

function mixedCase(str) {
    if (str === str.toUpperCase() || str === str.toLowerCase()) return false
    return true
}

function ValidateLoginAttempt() {
    if (!CheckCreds()[0]) { //Display the error message to client
        return
    }
    
    const LOGIN_CREDS = {//Make the obj with the login details
        Email:  emailInput.value,
        Password: passwordInput.value
    }
    //Send the data to the server to be validated by login handler
    fetch('/login')
}