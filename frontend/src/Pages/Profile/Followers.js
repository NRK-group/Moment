import config from '../../../config';

export default async function GetFollowers() {
    return await fetch(config.api + '/followers', {
        credentials: 'include',
    }).then(async (response) => {
        return await response.json();
    });
}
