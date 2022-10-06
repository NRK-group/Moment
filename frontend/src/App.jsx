import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home'
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';

function App() {
    return (
        <div className='App'>
            {/* <Header /> */}
            <Registration />
            <Footer />
        </div>
    );
}

export default App;
