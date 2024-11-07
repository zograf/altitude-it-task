import React, { useState, useEffect, useRef } from "react";
import axios from "axios";
import {jwtDecode} from 'jwt-decode'
import './LoginComponent.css'
import { API, IMG } from "../environment";

export function ProfileComponent() {
    const token = localStorage.getItem("token");

    const [username, setUsername] = useState("")
    const [name, setName] = useState("")
    const [lastName, setLastName] = useState("")
    const [image, setImage] = useState(null)
    const [birthday, setBirthday] = useState("")
    const imageRef = useRef(null)

    const [nameDisabled, setNameDisabled] = useState(true)
    const [lastNameDisabled, setLastNameDisabled] = useState(true)
    const [birthdayDisabled, setBirthdayDisabled] = useState(true)
    const [imageDisabled, setImageDisabled] = useState(true)
    const [displaySaveCancel, setDisplaySaveCancel] = useState("none")
    const [displayEdit, setDisplayEdit] = useState("")

    const handleLastName = (e) => setLastName(e.target.value);
    const handleUsername = (e) => setUsername(e.target.value);
    const handleBirthday = (e) => setBirthday(e.target.value);
    const handleName = (e) => setName(e.target.value);
    const handleImage = (e) => setImage(e.target.files[0]);

    useEffect(() => {
        const decoded = jwtDecode(token);
        const email = decoded.email;

        axios.get(API + "/user", { headers: {"Authorization" : `Bearer ${token}`} })
            .then(response => { 
                const data = response.data.user
                setName(data.name)
                setBirthday(data.birthday.split(" ")[0])
                setLastName(data.last_name)
                setUsername(data.email)
            })
            .catch(e => console.log(e))

        axios.get(IMG + "/" + email, {responseType: 'blob'})
            .then(response => { 
                setImage(response.data)
            })
            .catch(e => console.log(e))

    }, [])

    const handleEdit = (e) => {
        e.preventDefault()
        setNameDisabled(false)
        setLastNameDisabled(false)
        setBirthdayDisabled(false)
        setImageDisabled(false)
        setDisplayEdit("none")
        setDisplaySaveCancel("")
    }

    const handleCancel = (e) => {
        setNameDisabled(true)
        setLastNameDisabled(true)
        setBirthdayDisabled(true)
        setImageDisabled(true)
        setDisplaySaveCancel("none")
        setDisplayEdit("")

        const decoded = jwtDecode(token);
        const email = decoded.email;

        axios.get(API + "/user", { headers: {"Authorization" : `Bearer ${token}`} })
            .then(response => { 
                const data = response.data.user
                setName(data.name)
                setBirthday(data.birthday.split(" ")[0])
                setLastName(data.last_name)
                setUsername(data.email)
            })
            .catch(e => console.log(e))

        axios.get(IMG + "/" + email, {responseType: 'blob'})
            .then(response => { 
                setImage(response.data)
            })
            .catch(e => console.log(e))
    }

    const handleSave = (e) => {
        setDisplaySaveCancel("none")
        setDisplayEdit("")

        if (name == "" || name == null || lastName == "" || lastName == null || birthday == "" || birthday == null) {
            handleCancel()
            return
        }

        setNameDisabled(true)
        setLastNameDisabled(true)
        setBirthdayDisabled(true)
        setImageDisabled(true)

        let payload = new FormData()
        payload.append("image", image)
        payload.append("name", name)
        payload.append("last_name", lastName)
        payload.append("birthday", birthday)
        payload.append("email", username)
        for (var pair of payload.entries()) {
            console.log(pair[0]+ ', ' + pair[1]); 
        }


        axios.post(API + "/user", payload,  { headers: {"Authorization" : `Bearer ${token}`} })
            .then(response => {
            })
            .catch(e => {
                console.log(e)
                // TODO: Popup goes here
                alert("Failed to change information")
            })
    }

    const imageHandler = () => {
        imageRef.current.click()
    }

    return(
        <div className="card" style={{minWidth: "500px"}}>
            <div className="flex center justify-center">
                <h1 style={{marginBottom: "20px"}}>Your information</h1>
            </div>
            <div className="flex center justify-center">
                <img
                    alt="not found"
                    onClick={() => imageHandler()}
                    style={{width: "100px", height: "100px", borderRadius: "50%", border: "1px solid black", marginBottom: "10px"}}
                    src={image == null ? require('../img/default.png') : URL.createObjectURL(image)}
                />
            </div>

            <div className="flex gap-s">
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">person</span>
                    <input placeholder="Name" value={name}
                            onChange={handleName}
                            disabled={nameDisabled}
                    />
                </div>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">badge</span>
                    <input placeholder="Last name" value={lastName}
                            onChange={handleLastName}
                            disabled={lastNameDisabled}
                    />
                </div>
            </div>
            <div className="input-wrapper regular-border v-spacer-xs">
                <span className="material-symbols-outlined icon input-icon">mail</span>
                <input placeholder="Email" value={username} disabled
                        onChange={handleUsername}
                />
            </div>
            <div className="input-wrapper regular-border v-spacer-s">
                <span className="material-symbols-outlined icon input-icon">cake</span>
                <input placeholder="Birthday" type="date" value={birthday}
                        onChange={handleBirthday}
                        disabled={birthdayDisabled}
                />
            </div>
            <div className='flex gap-xs justify-center' style={{display: displayEdit}}>
                <button className='small-button solid-button register' onClick={handleEdit}>Edit</button>
            </div>
            <div className="flex gap-l justify-center" style={{display: displaySaveCancel}}>
                <button className='small-button solid-error-button' onClick={handleCancel} style={{width: "40%"}}>Cancel</button>
                <button className='small-button solid-button' onClick={handleSave} style={{width: "40%"}}>Save</button>
            </div>

            <input
                style={{display: "none"}}
                type="file"
                name="image"
                onChange={handleImage}
                ref={imageRef}
                disabled={imageDisabled}>
            </input>
        </div>
    )
}