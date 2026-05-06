export default function Sidebar() {
    return (
        <nav className="sidebar">
            <a className="sidebar__link sidebar__link--active" href="/">Applications</a>
            <a className="sidebar__link" href="/summary">Summary</a>
        </nav>
    );
}
