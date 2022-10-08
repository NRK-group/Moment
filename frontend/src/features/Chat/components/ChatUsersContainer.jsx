import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
export const ChatUsersContainer = ({ users, currentUserName }) => {
    return (
        <div className='chatUsersContainer'>
            <ChatContainerHeader userName={currentUserName} />
            <div className='chatUsers scrollbar-hidden'>
                {users.map(({ img, id, name, content }) => (
                    <MiniUserCard
                        key={id}
                        img={img}
                        propsId={id}
                        name={name}
                        content={content}
                    />
                ))}
            </div>
        </div>
    );
};
