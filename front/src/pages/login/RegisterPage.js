import { RegisterComponent } from '../../components/RegisterComponent'
import './RegisterPage.css'

export function RegisterPage() {
    return(
        <main className="mh-100 body">
            <div className="flex center justify-center register-card">
                <RegisterComponent/>

            </div>
        </main>
    )
}