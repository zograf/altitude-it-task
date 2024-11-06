import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
//import { LoginPage } from './pages/login/LoginPage'
import { RegisterPage } from './pages/login/RegisterPage'
import { LoginPage } from './pages/login/LoginPage';
//import { ValidatePage } from './pages/login/ValidatePage'

function App() {
    return (
        <main>
            <Router>
                <Routes>
                    <Route exact path='/' element={<LoginPage />} />
                    <Route exact path='/login' element={<LoginPage />} />
                    <Route exact path='/register' element={<RegisterPage />} />
                </Routes>
            </Router>
        </main>
    );
}

export default App;