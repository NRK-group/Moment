import {ValidateEmail} from "../loginPage/ValidateLogin";

export default function SendRegistration(values, div) {
    // ValidateRegistrationInfo(first, last, nick, about, email, password, confirmPassword, day, month, year)
    let result = ValidateRegistrationInfo(values)
    console.log(result)
    if (!result[0]) {
        div.innerHTML = result[1]
        return
    }
    console.log("PASSED")
}

function OnlyLetters(str) {
     return /^[A-Za-z-]*$/.test(str);
}
function ValidateRegistrationInfo(args) {
    console.log({args})
    const FULL = args.every((element, i) => {//Check if any values are empty (EXCEPT NICKNAME AND ABOUT ME)
        console.log({element} , i);
        if  (i == (args.length - 1)) return true
        if (element.trim().length === 0) return false
        return true
    });
    if (!FULL) return [false, "Required fields can't be empty"]//Check no required fields are empty
    if (args[5] != args[6]) return [false, "Passwords don't match"]//Check the passwords match
    if (!ValidateEmail(args[4])) return [false, "Please enter a valid email"]// Cehck the email is valid
    if (!OnlyLetters(args[0]) || !OnlyLetters(args[1]) || !OnlyLetters(args[2])) return [false, "Names can only contain letters and hyphens"]// Check names only contain hyphens and letters 
    if (!args[7]) return [false, "Please enter a valid date"]
    return[true, ""]
}