export default function SortControls({ sortBy, onSortByChange, sortDir, onSortDirChange }) {
    return (
        <div className="applications-control-bar__sort-controls">
            <div>
                <label htmlFor="sort-by">Sort By:</label>
                <select
                    id="sort-by"
                    name="sort-by"
                    defaultValue={sortBy}
                    onChange={() => onSortByChange(event.target.value)}
                >
                    <option value="">No Sort</option>
                    <option value="date">Date</option>
                    <option value="company">Company</option>
                    <option value="role">Role</option>
                    <option value="location">Location</option>
                </select>
            </div>

            <div>
                <label htmlFor="sort-dir">Sort Direction:</label>
                <select
                    id="sort-dir"
                    name="sort-dir"
                    defaultValue={sortDir}
                    onChange={() => onSortDirChange(event.target.value)}
                >
                    <option value="asc">Ascending</option>
                    <option value="desc">Descending</option>
                </select>
            </div>
        </div>
    );
}
