import config from '../../../config';
import { GetCookie } from './ProfileData';

export default function GetPosts(id) {
    if (!id) id = GetCookie('session_token').split('&')[0];
    return fetch(
        //Query the profile endpoint to get the data
        config.api +
            '/getUserPosts?' +
            new URLSearchParams({
                //Add params to specifiy to server which user
                userID: id,
            }),
        {
            credentials: 'include',
        }
    ).then((response) => {
        return response.json();
    });
}
