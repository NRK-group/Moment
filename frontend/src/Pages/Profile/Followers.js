export default async function GetFollowers() {
    return await fetch('http://localhost:5070/followers', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}