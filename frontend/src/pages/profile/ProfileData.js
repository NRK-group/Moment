export default async function GetProfile(user) {
    //Query the profile endpoint to get the data
    if (!user) user = GetCookie("session_token").split("&")[0]
    return await fetch(
        'http://localhost:5070/profile?' +
            new URLSearchParams({
                userID: user,
            }),
        {
            credentials: 'include',
        }
    ).then((response) => {
        return response.json();
    });
}

function GetCookie(name) {
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

// export { GetCookie, GetProfile };
