export default function Application({ app, onEdit }) {
    return (
        <li className="application-card">
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
                    <span className={"status-badge " + statusBadgeFromStatus(app.Status)}></span>
                    <p>{app.Status}</p>
                </div>
            </div>
            <div>
                <button className="button" onClick={() => onEdit(app)}>Edit</button>
            </div>
        </li>
    );
}

function hasDate(date) {
    return date != null && date != "" && date != undefined;
}

function statusBadgeFromStatus(status) {
    if (status == "AWAITING_RESPONSE") return "status-badge--awaiting-response";
    if (status == "NO_RESPONSE") return "status-badge--no-response";
    if (status == "AWAITING_INTERVIEW") return "status-badge--awaiting-interview";
    if (status == "REJECTED") return "status-badge--rejected";
    if (status == "ACCEPTED") return "status-badge--accepted";
}
