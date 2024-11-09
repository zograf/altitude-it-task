import { AdminSidebar } from "../components/Sidebar"
import './UserPage.css'

export default function AdminPage(props) {
    const handleLogout = () => { window.location.href = "/" }

    return(
        <main className="mh-100">
            <div className="sidebar-root flex gap-l">
                <AdminSidebar/>
                {props.children}
            </div>
        </main>
    )
}