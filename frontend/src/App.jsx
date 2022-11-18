import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import Chat from './features/Chat/Chat';
import Profile from './pages/profile/Profile';
import ProfileInfoPopUp from './features/profile/ProfileInfoPopUp';
import Stories from './pages/stories/stories';
import Comments from './features/Comments';
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
function App() {
    const [authorised, setAuthorised] = useState(false);
    Validation(setAuthorised);
    const [socket, setSocket] = useState(null);
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    const [isSearchModalOpen, setIsSearchModalOpen] = useState(false);
    const { width } = useWindowDimensions();
    const [query, setQuery] = useState('');
    let isMobile = width < 600;
    return (
        <div
            className='App'
            onClick={() => {
                setIsMenuOpen(false);
                setIsSearchModalOpen(false);
            }}>
            {authorised && (
                <Header
                    setSocket={setSocket}
                    setIsMenuOpen={setIsMenuOpen}
                    setIsSearchModalOpen={setIsSearchModalOpen}
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
                            element={<Notification socket={socket} />}
                        />
                        <Route path='/profile' element={<Profile />} />
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
                <Footer setIsSearchModalOpen={setIsSearchModalOpen} />
            ) : null}
        </div>
    );
}

export default App;
