import { useEffect, useState } from 'react';
import config from '../../../../config';
export const GetNotif = (type, setNotifications, newNotif) => {
    const [notif, setNotif] = useState([]);
    useEffect(() => {
        fetch(
            config.api +
                '/notification?' +
                new URLSearchParams({ notifType: type }),
            {
                credentials: 'include',
            }
        )
            .then(async (res) => {
                let data = await res.json();
                console.log(data);
                data = data
                    ? data.sort(
                          (a, b) =>
                              new Date(b.createdAt) - new Date(a.createdAt)
                      )
                    : [];
                data ? setNotif(data) : setNotif([]);
                data ? setNotifications(data) : setNotifications([]);
                return data;
            })
            .catch(() => {
                setNotif([]);
            });
    }, [newNotif]);
    return notif;
};
