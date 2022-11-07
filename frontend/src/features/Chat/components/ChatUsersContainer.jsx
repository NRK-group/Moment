import ChatContainerHeader from './ChatContainerHeader';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
export const ChatUsersContainer = ({
    currentUserName,
    styleName,
    setIsModalOpen,
    setcurrentReceiver,
}) => {
    let users = [];
    return (
        <div className={`chatUsersContainer ${styleName}`}>
            <ChatContainerHeader
                userName={currentUserName}
                setIsModalOpen={setIsModalOpen}
            />
            <div className='chatUsers scrollbar-hidden'>
                {users.map(({ img, id, name, content }) => (
                    <div
                        key={id}
                        onClick={() => {
                            setcurrentReceiver(name);
                        }}>
                        <MiniUserCard
                            key={id}
                            img={img}
                            propsId={`chat` + id}
                            name={name}>
                            {content}
                        </MiniUserCard>
                    </div>
                ))}
            </div>
        </div>
    );
};
