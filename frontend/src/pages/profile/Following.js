export default async function GetFollowing() {
    return await fetch('http://localhost:5070/getfollowing', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}