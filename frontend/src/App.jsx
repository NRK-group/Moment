import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home'
import Login from './pages/loginPage/LoginPage';

function App() {
    return (
        <div className='App'>
            {/* <Header /> */}
            <Login />
            <Footer />
        </div>
    );
}

export default App;
