import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import { useState } from 'react';
import Modal from './features/Modal';
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
                    <Route path='/modal' element={<Modal> <div className='ModalContent'/></Modal>} />
                    <Route
                        path='/notifications'
                        element={<h1>Notifications</h1>}
                    />
                    <Route path='/profile' element={<h1>Profile</h1>} />
                </Routes>
            </>
            <Footer />
        </div>
    );
}

export default App;
