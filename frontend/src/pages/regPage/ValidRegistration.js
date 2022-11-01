import {ValidateEmail} from "../loginPage/ValidateLogin";

export default function SendRegistration(first, last, nick, about, email, password, confirmPassword, day) {
    console.log(first, last, nick, about, email, password, confirmPassword, {day})
    // ValidateRegistrationInfo(first, last, nick, about, email, password, confirmPassword, day, month, year)
}

function OnlyLetters(str) {
     return /^[A-Za-z-]*$/.test(str);
}

function ValidDate(day, month, year) {
    const DOB = new Date(year, month, day)
    console.log({DOB});
    if (DOB == 'Invalid Date') return [false, "Invalid Date of Birth"]
}

// isValid(new Date(14, "FEB", 2022))
// ValidDate(14, "FEB", 2022)


function ValidateRegistrationInfo(first, last, nick, about, email, password, confirmPassword, day, month, year) {
    
     const EMPTY = arguments.every((element, i) => {//Check if any values are empty (EXCEPT NICKNAME AND ABOUT ME)
        if (i == 2 || i== 3 ) return false
        if (element.trim().length() = 0) return true
    });
    if (EMPTY) return [false, "Required fields can't be empty"]
    if (password != confirmPassword) return [false, "Passwords don't match"]//Check the passwords match
    if (!ValidateEmail(email)) return [false, "Invalid email"]
    if (!OnlyLetters(first) || !OnlyLetters(last) || !OnlyLetters(nick)) return [false, "Names can only contain letters and hyphens"]

    

}