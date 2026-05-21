export default function StatusFilter({ statusFilter, onStatusFilterChange }) {
    return (
        <div className="applications-control-bar__status-filter">
            <label for="status-filter">Filter:</label>
            <select
                id="status-filter"
                name="status-filter"
                defaultValue={statusFilter}
                onChange={() => onStatusFilterChange(event.target.value)}
            >
                <option value="">No Filter</option>
                <option value="AWAITING_RESPONSE">Awaiting Response</option>
                <option value="NO_RESPONSE">No Response</option>
                <option value="AWAITING_INTERVIEW">Awaiting Interview</option>
                <option value="REJECTED">Rejected</option>
                <option value="ACCEPTED">Accepted</option>
            </select>
        </div>
    );
}
