import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
export const ChatUsersContainer = ({ users }) => {
    return (
        <div className='chatUsersContainer'>
            <ChatContainerHeader userName='Moment' />
            <div className='chatUsers scrollbar-hidden'>
                {users.map(({ img, id, name, content }) => (
                    <MiniUserCard
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
