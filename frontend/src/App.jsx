import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/LoginPage';

function App() {
    return (
        <div className='App'>
            <Header />
            {/* <Login /> */}
            <>
                <Routes>
                    <Route path='/home' element={<Home />} />
                    <Route path='/search' element={<h1>Search</h1>} />
                    <Route path='/newpost' element={<h1>Newpost</h1>} />
                    <Route path='/messages' element={<h1>Messages</h1>} />
                    <Route path='/groups' element={<h1>Groups</h1>} />
                </Routes>
            </>
            <Footer />
        </div>
    );
}

export default App;
