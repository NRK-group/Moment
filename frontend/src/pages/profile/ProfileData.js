export default async function GetProfile() {
    //Access the current cookie
    const COOKIE_VALUE = getCookie('session_token').split('&')[0];
    if (!COOKIE_VALUE) return 'Unauthorised';
    //Query the profile endpoint to get the data
    let result = fetch(
        'http://localhost:5070/profile?' +
            new URLSearchParams({
                userID: COOKIE_VALUE,
            }),
        {
            credentials: 'include',
        }
    ).then( response => {
        return response.json()
    })
    return await result
}

function getCookie(name) {
    // Split cookie string and get all individual name=value pairs in an array
    var cookieArr = document.cookie.split(';');

    // Loop through the array elements
    for (var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split('=');
        /* Removing whitespace at the beginning of the cookie name
        and compare it with the given string */
        if (name == cookiePair[0].trim()) {
            // Decode the cookie value and return
            return cookiePair[1];
        }
    }

    // Return null if not found
    return null;
}
