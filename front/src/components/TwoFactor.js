import './TwoFactor.css'
import { useLocation } from 'react-router-dom';
import { useEffect, useMemo, useState } from 'react';
import { API } from "../environment";
import axios from "axios";

export function TwoFactor() {
    const [totp, setTotp] = useState("")

    const handleTotp = (e) => setTotp(e.target.value);

    const handleSubmit = () => {
        const email = localStorage.getItem("email")
        let payload = {
            "email": email,
            "totp_code": totp,
        }

        axios.post(API + "/totp", payload)
            .then(response => {
                localStorage.setItem("token", response.data.token)
                localStorage.setItem("isUser", !response.data.is_admin)
                if (!response.data.is_admin) {
                    window.location.href = "/profile"
                } else {
                    window.location.href = "/users"
                }
            })
            .catch(e => {
                console.log(e)
            })
    }

    return(
        <div>
            <div className="card">
                <h1 className='card-title v-spacer-s'> Enter your code </h1>
                    <div className="input-wrapper regular-border v-spacer-s">
                        <span className="material-symbols-outlined icon input-icon">123</span>
                        <input placeholder="Your code goes here" value={totp} onChange={handleTotp} />
                    </div>
                    <div className='flex gap-xs justify-center v-spacer-s'>
                        <button className='small-button solid-accent-button login' onClick={handleSubmit}>Login</button>
                    </div>
            </div>
        </div>
    )
}