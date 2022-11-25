import { useEffect, useState } from 'react';
import config from '../../../../config';
export const GetFollow = () => {
    const [following, setFollowing] = useState([]);
    useEffect(() => {
        fetch(config.api + '/message/new', {
            credentials: 'include',
        }).then(async (res) => {
            let data = await res.json();
            data ? setFollowing(data) : setFollowing([]);
        });
    }, []);
    return following;
};
