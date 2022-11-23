import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes, useLocation } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import Chat from './features/Chat/Chat';
import Profile from './pages/profile/Profile';
import ProfileInfoPopUp from './features/profile/ProfileInfoPopUp';
import Stories from './pages/stories/stories';
import Comments from './features/Comments';
import Groups from './pages/Groups';
import { useEffect, useState } from 'react';
import NewPost from './features/newpost/NewPost';
import { Notification } from './features/Notification/Notification';
import Validation from './components/Validation/Validation';
import { Menu } from './layouts/Menu/Menu';
import ValidRedirect from './components/Validation/ValidRedirect';
import useWindowDimensions from './components/hooks/useWindowDimensions';
import CloseFriendsUsers from './features/profile/CloseFriendsUsers';
import Followers from './features/profile/Followers';
import Following from './features/profile/Following';
import { SearchModal } from './features/Search/SearchModal';
import { CreateWebSocket } from './utils/createWebsocket';
function App() {
    const [authorised, setAuthorised] = useState(false);
    Validation(setAuthorised);
    const [socket, setSocket] = useState(null);
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    const [isSearchModalOpen, setIsSearchModalOpen] = useState(false);
    const { width } = useWindowDimensions();
    const [query, setQuery] = useState('');
    const [messageNotif, setMessageNotif] = useState(false);
    const [followNotif, setFollowNotif] = useState(false);
    const [followNotifContainer, setFollowNotifContainer] = useState();
    const { pathname } = useLocation();
    let isMobile = width < 600;
    useEffect(() => {
        if (authorised) {
            setSocket(CreateWebSocket());
        }
    }, [authorised]);
    let [newMessage, setNewMessage] = useState(0);
    const [chatList, setClist] = useState([]);
    if (socket) {
        socket.onmessage = (e) => {
            let data = JSON.parse(e.data);
            if (
                data.type === 'privateMessage' ||
                data.type === 'groupMessage'
            ) {
                console.log('new message');
                setNewMessage((prev) => prev + 1);
            }
            if (data.type === 'followRequest') {
                setFollowNotifContainer((prev) => {
                    console.log(prev);
                    for (let i = 0; i < prev.length; i++) {
                        if (prev[i].userId.id == data.data.userId.id) {
                            return prev;
                        }
                    }
                    return [data.data, ...prev];
                });
            }
        };
    }
    useEffect(() => {
        if (pathname !== '/notifications/follow') {
            if (Array.isArray(followNotifContainer)) {
                for (let i = 0; i < followNotifContainer.length; i++) {
                    if (followNotifContainer[i].read === 0) {
                        setFollowNotif(true);
                        return;
                    }
                }
            }
        }
    }, [followNotifContainer]);
    return (
        <div
            className='App'
            onClick={() => {
                setIsMenuOpen(false);
                setIsSearchModalOpen(false);
            }}>
            {authorised && (
                <Header
                    socket={socket}
                    setIsMenuOpen={setIsMenuOpen}
                    setIsSearchModalOpen={setIsSearchModalOpen}
                    messageNotif={messageNotif}
                    setMessageNotif={setMessageNotif}
                    setClist={setClist}
                    newMessage={newMessage}
                    chatList={chatList}
                    followNotif={followNotif}
                    setFollowNotifContainer={setFollowNotifContainer}
                    setFollowNotif={setFollowNotif}
                    followNotifContainer={followNotifContainer}
                    onChange={(e) => {
                        setQuery(e.target.value);
                    }}
                />
            )}
            <>
                {isMenuOpen ? (
                    <Menu setIsMenuOpen={setIsMenuOpen} auth={setAuthorised} />
                ) : null}
            </>
            <>{isSearchModalOpen ? <SearchModal query={query} /> : null}</>
            <>
                <Routes>
                    <Route
                        path='/'
                        element={
                            <ValidRedirect>
                                <Login auth={setAuthorised} />
                            </ValidRedirect>
                        }
                    />
                    <Route
                        path='/register'
                        element={
                            <ValidRedirect>
                                <Registration />
                            </ValidRedirect>
                        }
                    />
                    {/* need to be replaced */}
                    <Route path='*' element={<></>} />
                </Routes>
            </>
            {authorised && (
                <>
                    <Routes>
                        <Route
                            path='/home'
                            element={<Home isMobile={isMobile} />}
                        />
                        <Route path='/newpost' element={<NewPost />} />
                        <Route
                            path='/messages/*'
                            element={
                                <Chat
                                    isMobile={isMobile}
                                    socket={socket}
                                    chatList={chatList}
                                    setMessageNotif={setMessageNotif}
                                    setNewMessage={setNewMessage}
                                />
                            }
                        />
                        <Route
                            path='/groups'
                            element={<Groups socket={socket} />}
                        />
                        <Route
                            path='/comments'
                            element={<Comments isMobile={isMobile} />}
                        />
                        <Route path='notifications' element={<></>} />
                        <Route
                            path='notifications/:type'
                            element={
                                <Notification
                                    socket={socket}
                                    followNotif={followNotif}
                                    setFollowNotif={setFollowNotif}
                                    followNotifContainer={followNotifContainer}
                                    setFollowNotifContainer={
                                        setFollowNotifContainer
                                    }
                                />
                            }
                        />
                        <Route
                            path='/profile'
                            element={<Profile socket={socket} />}
                        />
                        <Route
                            path='/closefriends'
                            element={<CloseFriendsUsers />}
                        />
                        <Route path='/followers' element={<Followers />} />
                        <Route path='/following' element={<Following />} />
                        <Route
                            path='/update'
                            element={<ProfileInfoPopUp styleName='popUp' />}
                        />
                        <Route path='/stories' element={<Stories />} />
                    </Routes>
                </>
            )}
            {authorised ? (
                <Footer
                    setIsSearchModalOpen={setIsSearchModalOpen}
                    messageNotif={messageNotif}
                />
            ) : null}
        </div>
    );
}

export default App;
