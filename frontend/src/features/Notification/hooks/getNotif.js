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
                let data = await res.json();
                data = data
                    ? data.sort(
                          (a, b) =>
                              new Date(b.updatedAt) - new Date(a.updatedAt)
                      )
                    : [];
                data ? setNotif(data) : setNotif([]);
                data ? setNotifications(data) : setNotifications([]);
                return data;
            })
            .catch(() => {
                setNotif([]);
            });
    }, []);
    return notif;
};
