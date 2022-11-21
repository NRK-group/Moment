import { useEffect, useState } from 'react';

export const GetAllUser = () => {
    const [searchResult, setSearchResult] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/search', { credentials: 'include' }).then(
            async (resp) => {
                const data = await resp.json();
                setSearchResult(data);
            }
        );
    }, []);
    return searchResult;
};
