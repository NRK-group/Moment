import { useEffect, useState } from 'react';
export const GetNotif = (type, setNotifications) => {
    const [notif, setNotif] = useState([]);
    useEffect(() => {
        fetch(
            'http://localhost:5070/notification?' +
                new URLSearchParams({ notifType: type }),
            {
                credentials: 'include',
            }
        )
            .then(async (res) => {
                const data = await res.json();
                setNotif(data);
                setNotifications(data);
                return data;
            })
            .catch(() => {
                setNotif([]);
            });
    }, []);
    return notif;
};
