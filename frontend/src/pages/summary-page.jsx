import { useEffect, useState } from "react";
import * as Request from "../request.js";

export default function SummaryPage() {
    const [total, setTotal] = useState(null);
    const [awaitingResponse, setAwaitingResponse] = useState(null);
    const [noResponse, setNoResponse] = useState(null);
    const [awaitingInterview, setAwaitingInterview] = useState(null);
    const [rejected, setRejected] = useState(null);
    const [accepted, setAccepted] = useState(null);

    useEffect(() => {
        const getSummary = async () => {
            const response = await Request.getData(`${Request.apiUrl}/summary`);
            return await response.json();
        };


        getSummary().then((summary) => {
            setTotal(summary.Total);
            setAwaitingResponse(summary.TotalAwaitingResponse);
            setNoResponse(summary.TotalNoResponse);
            setAwaitingInterview(summary.TotalAwaitingInterview);
            setRejected(summary.TotalRejected);
            setAccepted(summary.TotalAccepted);
        });
    }, []);

    return (
        <div>
            <h1>Summary</h1>
            <div className="summary-list">
                <span className="summary-list__status summary-list__status--total">Total</span>
                <span></span>
                <span className="summary-list__count summary-list__count--total">{total}</span>

                <span className="summary-list__status">Awaiting Response</span>
                <span className="status-badge status-badge--awaiting-response"></span>
                <span className="summary-list__count">{awaitingResponse}</span>

                <span className="summary-list__status">No Response</span>
                <span className="status-badge status-badge--no-response"></span>
                <span className="summary-list__count">{noResponse}</span>

                <span className="summary-list__status">Awaiting Interview</span>
                <span className="status-badge status-badge--awaiting-interview"></span>
                <span className="summary-list__count">{awaitingInterview}</span>

                <span className="summary-list__status">Rejected</span>
                <span className="status-badge status-badge--rejected"></span>
                <span className="summary-list__count">{rejected}</span>

                <span className="summary-list__status">Accepted</span>
                <span className="status-badge status-badge--accepted"></span>
                <span className="summary-list__count">{accepted}</span>
            </div>
        </div>
    );
}
