import { TwoFactor } from '../components/TwoFactor'
import './TotpPage.css'

export function TotpPage() {

    return(
        <main className="mh-100">
            <div className="flex center justify-center validate-card">
                <TwoFactor/>

            </div>
        </main>
    )
}