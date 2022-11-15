import { GetCookie } from "./ProfileData";
//Checkfollowing checks the relationship between two accounts: Following, Not Following or Pending.
export default function CheckFollowing(userId) {
    return fetch(
        'http://localhost:5070/following?' +
            new URLSearchParams(
                {
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
function FollowRelationshipUpdate (followingId) {
    //Get the userId from the cookie
    let userId = GetCookie("session_token").split["&"][0]
    let FOLLOW_DATA = {
        FollowerId: userId,
        FollowingId: followingId,
    }
    fetch(
        'http://localhost:5070/following', {
            method: "PUT",
            credentials: "include",
            Accept: 'application/json',
            body: JSON.stringify(FOLLOW_DATA)
        }).then(response => {
            return response.json()
        }).then(resp=> {
            console.log(resp);
        })

}