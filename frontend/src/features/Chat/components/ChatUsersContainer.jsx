import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
export const ChatUsersContainer = ({ users, currentUserName, styleName }) => {
    console.log(styleName);
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader userName={currentUserName} />
            <div className='chatUsers scrollbar-hidden'>
                {users.map(({ img, id, name, content }) => (
                    <MiniUserCard
                        key={id}
                        img={img}
                        propsId={`chat` + id}
                        name={name}>
                        {content}
                    </MiniUserCard>
                ))}
            </div>
        </div>
    );
};
