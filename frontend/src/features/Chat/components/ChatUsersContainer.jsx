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
                {Array.isArray(chatList)
                    ? chatList.map(
                          ({
                              type,
                              chatId,
                              details,
                              member,
                              content,
                              notif,
                          }) => {
                              return (
                                  <NavLink
                                      key={details.id}
                                      to={`/messages/${chatId}`}
                                      state={{
                                          type: type,
                                          details: details,
                                          user: member,
                                      }}>
                                      <MiniUserCard
                                          img={details.img}
                                          propsId={`chat` + details.id}
                                          name={details.name}
                                          button={
                                              <>
                                                  {notif && notif != 0 ? (
                                                      <div className='messageNotifIndicator'>
                                                          {notif}
                                                      </div>
                                                  ) : null}
                                              </>
                                          }>
                                          {content.content}
                                      </MiniUserCard>
                                  </NavLink>
                              );
                          }
                      )
                    : null}
                <Outlet />
            </div>
        </div>
    );
};
