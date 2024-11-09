import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
//import { LoginPage } from './pages/login/LoginPage'
import { RegisterPage } from './pages/login/RegisterPage'
import { LoginPage } from './pages/login/LoginPage';
import { ValidatePage } from './pages/ValidatePage';
import ProfilePage from './pages/ProfilePage';
import { PasswordComponent } from './components/PasswordComponent';
import PasswordPage from './pages/PasswordPage';
import UsersPage from './pages/UsersPage';
import { TotpPage } from './pages/TotpPage';

function App() {
    return (
        <main>
            <Router>
                <Routes>
                    <Route exact path='/' element={<LoginPage />} />
                    <Route exact path='/login' element={<LoginPage />} />
                    <Route exact path='/register' element={<RegisterPage />} />
                    <Route exact path='/validate' element={<ValidatePage />} />
                    <Route exact path='/profile' element={<ProfilePage />} />
                    <Route exact path='/password' element={<PasswordPage />} />
                    <Route exact path='/users' element={<UsersPage />} />
                    <Route exact path='/2fa' element={<TotpPage />} />
                </Routes>
            </Router>
        </main>
    );
}

export default App;