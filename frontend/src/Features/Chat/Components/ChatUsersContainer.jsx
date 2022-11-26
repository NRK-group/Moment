import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../Components/MiniUserCard/MiniUserCard';
import { NavLink, Outlet } from 'react-router-dom';
import { useEffect, useState } from 'react';
export const ChatUsersContainer = ({
    currentUserInfo,
    styleName,
    chatList,
    socket,
    setNewMessageNotif,
    style,
    username,
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
        let isNotif = false;
        for (let i = 0; i < newMessage.length; i++) {
            if (newMessage[i].notif !== 0) {
                isNotif = true;
                return;
            }
        }
        setNewMessageNotif(isNotif);
    };
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader userName={username} style={style} />
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
                                      key={details.id}
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
                                              name={
                                                  style != 'mobile'
                                                      ? details.name
                                                      : details.name.slice(0, 1)
                                              }
                                              button={
                                                  <>
                                                      {notif != 0 ? (
                                                          <div className='messageNotifIndicator'>
                                                              {notif}
                                                          </div>
                                                      ) : null}
                                                  </>
                                              }>
                                              {style != 'mobile'
                                                  ? content.content
                                                  : null}
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
