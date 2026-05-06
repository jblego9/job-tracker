export default function ApplicationsList({ applications }) {
    return (
        <ol className="applications-list">
            {applications.map(app => (
                <li key={app.Id} className="application">
                    <div>
                        <div className="application-row">
                            <p>{app.Company}</p>
                            <p>-</p>
                            <p>{app.Role}</p>
                        </div>
                        <div className="application-row application-row-margin-under">
                            <p>{app.Location}</p>
                        </div>
                        <div className="application-row">
                            <p>{app.Date}</p>
                            <p>-</p>
                            <p>{app.Status}</p>
                        </div>
                    </div>
                    <div>
                        <button className="application-edit-button" onClick={() => { console.log(`Edit application ${app.Id}`); }}>Edit</button>
                    </div>
                </li>
            ))}
        </ol>
    )
}
