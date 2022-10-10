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
import Groups from './pages/Groups';
import { useState } from 'react';
function App() {
    const [isMobile, setIsMobile] = useState(false);
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
                    <Route path='/search' element={<h1>Search</h1>} />
                    <Route path='/newpost' element={<h1>Newpost</h1>} />
                    <Route
                        path='/messages'
                        element={
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
                        }
                    />
                    <Route
                        path='/groups'
                        element={
                            isMobile ? (
                                <Groups
                                    bodyStyleName='mobile'
                                    cardStyleName='mobileCard'
                                />
                            ) : (
                                <Groups
                                    bodyStyleName='desktop'
                                    cardStyleName='desktopCard'
                                />
                            )
                        }
                    />
                    <Route
                        path='/notifications'
                        element={<h1>Notifications</h1>}
                    />
                    <Route
                        path='/profile'
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
    );
}

export default App;
