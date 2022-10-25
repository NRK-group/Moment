import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import Chat from './features/Chat/Chat';
import Profile from './pages/profile/Profile';
import Stories from './pages/stories/stories';
import Comments from './features/Comments';
import { useEffect, useState } from 'react';
import NewPost from './features/newpost/NewPost';
import { Notification } from './features/Notification/Notification';
import { Search } from './features/Search/Search';
function App() {
    const [isMobile, setIsMobile] = useState(false);
    let socket;
    let generalNotif = [
        {
            name: 'John',
            id: 1,
            content: 'liked your post',
            optContent: '1h',
        },
    ];
    let followrequest = [{ name: 'Ken' }];
    let groupNotif = [];

    // for websocket
    const CreateWebSocket = () => {
        socket = new WebSocket('ws://' + 'localhost:5070' + '/ws');
        console.log('Attempting Connection...');
    };
    CreateWebSocket();
    return (
        socket && (
            <div
                className='App'
                ref={(boxRef) => {
                    boxRef &&
                        console.log(
                            boxRef.getBoundingClientRect().width,
                            boxRef.getBoundingClientRect().width >= 600
                        );
                    return (
                        boxRef &&
                        setIsMobile(boxRef.getBoundingClientRect().width < 600)
                    );
                }}>
                <Header />
                {/* <Login /> */}
                {/* <Registration /> */}
                <>
                    <Routes>
                        <Route
                            path='/home'
                            element={
                                isMobile ? (
                                    <Home
                                        bodyStyleName='mobile'
                                        cardStyleName='mobileCard'
                                    />
                                ) : (
                                    <Home
                                        bodyStyleName='desktop'
                                        cardStyleName='desktopCard'
                                    />
                                )
                            }
                        />
                        <Route path='/search' element={<Search />} />
                        <Route path='/newpost' element={<NewPost />} />
                        <Route
                            path='/messages'
                            element={
                                isMobile ? (
                                    <Chat
                                        bodyStyleName='mobile'
                                        cardStyleName='mobileCard'
                                        socket={socket}
                                    />
                                ) : (
                                    <Chat
                                        bodyStyleName='desktop'
                                        cardStyleName='desktopCard'
                                        socket={socket}
                                    />
                                )
                            }
                        />
                        <Route path='/groups' element={<h1>Groups</h1>} />
                        <Route
                            path='/comments'
                            element={
                                isMobile ? (
                                    <Comments
                                        bodyStyleName='mobile'
                                        cardStyleName='mobileCard'
                                    />
                                ) : (
                                    <Comments
                                        bodyStyleName='desktop'
                                        cardStyleName='desktopCard'
                                    />
                                )
                            }
                        />
                        <Route
                            path='notifications'
                            element={<Notification users={generalNotif} />}>
                            <Route
                                path='general'
                                element={<Notification users={generalNotif} />}
                            />
                            <Route
                                path='followrequest'
                                element={<Notification users={followrequest} />}
                            />
                            <Route
                                path='group'
                                element={<Notification users={groupNotif} />}
                            />
                        </Route>
                        <Route
                            path='profile'
                            element={
                                <Profile
                                    aboutMe='This section is where the bio goes. You should write 1-2 sentences about yourself.'
                                    fullname='Nathaniel Russell'
                                    nickname='Nate'
                                />
                            }
                        />
                        <Route path='/stories' element={<Stories />} />
                    </Routes>
                </>
                <Footer />
            </div>
        )
    );
}

export default App;
