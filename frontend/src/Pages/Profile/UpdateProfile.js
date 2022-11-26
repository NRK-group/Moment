import config from '../../../config';

export default async function UpdateProfile(data, priv) {
    console.log({ priv }, data);
    let isPub = 0;
    priv === 'Public' ? (isPub = 1) : (isPub = 0);
    data.isPublic = isPub;
    return fetch(config.api + '/updateprofileinfo', {
        method: 'PUT',
        credentials: 'include',
        Accept: 'application/json',
        body: JSON.stringify(data),
    }).then((response) => {
        return response.json();
    });
}
