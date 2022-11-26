import { SearchResultsContainer } from './Components/SearchResultsContainer';
import './Search.css';
export const SearchModal = ({ query }) => {
    return (
        <div className='searchModalContainer'>
            <div className='searchModal'>
                <div className='searchContainer'>
                    <SearchResultsContainer query={query} />
                </div>
            </div>
        </div>
    );
};
