import AddModal from "./add-modal";
import PageControls from "./page-controls";
import SortControls from "./sort-controls";
import StatusFilter from "./status-filter";

export default function ControlBar({ pageNumber, onPageNumberChange, pageSize, onPageSizeChange, addingApp, onAddingApp, onAdd, statusFilter, onStatusFilterChange, sortBy, onSortByChange, sortDir, onSortDirChange, onExportCSV }) {
    return (
        <div className="applications-control-bar">
            <div>
                <PageControls
                    pageNumber={pageNumber}
                    onPageNumberChange={onPageNumberChange}
                    pageSize={pageSize}
                    onPageSizeChange={onPageSizeChange}
                />

                <StatusFilter
                    statusFilter={statusFilter}
                    onStatusFilterChange={onStatusFilterChange}
                />

                <SortControls
                    sortBy={sortBy}
                    onSortByChange={onSortByChange}
                    sortDir={sortDir}
                    onSortDirChange={onSortDirChange}
                />
            </div>

            <div className="applications-control-bar__end-buttons">
                <button
                    className="button"
                    onClick={onExportCSV}
                >
                    Export CSV
                </button>

                <button
                    className="button"
                    onClick={() => onAddingApp(true)}
                >
                    Add Application
                </button>
                {addingApp && (
                    <AddModal
                        onAdd={onAdd}
                        onCancel={() => onAddingApp(false)}
                    />
                )}
            </div>
        </div>
    );
}
