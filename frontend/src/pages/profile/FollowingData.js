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
