function checkCreds() {
    //Check the values inside the login input fields
    const emailInput = document.querySelector('.loginEmailInput'), passwordInput = document.querySelector('.loginPasswordInput')
    if (passwordInput.value.length < 8 || passwordInput.value.length > 16 || !mixedCase(passwordInput.value)) return [false, "Password must be between 8-16 characters"]//Check the password isn't too short
    //Check the email is valid
    if (!ValidateEmail(emailInput.value)) return [false, "invalid email"]

}  
function ValidateEmail(email) {
    var re = /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    return re.test(email);
  }

function mixedCase(str) {
    if (str === str.toUpperCase() || str === str.toLowerCase()) return false
    return true
}

