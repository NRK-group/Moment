import { useEffect, useState } from 'react';
export const GetFollow = () => {
    const [following, setFollowing] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/message/new', {
            credentials: 'include',
        }).then(async (res) => {
            let data = await res.json();
            setFollowing(data);
        });
    }, []);
    return following;
};
