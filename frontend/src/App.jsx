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
import { useState } from 'react';
import NewPost from './features/newpost/NewPost';
import { Notification } from './features/Notification/Notification';
import { Search } from './features/Search/Search';
import Validation from './components/Validation/Validation';
function App() {
    const [isMobile, setIsMobile] = useState(false);
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
    return (
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
                        path='/'
                        element={
                            <Validation>
                                <Login />
                            </Validation>
                        }
                    />
                    <Route
                        path='/register'
                        element={
                            <Validation>
                                <Registration />
                            </Validation>
                        }
                    />

                    <Route
                        path='/home'
                        element={
                            <Validation>
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
                            </Validation>
                        }
                    />
                    <Route
                        path='/search'
                        element={
                            <Validation>
                                <Search />
                            </Validation>
                        }
                    />
                    <Route
                        path='/newpost'
                        element={
                            <Validation>
                                <NewPost />
                            </Validation>
                        }
                    />
                    <Route
                        path='/messages'
                        element={
                            <Validation>
                                isMobile ? (
                                <Chat
                                    bodyStyleName='mobile'
                                    cardStyleName='mobileCard'
                                />
                                ) : (
                                <Chat
                                    bodyStyleName='desktop'
                                    cardStyleName='desktopCard'
                                />
                                )
                            </Validation>
                        }
                    />
                    <Route path='/groups' element={<h1>Groups</h1>} />
                    <Route
                        path='/comments'
                        element={
                            <Validation>
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
                            </Validation>
                        }
                    />
                    <Route
                        path='notifications'
                        element={
                            <Validation>
                                <Notification users={generalNotif} />
                            </Validation>
                        }>
                        <Route
                            path='general'
                            element={
                                <Validation>
                                    <Notification users={generalNotif} />
                                </Validation>
                            }
                        />
                        <Route
                            path='followrequest'
                            element={
                                <Validation>
                                    <Notification users={followrequest} />
                                </Validation>
                            }
                        />
                        <Route
                            path='group'
                            element={
                                <Validation>
                                    <Notification users={groupNotif} />
                                </Validation>
                            }
                        />
                    </Route>
                    <Route
                        path='profile'
                        element={
                            <Validation>
                                <Profile
                                    aboutMe='This section is where the bio goes. You should write 1-2 sentences about yourself.'
                                    fullname='Nathaniel Russell'
                                    nickname='Nate'
                                />
                            </Validation>
                        }
                    />
                    <Route
                        path='/stories'
                        element={
                            <Validation>
                                <Stories />
                            </Validation>
                        }
                    />
                </Routes>
            </>
            <Footer />
        </div>
    );
}

export default App;
