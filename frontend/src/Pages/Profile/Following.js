import config from '../../../config';

export default async function GetFollowing() {
    return await fetch(config.api + '/getfollowing', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}
