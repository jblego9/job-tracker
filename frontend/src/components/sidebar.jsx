import { NavLink } from "react-router";

export default function Sidebar() {
    return (
        <nav className="sidebar">
            <NavLink
                to="/" 
                className={({ isActive }) => isActiveClass(isActive)}
            >
                Applications
            </NavLink>
            <NavLink 
                to="/summary"
                className={({ isActive }) => isActiveClass(isActive)}
            >
                Summary
            </NavLink>
        </nav>
    );
}

function isActiveClass(isActive) {
    return isActive ? "sidebar__link sidebar__link--active" : "sidebar__link";
}
