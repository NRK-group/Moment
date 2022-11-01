import {ValidateEmail} from "../loginPage/ValidateLogin";

export default async function SendRegistration(values, div) {
    let result = ValidateRegistrationInfo(values)
    if (!result[0]) {
        div.innerHTML = result[1]
        return
    }
    //Send registration request to the backend
    const REG_DETAILS = {
        FirstName : values[0],
        LastName: values[1],
        NickName: values[2],
        AboutMe: values[3],
        Email: values[4],
        Password:values[5],
        DateOfBirth: values[7],
    }
    console.log("FETCHING IN REG");

    let registered = await fetch('http://localhost:5070/registration', {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(REG_DETAILS),
    })
    .then(async response => {
        return await response.json()
    })
}

function OnlyLetters(str) {
     return /^[A-Za-z-]*$/.test(str);
}
function ValidateRegistrationInfo(args) {
    console.log({args})
    const FULL = args.every((element, i) => {//Check if any values are empty (EXCEPT NICKNAME AND ABOUT ME)
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