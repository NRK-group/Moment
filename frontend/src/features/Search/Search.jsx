import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import Input from '../../components/Input/Input';
import { SearchResultsContainer } from './components/SearchResultsContainer';
import './Search.css';
export const Search = () => {
    let searchResult = [];
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='searchPageContainer'>
                    <Input
                        styleName='searchPage'
                        type='search'
                        placeholder='Search'
                    />
                </div>
                <SearchResultsContainer searchResult={searchResult} />
            </Card>
        </Body>
    );
};
