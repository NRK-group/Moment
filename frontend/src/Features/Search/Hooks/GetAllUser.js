import { useEffect, useState } from 'react';
import config from '../../../../config';

export const GetAllUser = () => {
    const [searchResult, setSearchResult] = useState([]);
    useEffect(() => {
        fetch(config.api + '/search', { credentials: 'include' }).then(
            async (resp) => {
                const data = await resp.json();
                data ? setSearchResult(data) : setSearchResult([]);
            }
        );
    }, []);
    return searchResult;
};
