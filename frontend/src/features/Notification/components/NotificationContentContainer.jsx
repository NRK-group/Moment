import { useState } from 'react';
import { useParams } from 'react-router-dom';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoNotifications } from './NoNotifications';
export const NotificationContentContainer = () => {
    const [notif, setNotif] = useState([]);
    let generalNotif = [];
    let followNotif = [];
    let groupNotif = [];
    const { type } = useParams();
    console.log(type);
    return (
        <div className='notificationContentContainer'>
            {notif && notif.length !== 0 ? (
                notif.map(({ img, id, name, content, optContent }) => (
                    <MiniUserCard
                        key={id}
                        img={img}
                        propsId={`notif` + id}
                        name={name}
                        optContent={optContent}>
                        {content}
                    </MiniUserCard>
                ))
            ) : (
                <NoNotifications />
            )}
        </div>
    );
};
