import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NavLink, Outlet } from 'react-router-dom';
export const ChatUsersContainer = ({
    currentUserInfo,
    styleName,
    chatList,
}) => {
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader userName={currentUserInfo} />
            <div className='chatUsers scrollbar-hidden'>
                {chatList.map(({ chatId, user, content }) => (
                    <NavLink
                        key={user.userId}
                        to={`/messages/${user.userId}`}
                        state={{
                            chatId: chatId,
                            img: user.img,
                            name: user.username,
                        }}>
                        <MiniUserCard
                            img={user.img}
                            propsId={`chat` + user.userId}
                            name={user.username}>
                            {content.content}
                        </MiniUserCard>
                    </NavLink>
                ))}
                <Outlet />
            </div>
        </div>
    );
};
