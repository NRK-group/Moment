import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoNotifications } from './NoNotifications';
export const NotificationContentContainer = ({ notif }) => {
    console.log(notif);
    return (
        <div className='notificationContentContainer'>
            {notif && notif.length !== 0 ? (
                notif.map(({ img, id, name, content, optContent }) => (
                    <MiniUserCard
                        key={id}
                        img={img}
                        propsId={`chat` + id}
                        name={name}
                        content={content}
                        optContent={optContent}
                    />
                ))
            ) : (
                <NoNotifications />
            )}
        </div>
    );
};
