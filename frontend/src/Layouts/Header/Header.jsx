import './Header.css';
import { DesktopHeaderNav, MobileHeaderNav } from '../Navbar/Navbar';
import Input from '../../Components/Input/Input';
import { useEffect } from 'react';
import { GetChatList } from '../../Features/Chat/Hooks/GetChatList';
import { GetNotif } from '../../Features/Notification/Hooks/GetNotif';
const Header = ({
    setIsMenuOpen,
    setIsSearchModalOpen,
    messageNotif,
    onChange,
    setMessageNotif,
    setClist,
    chatList,
    followNotif,
    setFollowNotifContainer,
    followNotifContainer,
    setFollowNotif,
    groupNotif,
    setGroupNotifContainer,
    groupNotifContainer,
    setGroupNotif,
    socket,
    newMessageNotif,
}) => {
    GetNotif('follow', setFollowNotifContainer);
    GetNotif('group', setGroupNotifContainer);
    GetChatList(setClist, newMessageNotif);
    if (socket) {
        socket.onmessage = (e) => {
            let data = JSON.parse(e.data);
            if (
                data.type === 'eventNotif' ||
                data.type === 'groupInvitationJoin' ||
                data.type === 'groupInvitationRequest'
            ) {
                console.log('new group notif');
                setGroupNotif(true);
            }
            if (data.type === 'followRequest') {
                setFollowNotif(true);
            }
            if (
                data.type === 'privateMessage' ||
                data.type === 'groupMessage'
            ) {
                console.log('new message');
                setMessageNotif(true);
            }
        };
    }
    useEffect(() => {
        if (Array.isArray(chatList)) {
            for (let i = 0; i < chatList.length; i++) {
                if (chatList[i].notif > 0) {
                    setMessageNotif(true);
                    return;
                }
            }
            setMessageNotif(false);
        }
    }, [chatList]);
    useEffect(() => {
        if (Array.isArray(followNotifContainer)) {
            for (let i = 0; i < followNotifContainer.length; i++) {
                if (followNotifContainer[i].read === 0) {
                    setFollowNotif(true);
                    return;
                }
            }
        }
    }, [followNotifContainer]);
    useEffect(() => {
        if (Array.isArray(groupNotifContainer)) {
            for (let i = 0; i < groupNotifContainer.length; i++) {
                if (groupNotifContainer[i].read === 0) {
                    setGroupNotif(true);
                    return;
                }
            }
        }
    }, [groupNotifContainer]);
    return (
        <div className='headerContainer'>
            <div className='header'>
                <div className='headerLogo'>
                    <svg
                        version='1.2'
                        xmlns='http://www.w3.org/2000/svg'
                        // viewBox='0 0 1900 1080'
                        width='100'
                        height='40'>
                        <title>Moment</title>
                        <defs>
                            <image
                                width='100'
                                height='40'
                                id='img1'
                                href='data:image/png;base64,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'
                            />
                        </defs>
                        <use id='Moment' href='#img1' x='0' y='0' />
                    </svg>
                </div>
                <div
                    onClick={(e) => {
                        e.stopPropagation();
                        setIsSearchModalOpen(true);
                    }}>
                    <Input
                        styleName={'search'}
                        type={'search'}
                        placeholder='search'
                        onChange={onChange}
                    />
                </div>
                <DesktopHeaderNav
                    setIsMenuOpen={setIsMenuOpen}
                    messageNotif={messageNotif}
                    followNotif={followNotif}
                    groupNotif={groupNotif}
                />
                <MobileHeaderNav
                    setIsMenuOpen={setIsMenuOpen}
                    followNotif={followNotif}
                    groupNotif={groupNotif}
                />
            </div>
        </div>
    );
};
export default Header;
