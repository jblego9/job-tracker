import { useState } from "react";
import EditModal from "./edit-modal.jsx";
import AddModal from "./add-modal.jsx";

export default function ApplicationsList({ applications }) {
    const [appToEdit, setAppToEdit] = useState(null);
    const [addingApp, setAddingApp] = useState(false);

    return (
        <>
            <div className="applications-control-bar">
                <button className="button" onClick={() => setAddingApp(true)}>Add Application</button>
                { addingApp && (
                    <AddModal onClose={() => setAppToEdit(false)}/>
                )}
            </div>
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
                                {hasDate(app.Date) && (
                                    <>
                                        <p>{app.Date}</p>
                                        <p>-</p>
                                    </>
                                )}
                                <p>{app.Status}</p>
                            </div>
                        </div>
                        <div>
                            <button className="button" onClick={() => setAppToEdit(app)}>Edit</button>
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

function hasDate(date) {
    return date != null && date != "" && date != undefined;
}
