import { NavLink } from 'react-router-dom';
import MiniUserCard from '../../../Components/MiniUserCard/MiniUserCard';
import { GetAllUser } from '../Hooks/GetAllUser';
import { NoSearchResult } from './NoSearchResult';
export const SearchResultsContainer = ({ query }) => {
    let searchResult = GetAllUser();
    let filteredItems = searchResult.filter((item) => {
        let fullName = item.firstName + ' ' + item.lastName;
        return (
            item.name.toLowerCase().includes(query.toLowerCase()) ||
            fullName.toLowerCase().includes(query.toLowerCase())
        );
    });
    return (
        <>
            {filteredItems.length != 0 ? (
                <div className='searchResultsContent scrollbar-hidden'>
                    {filteredItems.map(
                        ({ name, id, firstName, lastName, img }) => {
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
