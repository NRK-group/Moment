import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
export const ChatUsersContainer = ({
    currentUserName,
    styleName,
    setIsModalOpen,
    setcurrentReceiver,
    chatList,
    setReceiverInfo,
}) => {
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader
                userName={currentUserName}
                setIsModalOpen={setIsModalOpen}
            />
            <div className='chatUsers scrollbar-hidden'>
                {chatList.map(({ chatId, user, content }) => (
                    <div
                        key={chatId}
                        onClick={() => {
                            setcurrentReceiver(user.username);
                            setReceiverInfo({
                                chatId: chatId,
                                userId: user.userId,
                                username: user.username,
                                img: user.img,
                            });
                        }}>
                        <MiniUserCard
                            key={user.userId}
                            img={user.img}
                            propsId={`chat` + user.userId}
                            name={user.username}>
                            {content.content}
                        </MiniUserCard>
                    </div>
                ))}
            </div>
        </div>
    );
};
