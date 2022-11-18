import { NavLink } from 'react-router-dom';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoSearchResult } from './NoSearchResult';
export const SearchResultsContainer = ({ searchResult, query }) => {
    let filteredItems = searchResult.filter((item) => {
        return (
            item.name.toLowerCase().includes(query.toLowerCase()) ||
            item.firstName.toLowerCase().includes(query.toLowerCase()) ||
            item.lastName.toLowerCase().includes(query.toLowerCase())
        );
    });
    return (
        <>
            {filteredItems ? (
                <div className='searchResultsContent scrollbar-hidden'>
                    {filteredItems.map(
                        ({ name, id, firstName, lastName, img }) => {
                            console.log(searchResult);
                            return (
                                <NavLink key={id} to={`/profile?id=${id}`}>
                                    <MiniUserCard
                                        propsId={id}
                                        name={firstName + ' ' + lastName}
                                        img={img}
                                        button={<></>}>
                                        {name}
                                    </MiniUserCard>
                                </NavLink>
                            );
                        }
                    )}
                </div>
            ) : (
                <NoSearchResult />
            )}
        </>
    );
};
