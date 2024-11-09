import {useEffect, useState} from "react";
import axios from "axios";
import {API} from "../environment";
import UserPage from "./UserPage";
import { UserSidebar } from "../components/Sidebar";
import { ProfileComponent } from "../components/ProfileComponent";
import { PasswordComponent } from "../components/PasswordComponent";

export default function PasswordPage() {
    const token = localStorage.getItem("token")
    const userId = localStorage.getItem("id")


    return (
        <UserPage>
            <div className="profile">
                <p className="page-title">Password</p>
                <PasswordComponent />
            </div>
        </UserPage>
        )
}