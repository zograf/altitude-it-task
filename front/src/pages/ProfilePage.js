import {useEffect, useState} from "react";
import axios from "axios";
import {API} from "../environment";
import UserPage from "./UserPage";
import { UserSidebar } from "../components/Sidebar";
import { ProfileComponent } from "../components/ProfileComponent";

export default function ProfilePage() {
    const token = localStorage.getItem("token")
    const userId = localStorage.getItem("id")


    return (
        <UserPage>
            <div className="profile">
                <p className="page-title">Profile</p>
                <ProfileComponent />
            </div>
        </UserPage>
        )
}