import React, { useState, useEffect, useRef } from "react";
import axios from "axios";
import {jwtDecode} from 'jwt-decode'
import './UsersComponent.css'
import { API, IMG } from "../environment";
import { DropDownSelect } from "./DropDown";

export function UsersComponent() {
    const token = localStorage.getItem("token");

    const [users, setUsers] = useState([]);
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);
    const [total, setTotal] = useState(0);
    const [birthdayFrom, setBirthdayFrom] = useState("");
    const [birthdayTo, setBirthdayTo] = useState("");
    const [email, setEmail] = useState("");
    const [enabled, setEnabled] = useState("");

    const handleBirthdayFrom = (e) => setBirthdayFrom(e.target.value);
    const handleBirthdayTo = (e) => setBirthdayTo(e.target.value);
    const handleEmail = (e) => setEmail(e.target.value);
    const handleClear = (e) => {
        setEmail("")
        setBirthdayFrom("")
        setBirthdayTo("")
    };

    useEffect(() => {
        sendRequest(page, pageSize, email, birthdayFrom, birthdayTo, enabled)
    }, [page, email, birthdayFrom, birthdayTo, enabled]);

    const sendRequest = (page, pageSize, email, birthdayFrom, birthdayTo, verified) => {
        const decoded = jwtDecode(token);

        let url = `/users?page=${page}&pageSize=${pageSize}&email=${email}&birthdayFrom=${birthdayFrom}&birthdayTo=${birthdayTo}&enabled=${verified}`
        console.log(url)

        axios.get(API + url, { headers: {"Authorization" : `Bearer ${token}`} })
            .then(response => { 
                const data = response.data
                console.log(data)
                setUsers(data.users)
                setPage(data.page)
                setPageSize(data.pageSize)
                setTotal(data.total)
            })
            .catch(e => console.log(e))
    }

    const handleDeleteRestore = (userId) => {
        const updatedUsers = users.map((user) =>
            user.ID === userId ? { ...user, IsDeleted: !user.IsDeleted } : user
        );
        setUsers(updatedUsers);
    };

    return (
        <div className="card" style={{width: "fit-content"}}>
            <h1 className="v-spacer-s">Users</h1>

            <div className="flex gap-s center v-spacer-xs">
                <p>Email:</p>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">mail</span>
                    <input placeholder="Email" value={email} onChange={handleEmail} />
                </div>
                <p>Verified:</p>
                <div>
                    <DropDownSelect
                    options={[{label: "None", value: null}, {label: "Yes", value: true}, {label: "No", value: false}]}
                    icon={"verified"}
                    placeholder={"Verified"}
                    callback={(value)=>{setEnabled(value)}}>
                    </DropDownSelect>
                </div>
                <p>From:</p>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">cake</span>
                    <input placeholder="Birthday from" type="date" value={birthdayFrom}
                            onChange={handleBirthdayFrom}/>
                </div>
                <p>To:</p>
                <div className="input-wrapper regular-border v-spacer-xs">
                    <span className="material-symbols-outlined icon input-icon">cake</span>
                    <input placeholder="Birthday to" type="date" value={birthdayTo}
                            onChange={handleBirthdayTo}/>
                </div>

                <button className='small-button solid-button' onClick={handleClear}>Clear</button>
            </div>

            <UserTable users={users} onDeleteRestore={handleDeleteRestore} />
        </div>
    );
};

const UserTable = ({ users, onDeleteRestore }) => {
  return (
    <div className="user-table">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Last Name</th>
            <th>Email</th>
            <th>Birthday</th>
            <th>Enabled</th>
            <th>Deleted</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.ID}>
              <td>{user.Name}</td>
              <td>{user.LastName}</td>
              <td>{user.Email}</td>
              <td>{new Date(user.Birthday).toLocaleDateString()}</td>
              <td>{user.IsEnabled ? "Yes" : "No"}</td>
              <td>{user.IsDeleted ? "Yes" : "No"}</td>
              <td className="flex justify-center">
                <button className={`small-button ${user.IsDeleted ? "solid-button" : "solid-error-button"}`}
                        onClick={() => onDeleteRestore(user.ID)}>
                        {user.IsDeleted ? "Restore" : "Delete"}
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};