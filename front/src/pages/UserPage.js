import { UserSidebar } from "../components/Sidebar"
import './UserPage.css'

export default function UserPage(props) {
    const handleLogout = () => { window.location.href = "/" }

    return(
        <main className="mh-100">
            <div className="sidebar-root flex gap-l">
                <UserSidebar/>
                {props.children}
            </div>
        </main>
    )
}