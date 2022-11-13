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
                {chatList.map(({ chatId, details, member, content }) => (
                    <NavLink
                        key={details.id}
                        to={`/messages/${chatId}`}
                        state={{
                            details: details,
                            user: member,
                        }}>
                        <MiniUserCard
                            img={details.img}
                            propsId={`chat` + details.id}
                            name={details.name}>
                            {content.content}
                        </MiniUserCard>
                    </NavLink>
                ))}
                <Outlet />
            </div>
        </div>
    );
};
