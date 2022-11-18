import { SearchResultsContainer } from './components/SearchResultsContainer';
import './Search.css';
export const SearchModal = () => {
    let searchResult = [];
    return (
        <div className='searchModalContainer'>
            <div className='searchModal'>
                <div className='searchContainer'>
                    <SearchResultsContainer searchResult={searchResult} />
                </div>
            </div>
        </div>
    );
};
