export default async function GetFollowing() {
    return await fetch('http://localhost:5070/following', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}