import React, { useState, useMemo } from "react";
import axios from "axios";
import jwt from 'jwt-decode'
import './LoginComponent.css'
import { API } from "../environment";

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
                localStorage.setItem("token", response.data.token)
                localStorage.setItem("isUser", !response.data.is_admin)
                console.log(response);
            })
            .catch(e => {
                console.log(e);
                //TODO: Popup goes here
                alert("Incorrect combination of username and password")
            }
            )
    }

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
                <div className="v-spacer-s">
                    <a href="/register" className="register-link">Don't have an account?</a>
                </div>
                <div className='flex gap-xs justify-center'>
                    <button className='small-button solid-accent-button' onClick={handleSubmit}>Login</button>
                </div>
            </div>
        </div>
    )
}