import { GetCookie } from './ProfileData';
//Checkfollowing checks the relationship between two accounts: Following, Not Following or Pending.
export default function CheckFollowing(userId) {
    return fetch(
        'http://localhost:5070/following?' +
            new URLSearchParams({
                //Add params to specifiy to server which user
                followingID: userId,
            }),
        {
            credentials: 'include',
        }
    ).then(async (response) => {
        return await response.json();
    });
}

//FollowRelationshipUpdate will send a request to either follow or unfollow a user.
function FollowRelationshipUpdate(followingId) {
    //Get the userId from the cookie
    let cookieVal = GetCookie('session_token');
    let userId;
    if (cookieVal) userId = cookieVal.split('&')[0];
    else return { Message: 'Error' };
    let FOLLOW_DATA = {
        FollowerId: userId,
        FollowingId: followingId,
    };
    return fetch('http://localhost:5070/followrequest', {
        method: 'PUT',
        credentials: 'include',
        Accept: 'application/json',
        body: JSON.stringify(FOLLOW_DATA),
    }).then((response) => {
        return response.json();
    });
}

//UpdateRelationshipBtn updates the style of the rel button
function UpdateRelationshipBtn(relationship, setter) {
    console.log({ relationship });
    switch (relationship) {
        case 'follow':
            setter('Following');
            break;
        case 'unfollow':
            setter('Follow');
            break;
        case 'pending':
            setter('Pending');
            break;
        default:
            setter('Error');
            break;
    }
}

//SetRelBtn sets the class and value of the relationship button when the page is first loaded
function SetRelBtn(rel, setter) {
    switch (rel) {
        case 'Not Following':
            setter('Follow');
            break;
        case 'Following':
            setter('Following');
            break;
        case 'Pending':
            setter('Pending');
            break;
        case 'Close Friend':
            setter('Close Friend');
            break;
        default:
            setter('Error');
            break;
    }
}

export { FollowRelationshipUpdate, UpdateRelationshipBtn, SetRelBtn };
