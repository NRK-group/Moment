import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import Profile from './pages/profile/Profile';
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
                    <Route path='/messages' element={<h1>Messages</h1>} />
                    <Route path='/groups' element={<h1>Groups</h1>} />
                    <Route
                        path='/notifications'
                        element={<h1>Notifications</h1>}
                    />
                    <Route path='/profile' element={<Profile />} />
                </Routes>
            </>
            <Footer />
        </div>
    );
}

export default App;
