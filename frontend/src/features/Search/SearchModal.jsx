import { useEffect, useState } from 'react';
import { SearchResultsContainer } from './components/SearchResultsContainer';
import './Search.css';
export const SearchModal = ({ query }) => {
    const [searchResult, setSearchResult] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/search', { credentials: 'include' }).then(
            async (resp) => {
                const data = await resp.json();
                setSearchResult(data);
            }
        );
    }, []);
    return (
        <div className='searchModalContainer'>
            <div className='searchModal'>
                <div className='searchContainer'>
                    <SearchResultsContainer
                        searchResult={searchResult}
                        query={query}
                    />
                </div>
            </div>
        </div>
    );
};
