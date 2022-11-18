import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoSearchResult } from './NoSearchResult';
export const SearchResultsContainer = ({ searchResult }) => {
    return (
        <>
            {searchResult && searchResult.length !== 0 ? (
                <div className='searchResultsContent scrollbar-hidden'>
                    {searchResult.map(({ name, id, content, img }) => (
                        <MiniUserCard
                            key={id}
                            name={name}
                            propsId={id}
                            content={content}
                            img={img}
                        />
                    ))}
                </div>
            ) : (
                <NoSearchResult />
            )}
        </>
    );
};
