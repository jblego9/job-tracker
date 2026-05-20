import Application from "./application.jsx";

export default function ApplicationsList({ apps, onEdit }) {
    return (
        <ol className="applications-list">
            {apps.map(app => (<Application key={app.Id} app={app} onEdit={onEdit} />))}
        </ol>
    );
}
