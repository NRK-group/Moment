import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';

function App() {
    return (
        <div className='App'>
            <Header />
            {/* <Login /> */}
            {/* <Registration /> */}
            <>
                <Routes>
                    <Route path='/home' element={<Home />} />
                    <Route path='/search' element={<h1>Search</h1>} />
                    <Route path='/newpost' element={<h1>Newpost</h1>} />
                    <Route path='/messages' element={<h1>Messages</h1>} />
                    <Route path='/groups' element={<h1>Groups</h1>} />
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
