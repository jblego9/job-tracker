import { useState } from "react";
import EditModal from "./edit-modal";

export default function ApplicationsList({ applications }) {
    const [appToEdit, setAppToEdit] = useState(null);

    return (
        <>
            <ol className="applications-list">
                {applications.map(app => (
                    <li key={app.Id} className="application-card">
                        <div>
                            <div className="application-card__row application-card__row--main">
                                <p>{app.Company}</p>
                                <p>-</p>
                                <p>{app.Role}</p>
                            </div>
                            <div className="application-card__row application-card__row--details">
                                <p>{app.Location}</p>
                            </div>
                            <div className="application-card__row application-card__row--details application-card__row--spaced">
                                <p>{app.Date}</p>
                                <p>-</p>
                                <p>{app.Status}</p>
                            </div>
                        </div>
                        <div>
                            <button className="application-card__edit-button" onClick={() => setAppToEdit(app)}>Edit</button>
                        </div>
                    </li>
                ))}
            </ol>
            {appToEdit && (
                <EditModal app={appToEdit} onClose={() => setAppToEdit(null)}/>
            )}
        </>
    );
}
