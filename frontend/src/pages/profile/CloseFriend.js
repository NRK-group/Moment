export default async function GetCloseFriends() {
    return await fetch('http://localhost:5070/getclosefriend', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}
