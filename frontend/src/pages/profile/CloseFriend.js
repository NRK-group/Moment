import { GetCookie } from './ProfileData';

export default async function GetCloseFriends() {
    return await fetch('http://localhost:5070/getclosefriend', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}

async function UpdateCloseFriends(closeFriend, set) {
    const USER_ID = GetCookie('session_token').split('&')[0];
    const DATA = {
        UserID: USER_ID,
        CloseFriendId: closeFriend,
    };
    return await fetch('http://localhost:5070/closefriend', {
        credentials: 'include',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        method: "POST",
        body: JSON.stringify(DATA),
    }).then(async (response) => {
        return await response.json();
    }).then(resp => {
    if (resp.Message === "Removed") set("Add")
    if (resp.Message === "Added") set("Remove")

     console.log("UPDATED TO:   ", resp.Message)
    });
}

export { UpdateCloseFriends };
