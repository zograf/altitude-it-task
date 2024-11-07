import { LoginComponent } from '../../components/LoginComponent'
import './LoginPage.css'

export function LoginPage() {
    return(
        <main className="mh-100 body">
            <div className="flex center justify-center login-card">
                <LoginComponent/>

            </div>
        </main>
    )
}