import AdminPage from "./AdminPage";
import { ProfileComponent } from "../components/ProfileComponent";
import { UsersComponent } from "../components/UsersComponent";

export default function UsersPage() {
    const token = localStorage.getItem("token")
    const userId = localStorage.getItem("id")


    return (
        <AdminPage>
            <div className="profile">
                <p className="page-title">Admin page</p>
                <UsersComponent />
            </div>
        </AdminPage>
        )
}