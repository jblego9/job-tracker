import AddModal from "./add-modal";
import PageControls from "./page-controls";
import StatusFilter from "./status-filter";

export default function ControlBar({ pageNumber, onPageNumberChange, pageSize, onPageSizeChange, addingApp, onAddingApp, onAdd, statusFilter, onStatusFilterChange }) {
    return (
        <div className="applications-control-bar">
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
    );
}
