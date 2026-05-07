import { NavLink } from "react-router";

export default function Header() {
    return (
        <header className="header">
            <NavLink to="/">Job Tracker</NavLink>
        </header>
    );
}
