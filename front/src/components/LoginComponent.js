import React, { useState, useMemo } from "react";
import { GoogleLogin, GoogleOAuthProvider } from '@react-oauth/google';
import axios from "axios";
import jwt from 'jwt-decode'
import './LoginComponent.css'
import { API } from "../environment";
import { MessagePopUp, StandardPopUp, usePopup } from "./PopUp";

export function LoginComponent() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const handleUsername = (e) => setUsername(e.target.value);
    const handlePassword = (e) => setPassword(e.target.value);

    useMemo(() => {
        localStorage.setItem("token", "")
        localStorage.setItem("isUser", false)
    }, [])

    const handleSubmit = (e) => {
        e.preventDefault()
        let payload = { "email": username, "password": password }

        axios.post(API + "/login", payload)
            .then(response => {
                if (response.status == 204) {
                    localStorage.setItem("email", username)
                    window.location.href = "/2fa"
                    return
                }
                localStorage.setItem("token", response.data.token)
                localStorage.setItem("isUser", !response.data.is_admin)
                if (!response.data.is_admin) {
                    window.location.href = "/profile"
                } else {
                    window.location.href = "/users"
                }
            })
            .catch(e => {
                if (e.response.data.error) {
                    setPopUpMessage(e.response.data.error)
                } else {
                    setPopUpMessage("An error occured")
                }
                notificationPopUp.showPopup()
            }
            )
    }

    const notificationPopUp = usePopup()
    const [popUpTitle, setPopUpTitle] = useState("Notification")
    const [popUpMessage, setPopUpMessage] = useState("")

    return(
        <div>
            <div className="card">
                <div className="flex center justify-center">
                    <h1 style={{marginBottom: "20px"}}>Login</h1>
                </div>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">mail</span>
                    <input placeholder="Username" value={username} onChange={handleUsername} />
                </div>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">key</span>
                    <input placeholder="Password" type="Password" value={password} onChange={handlePassword}/>
                </div>
                <div className="v-spacer-m">
                    <a href="/register" className="register-link">Don't have an account?</a>
                </div>
                <div className='flex gap-xs justify-center v-spacer-m'>
                    <button className='small-button solid-accent-button login' onClick={handleSubmit}>Login</button>
                </div>
                <div className="v-spacer-m" style={{ display: 'flex', alignItems: 'center', margin: '20px 0' }}>
                    <hr style={{ flex: 1, border: 'none', borderTop: '1px solid #ccc' }} />
                    <span style={{ padding: '0 10px', color: '#666' }}>Or</span>
                    <hr style={{ flex: 1, border: 'none', borderTop: '1px solid #ccc' }} />
                </div>

                <div className='v-spacer-xs'>
                    <GoogleOAuthProvider clientId="85724573809-pgeu3te2gm2198mm7feon6e725vaf9k1.apps.googleusercontent.com" locale="fr">
                        <GoogleLogin
                            onSuccess={(credentialResponse) => {
                                const { credential } = credentialResponse;
                                fetch(API + '/auth/google', {
                                    method: 'POST',
                                    headers: { 'Content-Type': 'application/json' },
                                    body: JSON.stringify({ token: credential }),
                                })
                                .then(response => response.json())
                                .then(data => {
                                    // Handle token and login status
                                    console.log(data);
                                });
                            }}
                            onError={() => {
                                console.log('Login failed');
                            }}
                        />
                    </GoogleOAuthProvider>
                </div>
                <MessagePopUp popup={notificationPopUp} title={popUpTitle} message={popUpMessage}/>
            </div>
        </div>
    )
}