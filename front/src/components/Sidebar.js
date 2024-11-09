import './Sidebar.css'

export function UserSidebar() {
    return(
        <main className='mh-100 sidebar'>
            <div className="">
                <SidebarTile icon={"logout"} label={"Logout"} path="/login"/>
            </div>
            <SidebarDivider />
            <div className="" style={{padding: '0 0 14px 0'}}>
                <SidebarTile icon={"person"} label={"Profile"} path="/profile"/>
                <SidebarTile icon={"key"} label={"Password"} path="/password"/>
            </div>
        </main>
    )
}

export function AdminSidebar() {
    return(
        <main className='mh-100 sidebar'>
            <div className="">
                <SidebarTile icon={"logout"} label={"Logout"} path="/login"/>
            </div>
            <SidebarDivider />
            <div className="" style={{padding: '0 0 14px 0'}}>
                <SidebarTile icon={"group"} label={"Users"} path="/users"/>
            </div>
        </main>
    )
}

export function SidebarTile({icon, label, path}) {
    const handleClick = () => {window.location.href=path}
    return(
        <div className="sidebar-tile-container" onClick={handleClick}>
            <span className="sidebar-tooltip">{label}</span>
            <span className="material-symbols-outlined icon">{icon}</span>
        </div>
    )
}

export function SidebarDivider() {
    return( <hr style={{width: "calc(100% - 24px)", marginTop: "2px", marginBottom: "2px"}}/> )
}