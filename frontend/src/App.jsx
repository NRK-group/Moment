import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home'

function App() {
    return (
        <div className='App'>
            <Header />
            <Home/>
            <Footer />
        </div>
    );
}

export default App;
