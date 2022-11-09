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
function App() {
    const [isMobile, setIsMobile] = useState(false);
    const [authorised, setAuthorised] = useState(false);
    const [socket, setSocket] = useState(null);
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    return (
        <div
            className='App'
            onClick={() => {
                setIsMenuOpen(false);
            }}
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
                    <Route
                        path='/home'
                        element={
                            <Validation auth={setAuthorised}>
                                <Home isMobile={isMobile} />
                            </Validation>
                        }
                    />
                    <Route
                        path='/search'
                        element={
                            <Validation auth={setAuthorised}>
                                <Search />
                            </Validation>
                        }
                    />
                    <Route
                        path='/newpost'
                        element={
                            <Validation auth={setAuthorised}>
                                <NewPost />
                            </Validation>
                        }
                    />
                    <Route
                        path='/messages'
                        element={
                            <Validation auth={setAuthorised}>
                                <Chat isMobile={isMobile} socket={socket} />
                            </Validation>
                        }
                    />
                    <Route
                        path='/groups'
                        element={
                            <Validation auth={setAuthorised}>
                                <h1>Groups</h1>
                            </Validation>
                        }
                    />
                    <Route
                        path='/comments'
                        element={
                            <Validation auth={setAuthorised}>
                                <Comments isMobile={isMobile} />
                            </Validation>
                        }
                    />
                    <Route
                        path='notifications/:type'
                        element={
                            <Validation auth={setAuthorised}>
                                <Notification />
                            </Validation>
                        }
                    />
                    <Route
                        path='profile'
                        element={
                            <Validation auth={setAuthorised}>
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
                            <Validation auth={setAuthorised}>
                                <Stories />
                            </Validation>
                        }
                    />
                    <Route path='*' element={<h1>404</h1>} />
                </Routes>
            </>
            {authorised ? <Footer /> : null}
        </div>
    );
}

export default App;
