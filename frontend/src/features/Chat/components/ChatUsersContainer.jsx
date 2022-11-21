import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NavLink, Outlet } from 'react-router-dom';
import { useEffect, useState } from 'react';
export const ChatUsersContainer = ({
    currentUserInfo,
    styleName,
    chatList,
    socket,
}) => {
    const [newMessage, setNewMessage] = useState([]);
    useEffect(() => {
        setNewMessage(chatList);
    }, [chatList]);
    const handleOpenChat = (chatId) => {
        socket.send(
            JSON.stringify({
                type: 'deleteNotif',
                chatId: chatId,
                receiverId: currentUserInfo,
            })
        );
        setNewMessage(
            newMessage.map((msg) => {
                if (msg.chatId === chatId) {
                    msg.notif = 0;
                }
                return msg;
            })
        );
    };
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader userName={currentUserInfo} />
            <div className='chatUsers scrollbar-hidden'>
                {Array.isArray(newMessage)
                    ? newMessage.map(
                          ({
                              type,
                              chatId,
                              details,
                              member,
                              content,
                              notif,
                          }) => {
                              return (
                                  <div
                                      onClick={() => {
                                          handleOpenChat(chatId);
                                      }}>
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
                                  </div>
                              );
                          }
                      )
                    : null}
                <Outlet />
            </div>
        </div>
    );
};
