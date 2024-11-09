import React, { useState, useEffect, useRef } from "react";
import axios from "axios";
import {jwtDecode} from 'jwt-decode'
import './PasswordComponent.css'
import { API, IMG } from "../environment";

export function PasswordComponent() {
    const token = localStorage.getItem("token");

    const [username, setUsername] = useState("")
    const [oldPassword, setOldPassword] = useState("")
    const [newPassword, setNewPassword] = useState("")
    const [repeatPassword, setRepeatPassword] = useState("")

    const handleUsername = (e) => setUsername(e.target.value);
    const handleOldPassword = (e) => setOldPassword(e.target.value);
    const handleNewPassword = (e) => setNewPassword(e.target.value);
    const handleRepeatPassword = (e) => setRepeatPassword(e.target.value);

    const handleSubmit = (e) => {
        e.preventDefault()

        let payload = {
            "old_password": oldPassword,
            "new_password": newPassword,
            "repeat_password": repeatPassword,
        }
        console.log(payload)

        axios.post(API + "/user/password", payload, { headers: {"Authorization" : `Bearer ${token}`} })
            .then(response => { 
                setOldPassword("")
                setNewPassword("")
                setRepeatPassword("")
            })
            .catch(e => console.log(e))
    }

    return(
        <div className="card" style={{minWidth: "400px"}}>
            <div className="flex center justify-center">
                <h1 style={{marginBottom: "20px"}}>Change your password</h1>
            </div>
                <div className="input-wrapper regular-border v-spacer-l" style={{marginBottom: "40px"}}>
                    <span className="material-symbols-outlined icon input-icon">lock_open</span>
                    <input type="password" placeholder="Old password" value={oldPassword}
                            onChange={handleOldPassword}
                    />
                </div>

                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">key</span>
                    <input type="password" placeholder="New password" value={newPassword}
                            onChange={handleNewPassword}
                    />
                </div>

                <div className="input-wrapper regular-border v-spacer-s">
                    <span className="material-symbols-outlined icon input-icon">lock</span>
                    <input type="password" placeholder="Repeat password" value={repeatPassword}
                            onChange={handleRepeatPassword}
                    />
                </div>
            <div className='flex gap-xs justify-center'>
                <button className='small-button solid-button' onClick={handleSubmit} style={{width: "90%"}}>Change password</button>
            </div>
        </div>
    )
}