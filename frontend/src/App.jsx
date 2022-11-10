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
import { Menu } from './layouts/Menu/Menu';
import ValidRedirect from './components/Validation/ValidRedirect';
import useWindowDimensions from './components/hooks/useWindowDimensions';
function App() {
    const [auth, setAuthorised] = useState(false);
    const authorised = Validation(auth);
    const [socket, setSocket] = useState(null);
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    const { width } = useWindowDimensions();
    let isMobile = width < 600;
    return (
        <div
            className='App'
            onClick={() => {
                setIsMenuOpen(false);
            }}>
            {authorised && (
                <Header setSocket={setSocket} setIsMenuOpen={setIsMenuOpen} />
            )}
            <>
                {isMenuOpen && (
                    <Menu setIsMenuOpen={setIsMenuOpen} auth={setAuthorised} />
                )}
            </>
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
                        <Route path='/search' element={<Search />} />
                        <Route path='/newpost' element={<NewPost />} />
                        <Route
                            path='/messages'
                            element={
                                <Chat isMobile={isMobile} socket={socket} />
                            }
                        />
                        <Route path='/groups' element={<h1>Groups</h1>} />
                        <Route
                            path='/comments'
                            element={<Comments isMobile={isMobile} />}
                        />
                        <Route path='notifications' element={<></>} />
                        <Route
                            path='notifications/:type'
                            element={<Notification />}
                        />
                        <Route path='profile' element={<Profile />} />
                        <Route path='/stories' element={<Stories />} />
                    </Routes>
                </>
            )}
            {authorised ? <Footer /> : null}
        </div>
    );
}

export default App;
